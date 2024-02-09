package handler

import (
	"log"

	pb "github.com/Shakezidin/pkg/user/userpb"
	"golang.org/x/net/context"
)

func (c *UserHandler) UserViewPackage(ctx context.Context, p *pb.UserView) (*pb.UserPackage, error) {
	resp, err := c.SVC.ViewPackageSvc(p)
	if err != nil {
		log.Printf("fetching package error", err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler)UserViewCatagories(ctx context.Context, p *pb.UserView)(*pb.UserCategories,error){
	resp, err := c.SVC.ViewCatagoriesSvc(p)
	if err != nil {
		log.Printf("fetching catagories error", err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler)UserViewPackages(ctx context.Context,p *pb.UserView)(*pb.UserPackages,error){
	resp, err := c.SVC.ViewPackagesSvc(p)
	if err != nil {
		log.Printf("fetching packages error", err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler)UserFilterPackage(ctx context.Context,p *pb.UserFilter)(*pb.UserPackages,error){
	resp, err := c.SVC.FilterPackageSvc(p)
	if err != nil {
		log.Printf("fetching packages error", err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler)UserViewFoodMenu(ctx context.Context,p *pb.UserView)(*pb.UserFoodMenus,error){
	resp, err := c.SVC.ViewFoodMenusSvc(p)
	if err != nil {
		log.Printf("fetching food menus error", err.Error())
		return nil, err
	}
	return resp, nil
}