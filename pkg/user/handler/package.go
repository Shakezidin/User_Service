package handler

import (
	"log"

	cpb "github.com/Shakezidin/pkg/user/pb"
	"golang.org/x/net/context"
)

func (c *UserHandler) UserViewPackage(ctx context.Context, p *cpb.UserView) (*cpb.UserPackage, error) {
	resp, err := c.SVC.ViewPackageSvc(p)
	if err != nil {
		log.Printf("fetching package error", err.Error())
		return nil, err
	}
	return resp, nil
}
