package service

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/Shakezidin/config"
	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/Shakezidin/utils"
	"gorm.io/gorm"
)

// ForgetPasswordSVC initiates the process of forgetting the password by sending an OTP to the user's phone.
func (c *UserSVC) ForgetPasswordSVC(p *pb.UserforgetPassword) (*pb.UserResponce, error) {
	resp, err := c.twilio.SendTwilioOTP(p.Phone)
	if err != nil {
		return &pb.UserResponce{
			Status:  "fail",
			Message: "error while sending OTP",
		}, errors.New("error while sending OTP")
	} else {
		if resp.Status != nil {
			log.Println(*resp.Status)
		} else {
			log.Println(resp.Status)
		}
	}

	c.redis.Set(context.Background(), "phoneNo", p.Phone, time.Minute*2)
	return &pb.UserResponce{
		Status: "success",
	}, nil
}

// ForgetPasswordVerifySVC verifies the OTP provided by the user during the forget password process.
func (c *UserSVC) ForgetPasswordVerifySVC(p *pb.UserforgetPasswordVerify) (*pb.UserResponce, error) {
	redisVal := c.redis.Get(context.Background(), "phoneNo")
	if redisVal.Err() != nil {
		return nil, redisVal.Err()
	}
	savedPhone, err := redisVal.Result()
	if err != nil {
		return &pb.UserResponce{
			Status:  "fail",
			Message: "error while getting data from redis",
		}, err
	}
	if savedPhone != p.Phone {
		return &pb.UserResponce{
			Status:  "fail",
			Message: "provided phone number does not match the saved phone number",
		}, errors.New("phone number verification failed")
	}

	resp, err := c.twilio.VerifyTwilioOTP(p.Phone, p.Otp)
	if err != nil {
		return &pb.UserResponce{
			Status:  "fail",
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
		return nil, errors.New("user not found for this phone number")
	}

	userID := strconv.Itoa(int(user.ID))
	token, _, err := utils.GenerateToken(user.Email, user.Role, userID, config.LoadConfig().SECRETKEY)
	if err != nil {
		return nil, errors.New("error while generating JWT")
	}

	return &pb.UserResponce{
		Status:  "success",
		Message: token,
		Id:      int64(user.ID),
	}, nil
}

// NewPasswordSVC sets a new password for the user after verifying the OTP.
func (c *UserSVC) NewPasswordSVC(p *pb.Usernewpassword) (*pb.UserResponce, error) {
	hashPassword, err := utils.HashPassword(p.Newpassword)
	if err != nil {
		return nil, errors.New("error in password mismatch")
	}
	id, _ := strconv.Atoi(p.Id)
	user, _ := c.Repo.FindUserById(uint(id))
	user.Password = string(hashPassword)
	err = c.Repo.UpdateUser(user)
	if err != nil {
		return nil, errors.New("password updating error")
	}

	return &pb.UserResponce{
		Status:  "Success",
		Message: "Password update success",
		Id:      int64(id),
	}, nil
}

// UpdateProfileSVC updates the profile of the user.
func (c *UserSVC) UpdateProfileSVC(p *pb.UserSignup) (*pb.UserResponce, error) {
	user, err := c.Repo.FindUserById(uint(p.Id))
	if err != nil {
		return &pb.UserResponce{
			Status:  "fail",
			Message: "user not found",
		}, errors.New("user not found")
	}

	if p.Email != "" {
		_, err = c.Repo.FindUserByEmail(p.Email)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user already exists with the same email")
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
			return nil, errors.New("user already exists with the same phone")
		}
		user.Phone = phone
	}

	err = c.Repo.UpdateUser(user)
	if err != nil {
		return &pb.UserResponce{
			Status:  "fail",
			Message: "failed to update user profile",
		}, errors.New("error while updating user profile")
	}

	// Return a success response
	return &pb.UserResponce{
		Status:  "success",
		Message: "user profile updated successfully",
		Id:      int64(user.ID),
	}, nil
}

// UserCountSVC retrieves the count of users.
func (c *UserSVC) UserCountSVC(p *pb.UserView) (*pb.Usercount, error) {
	count := c.Repo.UserCount()

	return &pb.Usercount{
		Usercount: count,
	}, nil
}

// ViewUsersSVC retrieves a list of users based on pagination.
func (c *UserSVC) ViewUsersSVC(p *pb.UserView) (*pb.UserUsers, error) {
	offset := 10 * (p.Page - 1)
	limit := 10

	user, err := c.Repo.FindUsers(int(offset), limit)
	if err != nil {
		return nil, errors.New("no users found")
	}
	var users []*pb.UserSignup

	for _, usr := range user {
		phone := strconv.Itoa(usr.Phone)
		users = append(users, &pb.UserSignup{
			Name:  usr.Name,
			Email: usr.Email,
			Phone: phone,
			Role:  usr.Role,
			Id:    int64(usr.ID),
		})
	}

	return &pb.UserUsers{
		Users: users,
	}, nil
}

// ViewUserSVC retrieves details of a specific user.
func (c *UserSVC) ViewUserSVC(p *pb.UserView) (*pb.UserSignup, error) {
	user, err := c.Repo.FindUserById(uint(p.Id))
	if err != nil {
		return nil, errors.New("error while finding user")
	}

	phone := strconv.Itoa(user.Phone)
	return &pb.UserSignup{
		Name:  user.Name,
		Email: user.Email,
		Phone: phone,
		Role:  user.Role,
		Id:    int64(user.ID),
	}, nil
}
