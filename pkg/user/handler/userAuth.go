package handler

import (
	"errors"
	"fmt"
	"log"
	"time"

	cpb "github.com/Shakezidin/pkg/user/pb"
	"golang.org/x/net/context"
)

func (c *UserHandler) UserSignupRequest(ctx context.Context, p *cpb.UserSignup) (*cpb.UserResponce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	fmt.Println("hello")
	result, err := c.SVC.SignupSVC(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *UserHandler) UserSignupVerifyRequest(ctx context.Context, p *cpb.UserVerify) (*cpb.UserResponce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := c.SVC.VerifySVC(p)
	if err != nil {
		log.Printf("Unable to verify %v of email == %v, err: %v", p.Email, err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler) UserLoginRequest(ctx context.Context, p *cpb.UserLogin) (*cpb.UserLoginResponce, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	respnc, err := c.SVC.UserLogin(p)
	if err != nil {
		log.Printf("Unable to login %v of email == %v, err: %v", p.Role, p.Email, err.Error())
		return nil, err
	}
	return respnc, nil
}
