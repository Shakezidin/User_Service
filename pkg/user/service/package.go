package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/userpb"
)

// ViewPackageSvc retrieves details of a specific package.
func (c *UserSVC) ViewPackageSvc(p *pb.UserView) (*pb.UserPackage, error) {
	ctx := context.Background()
	result, err := c.codClient.CoordinatorViewPackage(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return nil, err
	}

	var destinations []*pb.UserDestination
	for _, pkg := range result.Destinations {
		destinations = append(destinations, &pb.UserDestination{
			Description:     pkg.Description,
			DestinationId:   pkg.DestinationId,
			DestinationName: pkg.DestinationName,
			Image:           pkg.Image,
			MaxCapacity:     pkg.MaxCapacity,
			Minprice:        int64(pkg.Minprice),
		})
	}

	category := &pb.UserCategory{
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
		Category:         category,
		Destinations:     destinations,
	}, nil
}

// ViewPackagesSvc retrieves available packages.
func (c *UserSVC) ViewPackagesSvc(p *pb.UserView) (*pb.UserPackages, error) {
	ctx := context.Background()
	result, err := c.codClient.AvailablePackages(ctx, &cpb.View{
		Status: p.Status,
		Page:   p.Page,
	})
	if err != nil {
		return nil, err
	}

	var packages []*pb.UserPackage
	for _, pkg := range result.Packages {
		packages = append(packages, &pb.UserPackage{
			PackageId:        pkg.PackageId,
			Destination:      pkg.Destination,
			DestinationCount: int64(pkg.DestinationCount),
			Enddate:          pkg.Enddate,
			Image:            pkg.Image,
			Packagename:      pkg.Packagename,
			AvailableSpace:   pkg.AvailableSpace,
			Price:            int64(pkg.Price),
			Startdate:        pkg.Startdate,
			Starttime:        pkg.Starttime,
			Startlocation:    pkg.Startlocation,
			Description:      pkg.Description,
		})
	}

	return &pb.UserPackages{
		Packages: packages,
	}, nil
}

// ViewCatagoriesSvc retrieves available categories.
func (c *UserSVC) ViewCatagoriesSvc(p *pb.UserView) (*pb.UserCategories, error) {
	ctx := context.Background()
	result, err := c.codClient.ViewCatagories(ctx, &cpb.View{
		Page: p.Page,
	})
	if err != nil {
		return nil, err
	}

	var categories []*pb.UserCategory
	for _, category := range result.Catagories {
		categories = append(categories, &pb.UserCategory{
			CategoryId:   category.CatagoryId,
			CategoryName: category.CategoryName,
		})
	}

	return &pb.UserCategories{
		Catagory: categories,
	}, nil
}

// FilterPackageSvc filters available packages based on criteria.
func (c *UserSVC) FilterPackageSvc(p *pb.UserFilter) (*pb.UserPackages, error) {
	ctx := context.Background()
	result, err := c.codClient.FilterPackage(ctx, &cpb.Filter{
		Status:       p.Status,
		Page:         p.Page,
		Departurtime: p.Departurtime,
		MinPrice:     p.MinPrice,
		MaxPrice:     p.MaxPrice,
		CategoryId:   p.CategoryId,
		OrderBy:      p.OrderBy,
	})
	if err != nil {
		return nil, err
	}

	var packages []*pb.UserPackage
	for _, pkg := range result.Packages {
		packages = append(packages, &pb.UserPackage{
			PackageId:        pkg.PackageId,
			Destination:      pkg.Destination,
			DestinationCount: int64(pkg.DestinationCount),
			Enddate:          pkg.Enddate,
			Image:            pkg.Image,
			Packagename:      pkg.Packagename,
			AvailableSpace:   pkg.AvailableSpace,
			Price:            int64(pkg.Price),
			Startdate:        pkg.Startdate,
			Starttime:        pkg.Starttime,
			Startlocation:    pkg.Startlocation,
			Description:      pkg.Description,
		})
	}

	return &pb.UserPackages{
		Packages: packages,
	}, nil
}

// ViewFoodMenusSvc retrieves food menus for a specific package.
func (c *UserSVC) ViewFoodMenusSvc(p *pb.UserView) (*pb.UserFoodMenus, error) {
	ctx := context.Background()
	result, err := c.codClient.CoordinatorViewFoodMenu(ctx, &cpb.View{
		Id:   p.Id,
		Page: p.Page,
	})
	if err != nil {
		return nil, err
	}

	var foodMenus []*pb.UserFoodMenu
	for _, menu := range result.Foodmenu {
		foodMenus = append(foodMenus, &pb.UserFoodMenu{
			FoodMenuId: menu.FoodMenuId,
			PackageID:  menu.PackageID,
			Breakfast:  menu.Breakfast,
			Lunch:      menu.Lunch,
			Dinner:     menu.Dinner,
			Date:       menu.Date,
		})
	}

	return &pb.UserFoodMenus{
		Foodmenu: foodMenus,
	}, nil
}
