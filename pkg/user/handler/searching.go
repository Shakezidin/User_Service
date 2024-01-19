package handler

import (
	"context"
	"log"

	pb "github.com/Shakezidin/pkg/user/pb"
)

func (c *UserHandler) UserSearchPacakge(ctx context.Context, p *pb.UserSearch) (*pb.UserPacakges, error) {
	resp, err := c.SVC.SearchPackageSvc(p)
	if err != nil {
		log.Printf("fetching package error", err.Error())
		return nil, err
	}
	return resp, nil
}
