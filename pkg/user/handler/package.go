package handler

import (
	pb "github.com/Shakezidin/pkg/user/userpb"
	"golang.org/x/net/context"
)

func (c *UserHandler) UserViewPackage(ctx context.Context, p *pb.UserView) (*pb.UserPackage, error) {
	resp, err := c.SVC.ViewPackageSvc(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UserViewCategories(ctx context.Context, p *pb.UserView) (*pb.UserCategories, error) {
	resp, err := c.SVC.ViewCatagoriesSvc(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UserViewPackages(ctx context.Context, p *pb.UserView) (*pb.UserPackages, error) {
	resp, err := c.SVC.ViewPackagesSvc(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UserFilterPackage(ctx context.Context, p *pb.UserFilter) (*pb.UserPackages, error) {
	resp, err := c.SVC.FilterPackageSvc(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UserViewFoodMenu(ctx context.Context, p *pb.UserView) (*pb.UserFoodMenus, error) {
	resp, err := c.SVC.ViewFoodMenusSvc(p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
