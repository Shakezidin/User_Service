package handler

import (
	"context"

	pb "github.com/Shakezidin/pkg/user/userpb"
)

func (c *UserHandler) UserRefreshToken(ctx context.Context, p *pb.TokenData) (*pb.UserLoginResponce, error) {
	resp, err := c.SVC.RefreshTokenSVC(p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
