package handler

import (
	"context"

	pb "github.com/Shakezidin/pkg/user/userpb"
)

func (c *UserHandler) UserTravellerDetails(ctx context.Context, p *pb.UserTravellerRequest) (*pb.UserTravellerResponse, error) {
	resp, err := c.SVC.TraverllerDetailSVC(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
