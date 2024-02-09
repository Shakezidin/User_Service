package handler

import (
	"errors"
	"log"
	"time"

	cpb "github.com/Shakezidin/pkg/user/userpb"
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

func (c *UserHandler) ProfileUpdate(ctx context.Context, p *cpb.UserSignup) (*cpb.UserResponce, error) {
	resp, err := c.SVC.UpdateProfileSVC(p)
	if err != nil {
		log.Printf("Unable to update profile, err: %v", err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler) UsersCount(ctx context.Context, p *cpb.UserView) (*cpb.Usercount, error) {
	resp, err := c.SVC.UserCountSVC(p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler) UserViewUsers(ctx context.Context, p *cpb.UserView) (*cpb.UserUsers, error) {
	resp, err := c.SVC.ViewUsersSVC(p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler) UserViewUser(ctx context.Context, p *cpb.UserView) (*cpb.UserSignup, error) {
	resp, err := c.SVC.ViewUserSVC(p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
