package handler

import (
	"context"

	pb "github.com/Shakezidin/pkg/user/userpb"
)

func (c *UserHandler) UserViewDestination(ctx context.Context, p *pb.UserView) (*pb.UserDestination, error) {
	resp, err := c.SVC.ViewDestinationSvc(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UserViewActivity(ctx context.Context, p *pb.UserView) (*pb.UserActivity, error) {
	resp, err := c.SVC.ViewActivitySvc(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
