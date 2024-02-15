package handler

import (
	"context"

	pb "github.com/Shakezidin/pkg/user/userpb"
)

func (c *UserHandler) UserSearchPackage(ctx context.Context, p *pb.UserSearch) (*pb.UserPackages, error) {
	resp, err := c.SVC.SearchPackageSvc(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
