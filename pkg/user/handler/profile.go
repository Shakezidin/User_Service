package handler

import (
	"errors"
	"log"
	"time"

	cpb "github.com/Shakezidin/pkg/user/pb"
	"golang.org/x/net/context"
)

func (c *UserHandler) ForgetPassword(ctx context.Context, p *cpb.UserforgetPassword) (*cpb.UserResponce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := c.SVC.ForgetPasswordSVC(p)
	if err != nil {
		log.Printf("Unable to verify sent of otp for phone == %v, err: %v", p.Phone, err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler) ForgetPasswordVerify(ctx context.Context, p *cpb.UserforgetPasswordVerify) (*cpb.UserResponce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := c.SVC.ForgetPasswordVerifySVC(p)
	if err != nil {
		log.Printf("Unable to verify otp for phone == %v, err: %v", p.Phone, err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler) NewPassword(ctx context.Context, p *cpb.Usernewpassword) (*cpb.UserResponce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := c.SVC.NewPasswordSVC(p)
	if err != nil {
		log.Printf("Unable to update password, err: %v", err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler)ProfileUpdate(ctx context.Context,p *cpb.UserSignup)(*cpb.UserResponce,error){
	resp, err := c.SVC.UpdateProfileSVC(p)
	if err != nil {
		log.Printf("Unable to update profile, err: %v", err.Error())
		return nil, err
	}
	return resp, nil
}
