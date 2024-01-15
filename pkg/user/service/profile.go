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
)

func (c *UserSVC) ForgetPassword(p *pb.ForgetPassword) (*pb.ForgetPasswordResponce, error) {
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
	return &pb.ForgetPasswordResponce{
		Responce: "otp sent success",
	}, nil
}

func (c *UserSVC) ForgetPasswordVerify(p *pb.ForgetPasswordVerify) (*pb.ForgetPasswordVerifyResponce, error) {
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
		return nil, err
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

	return &pb.ForgetPasswordVerifyResponce{
		Status: "Now you can create new password",
		Token:  token,
	}, nil
}

func (c *UserSVC) NewPassword(p *pb.Newpassword) (*pb.Newpasswordresponce, error) {
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

	return &pb.Newpasswordresponce{
		Status: "Password update success",
	}, nil
}
