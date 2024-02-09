package handler

import (
	"context"
	"log"

	pb "github.com/Shakezidin/pkg/user/userpb"
)

func (c *UserHandler) UserTravellerDetails(ctx context.Context, p *pb.UserTravellerRequest) (*pb.UserTravellerResponse, error) {
	resp, err := c.SVC.TraverllerDetailSVC(p)
	if err != nil {
		log.Printf("error while adding traveller details", err.Error())
		return nil, err
	}
	return resp, nil
}
