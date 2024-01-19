package handler

import (
	"context"
	"log"

	pb "github.com/Shakezidin/pkg/user/pb"
)

func (c *UserHandler) UserViewDestination(ctx context.Context, p *pb.UserView) (*pb.UserDestination, error) {
	resp, err := c.SVC.ViewDestinationSvc(p)
	if err != nil {
		log.Printf("fetching destination error", err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler)UserViewActivity(ctx context.Context, p *pb.UserView)(*pb.UserActivity,error){
	resp, err := c.SVC.ViewActivitySvc(p)
	if err != nil {
		log.Printf("fetching activity error", err.Error())
		return nil, err
	}
	return resp, nil
}