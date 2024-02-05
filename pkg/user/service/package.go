package service

import (
	"context"
	"fmt"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/pb"
)

func (c *UserSVC) ViewPackageSvc(p *pb.UserView) (*pb.UserPackage, error) {
	var ctxt = context.Background()
	result, err := c.codClient.CoordinatorViewPackage(ctxt, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		fmt.Println("fetching available packages error")
		return nil, err
	} else {
		var dsts []*pb.UserDestination
		for _, pakg := range result.Destinations {
			var dst pb.UserDestination
			dst.Description = pakg.Description
			dst.DestinationId = pakg.DestinationId
			dst.DestinationName = pakg.DestinationName
			dst.Image = pakg.Image
			dst.MaxCapacity = pakg.MaxCapacity
			dst.Minprice = int64(pakg.Minprice)
			dsts = append(dsts, &dst)

		}
		var catogry = &pb.UserCategory{
			CategoryName: result.Category.CategoryName,
		}
		return &pb.UserPackage{
			Packagename:      result.Packagename,
			Startlocation:    result.Startlocation,
			Startdate:        result.Startdate,
			Starttime:        result.Starttime,
			Enddate:          result.Enddate,
			Price:            result.Price,
			Image:            result.Image,
			AvailableSpace:   result.AvailableSpace,
			DestinationCount: result.DestinationCount,
			Destination:      result.Destination,
			PackageId:        result.PackageId,
			Description:      result.Description,
			Category:         catogry,
			Destinations:     dsts,
		}, nil
	}
}

func (c *UserSVC) ViewPackagesSvc(p *pb.UserView) (*pb.UserPackages, error) {
	var ctxt = context.Background()
	result, err := c.codClient.AvailablePackages(ctxt, &cpb.View{
		Status: p.Status,
		Page:   p.Page,
	})
	var pkgs []*pb.UserPackage
	if err != nil {
		fmt.Println("fetching available packages error")
		return &pb.UserPackages{
			Packages: nil,
		}, err
	} else {
		for _, pakg := range result.Packages {
			var pkg pb.UserPackage
			pkg.PackageId = pakg.PackageId
			pkg.Destination = pakg.Destination
			pkg.DestinationCount = int64(pakg.DestinationCount)
			pkg.Enddate = pakg.Enddate
			pkg.Image = pakg.Image
			pkg.Packagename = pakg.Packagename
			pkg.AvailableSpace = pakg.AvailableSpace
			pkg.Price = int64(pakg.Price)
			pkg.Startdate = pakg.Startdate
			pkg.Starttime = pakg.Starttime
			pkg.Startlocation = pakg.Startlocation
			pkg.Description = pakg.Description
			pkgs = append(pkgs, &pkg)

		}
	}

	return &pb.UserPackages{
		Packages: pkgs,
	}, nil
}

func (c *UserSVC) ViewCatagoriesSvc(p *pb.UserView) (*pb.UserCategories, error) {
	var ctxt = context.Background()
	result, err := c.codClient.ViewCatagories(ctxt, &cpb.View{
		Page: p.Page,
	})
	if err != nil {
		fmt.Println("fetching available packages error")
		return nil, err
	}

	var ctgries []*pb.UserCategory

	for _, ctgr := range result.Catagories {
		var ctgry pb.UserCategory
		ctgry.CategoryId = ctgr.CatagoryId
		ctgry.CategoryName = ctgr.CategoryName
		ctgries = append(ctgries, &ctgry)
	}

	return &pb.UserCategories{
		Catagory: ctgries,
	}, nil
}

func (c *UserSVC) FilterPackageSvc(p *pb.UserFilter) (*pb.UserPackages, error) {
	var ctxt = context.Background()
	result, err := c.codClient.FilterPackage(ctxt, &cpb.Filter{
		Status:       p.Status,
		Page:         p.Page,
		Departurtime: p.Departurtime,
		MinPrice:     p.MinPrice,
		MaxPrice:     p.MaxPrice,
		CategoryId:   p.CategoryId,
		OrderBy:      p.OrderBy,
	})
	var pkgs []*pb.UserPackage
	if err != nil {
		fmt.Println("fetching available packages error")
		return &pb.UserPackages{
			Packages: nil,
		}, err
	} else {
		for _, pakg := range result.Packages {
			var pkg pb.UserPackage
			pkg.PackageId = pakg.PackageId
			pkg.Destination = pakg.Destination
			pkg.DestinationCount = int64(pakg.DestinationCount)
			pkg.Enddate = pakg.Enddate
			pkg.Image = pakg.Image
			pkg.Packagename = pakg.Packagename
			pkg.AvailableSpace = pakg.AvailableSpace
			pkg.Price = int64(pakg.Price)
			pkg.Startdate = pakg.Startdate
			pkg.Starttime = pakg.Starttime
			pkg.Startlocation = pakg.Startlocation
			pkg.Description = pakg.Description
			pkgs = append(pkgs, &pkg)

		}
	}

	return &pb.UserPackages{
		Packages: pkgs,
	}, nil
}

func (c *UserSVC) ViewFoodMenusSvc(p *pb.UserView) (*pb.UserFoodMenus, error) {
	var ctx = context.Background()
	rslt, err := c.codClient.CoordinatorViewFoodMenu(ctx, &cpb.View{
		Id:   p.Id,
		Page: p.Page,
	})
	if err != nil {
		fmt.Println("fetching food menu error")
		return &pb.UserFoodMenus{}, err
	}

	var foodmenus []*pb.UserFoodMenu
	for _, menu := range rslt.Foodmenu {
		foodmenus = append(foodmenus, &pb.UserFoodMenu{
			FoodMenuId: menu.FoodMenuId,
			PackageID:  menu.PackageID,
			Breakfast:  menu.Breakfast,
			Lunch:      menu.Lunch,
			Dinner:     menu.Dinner,
			Date:       menu.Date,
		})
	}
	return &pb.UserFoodMenus{
		Foodmenu: foodmenus,
	}, nil
}
