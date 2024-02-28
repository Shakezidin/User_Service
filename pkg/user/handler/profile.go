package handler

import (
	"errors"
	"time"

	cpb "github.com/Shakezidin/pkg/user/userpb"
	"golang.org/x/net/context"
)

func (c *UserHandler) UserForgetPassword(ctx context.Context, p *cpb.UserForget_Password) (*cpb.UserResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := c.SVC.ForgetPasswordSVC(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UserForgetPasswordVerify(ctx context.Context, p *cpb.UserForget_PasswordVerify) (*cpb.UserResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := c.SVC.ForgetPasswordVerifySVC(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UserNewPassword(ctx context.Context, p *cpb.UserNew_Password) (*cpb.UserResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := c.SVC.NewPasswordSVC(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UserProfileUpdate(ctx context.Context, p *cpb.UserSignup) (*cpb.UserResponse, error) {
	resp, err := c.SVC.UpdateProfileSVC(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UsersCount(ctx context.Context, p *cpb.UserView) (*cpb.UserCount, error) {
	resp, err := c.SVC.UserCountSVC(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UserViewUsers(ctx context.Context, p *cpb.UserView) (*cpb.UserUsers, error) {
	resp, err := c.SVC.ViewUsersSVC(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UserViewUser(ctx context.Context, p *cpb.UserView) (*cpb.UserSignup, error) {
	resp, err := c.SVC.ViewUserSVC(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
