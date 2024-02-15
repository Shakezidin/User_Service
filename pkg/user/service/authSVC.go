package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Shakezidin/config"
	DOM "github.com/Shakezidin/pkg/dom"
	pb "github.com/Shakezidin/pkg/user/userpb"

	"github.com/Shakezidin/utils"
	"gorm.io/gorm"
)

// SignupSVC handles the user signup process.
func (c *UserSVC) SignupSVC(p *pb.UserSignup) (*pb.UserResponce, error) {
	// Hash the user's password
	hashPassword, err := utils.HashPassword(p.Password)
	if err != nil {
		return nil, errors.New("error while hashing password")
	}

	// Convert phone to integer
	phone, _ := strconv.Atoi(p.Phone)

	// Create a new user object
	user := &DOM.User{
		Phone:    phone,
		Email:    p.Email,
		Password: string(hashPassword),
		Name:     p.Name,
		Role:     p.Role,
	}

	// Send OTP to the phone number
	resp, err := c.twilio.SendTwilioOTP(p.Phone)
	if err != nil {
		return nil, err
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}

	// Marshal user data to JSON
	userData, err := json.Marshal(&user)
	if err != nil {
		return nil, err
	}

	// Set user data in Redis
	registerUser := fmt.Sprintf("register_user_%v", p.Email)
	c.redis.Set(context.Background(), registerUser, userData, time.Minute*2)

	// Return success response
	return &pb.UserResponce{
		Status:  "success",
		Message: "user creation initiated, check message for OTP",
	}, nil
}

// VerifySVC handles the user verification process.
func (c *UserSVC) VerifySVC(p *pb.UserVerify) (*pb.UserResponce, error) {
	// Retrieve user data from Redis
	registerUser := fmt.Sprintf("register_user_%v", p.Email)
	redisVal := c.redis.Get(context.Background(), registerUser)
	if redisVal.Err() != nil {
		return nil, redisVal.Err()
	}

	// Unmarshal user data
	var userData DOM.User
	err := json.Unmarshal([]byte(redisVal.Val()), &userData)
	if err != nil {
		return nil, err
	}

	// Verify OTP
	code := fmt.Sprintf("%v", p.OTP)
	phone := strconv.Itoa(userData.Phone)
	resp, err := c.twilio.VerifyTwilioOTP(phone, code)
	if err != nil {
		return nil, errors.New("missmatch OTP")
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}

	// Check if the user with the given email exists
	_, err = c.Repo.FindUserByEmail(userData.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user already exists with the same email")
	}

	// Check if the user with the given phone exists
	_, err = c.Repo.FindUserByPhone(userData.Phone)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user already exists with the same phone")
	}

	// Create the user
	err = c.Repo.CreateUser(&userData)
	if err != nil {
		return nil, errors.New("error while creating user")
	}

	// Return success response
	return &pb.UserResponce{
		Status:  "Success",
		Message: "User creation done",
		Id:      int64(userData.ID),
	}, nil
}

// UserLogin handles the user login process.
func (c *UserSVC) UserLogin(p *pb.UserLogin) (*pb.UserLoginResponce, error) {
	// Find user by email
	user, err := c.Repo.FindUserByEmail(p.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("unable to login")
	}

	// Check password
	check := utils.CheckPasswordMatch([]byte(user.Password), p.Password)
	if !check {
		return nil, fmt.Errorf("password mismatch for user %v", p.Email)
	}

	// Generate token
	userIdstr := strconv.Itoa(int(user.ID))
	token, rt, err := utils.GenerateToken(p.Email, p.Role, userIdstr, config.LoadConfig().SECRETKEY)
	if err != nil {
		return nil, errors.New("error while generating token")
	}

	// Return success response
	return &pb.UserLoginResponce{
		Token:        token,
		RefreshToken: rt,
	}, nil
}
