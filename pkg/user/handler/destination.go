package handler

import (
	"context"
	"log"

	pb "github.com/Shakezidin/pkg/user/pb"
)

func (c *UserHandler) UserViewDestination(ctx context.Context, p *pb.UserView) (*pb.UserDestination, error) {
	resp, err := c.SVC.ViewDestinationSvc(p)
	if err != nil {
		log.Printf("fetching package error", err.Error())
		return nil, err
	}
	return resp, nil
}
