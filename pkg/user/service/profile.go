package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Shakezidin/config"
	pb "github.com/Shakezidin/pkg/user/pb"
	"github.com/Shakezidin/utils"
	"gorm.io/gorm"
)

func (c *UserSVC) ForgetPasswordSVC(p *pb.UserforgetPassword) (*pb.UserResponce, error) {
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
	c.redis.Set(context.Background(), "phoneNo", p.Phone, time.Minute*2)
	return &pb.UserResponce{
		Status:  "success",
		Message: "otp sent successfully",
	}, nil
}

func (c *UserSVC) ForgetPasswordVerifySVC(p *pb.UserforgetPasswordVerify) (*pb.UserResponce, error) {
	redisVal := c.redis.Get(context.Background(), "phoneNo")
	if redisVal.Err() != nil {
		log.Printf("unable to get value from redis err: %v", redisVal.Err().Error())
		return nil, redisVal.Err()
	}
	savedPhone, err := redisVal.Result()
	if err != nil {
		log.Printf("Unable to get saved phone number from Redis err: %v", err.Error())
		return nil, err
	}
	if savedPhone != p.Phone {
		log.Println("Provided phone number does not match the saved phone number.")
		return nil, errors.New("phone number verification failed")
	}

	resp, err := c.twilio.VerifyTwilioOTP(p.Phone, p.Otp)
	if err != nil {
		return &pb.UserResponce{
			Status:  "failed",
			Message: "OTP verification failed",
		}, err
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}

	phoneInt, _ := strconv.Atoi(p.Phone)

	user, err := c.Repo.FindUserByPhone(phoneInt)
	if err != nil {
		fmt.Println("user not found in this number")
		return nil, errors.New("user not found in this number")
	}

	userid := strconv.Itoa(int(user.ID))
	token, err := utils.GenerateToken(user.Email, user.Role, userid, config.LoadConfig().SECRETKEY)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", user.Email, err.Error())
		return nil, err
	}

	return &pb.UserResponce{
		Status:  "success",
		Message: token,
	}, nil
}

func (c *UserSVC) NewPasswordSVC(p *pb.Usernewpassword) (*pb.UserResponce, error) {
	hashPassword, err := utils.HashPassword(p.Newpassword)
	if err != nil {
		log.Printf("unable to hash password in CoordinatorSvc() - service, err: %v", err.Error())
		return nil, err
	}
	id, _ := strconv.Atoi(p.Id)
	err = c.Repo.UpdatePassword(uint(id), string(hashPassword))
	if err != nil {
		fmt.Println("password updating error")
		return nil, errors.New("password updating error")
	}

	return &pb.UserResponce{
		Status:  "Success",
		Message: "Password update success",
	}, nil
}

func (c *UserSVC) UpdateProfileSVC(p *pb.UserSignup) (*pb.UserResponce, error) {
	user, err := c.Repo.FindUserById(uint(p.Id))
	if err != nil {
		return &pb.UserResponce{
			Status:  "fail",
			Message: "user not found",
		}, err
	}

	if p.Email != "" {
		_, err = c.Repo.FindUserByEmail(p.Email)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Existing email found of a user %v", p.Email)
			return nil, errors.New("user already exists with the same role")

		}
		user.Email = p.Email
	}
	if p.Name != "" {
		user.Name = p.Name
	}
	if p.Phone != "" {
		phone, _ := strconv.Atoi(p.Phone)
		_, err = c.Repo.FindUserByPhone(phone)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Existing phone found of a user %v", p.Email)
			return nil, errors.New("user already exists with the same role")

		}
		user.Phone = phone
	}

	err = c.Repo.UpdateUser(user)
	if err != nil {
		return &pb.UserResponce{
			Status:  "fail",
			Message: "failed to update user profile",
		}, err
	}

	// Return a success response
	return &pb.UserResponce{
		Status:  "success",
		Message: "user profile updated successfully",
	}, nil
}
