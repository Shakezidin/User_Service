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

func (c *UserSVC) SignupSVC(p *pb.UserSignup) (*pb.UserResponce, error) {
	hashPassword, err := utils.HashPassword(p.Password)
	if err != nil {
		log.Printf("unable to hash password in CoordinatorSvc() - service, err: %v", err.Error())
		return nil, err
	}
	phone, _ := strconv.Atoi(p.Phone)
	user := &DOM.User{
		Phone:    phone,
		Email:    p.Email,
		Password: string(hashPassword),
		Name:     p.Name,
		Role:     p.Role,
	}
	// send otp to phone number

	resp, err := c.twilio.SentTwilioOTP(p.Phone)
	if err != nil {
		return nil, err
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}
	userData, err := json.Marshal(&user)
	if err != nil {
		log.Printf("error parsing JSON, err: %v", err.Error())
		return nil, err
	}

	registerUser := fmt.Sprintf("register_user_%v", p.Email)
	c.redis.Set(context.Background(), registerUser, userData, time.Minute*2)
	return &pb.UserResponce{
		Status:  "success",
		Message: "user creation initiated, check mail for OTP",
	}, nil
}

func (c *UserSVC) VerifySVC(p *pb.UserVerify) (*pb.UserResponce, error) {
	registerUser := fmt.Sprintf("register_user_%v", p.Email)
	redisVal := c.redis.Get(context.Background(), registerUser)

	if redisVal.Err() != nil {
		log.Printf("unable to get value from redis err: %v", redisVal.Err().Error())
		return nil, redisVal.Err()
	}

	var userData DOM.User
	err := json.Unmarshal([]byte(redisVal.Val()), &userData)
	if err != nil {
		log.Println("unable to unmarshal json")
		return nil, err
	}

	code := fmt.Sprintf("%v", p.OTP)
	phone := strconv.Itoa(userData.Phone)
	resp, err := c.twilio.VerifyTwilioOTP(phone, code)
	if err != nil {
		return nil, err
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
		log.Printf("Existing email found of a user %v", p.Email)
		return nil, errors.New("user already exists with the same role")

	}

	// Check if the user with the given phone exists
	_, err = c.Repo.FindUserByPhone(userData.Phone)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Existing phone found of a user %v", p.Email)

		// Check if the roles match
		return nil, errors.New("user already exists with the same role")

	}

	err = c.Repo.CreateUser(&userData)
	if err != nil {
		return nil, errors.New("error while creating user")
	}
	return &pb.UserResponce{
		Status:  "Success",
		Message: "User creation done",
		Id:      int64(userData.ID),
	}, nil

}

func (c *UserSVC) UserLogin(p *pb.UserLogin) (*pb.UserLoginResponce, error) {
	user, err := c.Repo.FindUserByEmail(p.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", p.Email)
			return nil, errors.New("user not found")
		} else {
			log.Printf("unable to login %v, err: %v", p.Email, err.Error())
			return nil, errors.New("unable to login")
		}
	}

	check := utils.CheckPasswordMatch([]byte(user.Password), p.Password)
	if !check {
		log.Printf("password mismatch for user %v", p.Email)
		return nil, fmt.Errorf("password mismatch for user %v", p.Email)
	}
	userIdstr := strconv.Itoa(int(user.ID))
	token, rt, err := utils.GenerateToken(p.Email, p.Role, userIdstr, config.LoadConfig().SECRETKEY)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", p.Email, err.Error())
		return nil, err
	}

	return &pb.UserLoginResponce{
		Token:        token,
		RefreshToken: rt,
	}, nil

}
