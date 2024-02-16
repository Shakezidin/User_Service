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
		Id: p.ID,
	})
	if err != nil {
		return nil, err
	}

	var destinations []*pb.UserDestination
	for _, pkg := range result.Destinations {
		destinations = append(destinations, &pb.UserDestination{
			Description:      pkg.Description,
			Destination_ID:   pkg.DestinationId,
			Destination_Name: pkg.DestinationName,
			Image:            pkg.Image,
			Max_Capacity:     pkg.MaxCapacity,
			Min_Price:        int64(pkg.Minprice),
		})
	}

	category := &pb.UserCategory{
		Category_Name: result.Category.CategoryName,
	}

	return &pb.UserPackage{
		Package_Name:      result.Packagename,
		Start_Location:    result.Startlocation,
		Start_Date:        result.Startdate,
		Start_Time:        result.Starttime,
		End_Date:          result.Enddate,
		Price:             result.Price,
		Image:             result.Image,
		Available_Space:   result.AvailableSpace,
		Destination_Count: result.DestinationCount,
		Destination:       result.Destination,
		Package_ID:        result.PackageId,
		Description:       result.Description,
		Category:          category,
		Destinations:      destinations,
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
			Package_ID:        pkg.PackageId,
			Destination:       pkg.Destination,
			Destination_Count: int64(pkg.DestinationCount),
			End_Date:          pkg.Enddate,
			Image:             pkg.Image,
			Package_Name:      pkg.Packagename,
			Available_Space:   pkg.AvailableSpace,
			Price:             int64(pkg.Price),
			Start_Date:        pkg.Startdate,
			Start_Time:        pkg.Starttime,
			Start_Location:    pkg.Startlocation,
			Description:       pkg.Description,
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
			Category_ID:   category.CatagoryId,
			Category_Name: category.CategoryName,
		})
	}

	return &pb.UserCategories{
		Categories: categories,
	}, nil
}

// FilterPackageSvc filters available packages based on criteria.
func (c *UserSVC) FilterPackageSvc(p *pb.UserFilter) (*pb.UserPackages, error) {
	ctx := context.Background()
	result, err := c.codClient.FilterPackage(ctx, &cpb.Filter{
		Status:       p.Status,
		Page:         p.Page,
		Departurtime: p.Departure_Time,
		MinPrice:     p.Min_Price,
		MaxPrice:     p.Max_Price,
		CategoryId:   p.Category_ID,
		OrderBy:      p.Order_By,
	})
	if err != nil {
		return nil, err
	}

	var packages []*pb.UserPackage
	for _, pkg := range result.Packages {
		packages = append(packages, &pb.UserPackage{
			Package_ID:        pkg.PackageId,
			Destination:       pkg.Destination,
			Destination_Count: int64(pkg.DestinationCount),
			End_Date:          pkg.Enddate,
			Image:             pkg.Image,
			Package_Name:      pkg.Packagename,
			Available_Space:   pkg.AvailableSpace,
			Price:             int64(pkg.Price),
			Start_Date:        pkg.Startdate,
			Start_Time:        pkg.Starttime,
			Start_Location:    pkg.Startlocation,
			Description:       pkg.Description,
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
		Id:   p.ID,
		Page: p.Page,
	})
	if err != nil {
		return nil, err
	}

	var foodMenus []*pb.UserFoodMenu
	for _, menu := range result.Foodmenu {
		foodMenus = append(foodMenus, &pb.UserFoodMenu{
			Food_Menu_ID: menu.FoodMenuId,
			Package_ID:   menu.PackageID,
			Breakfast:    menu.Breakfast,
			Lunch:        menu.Lunch,
			Dinner:       menu.Dinner,
			Date:         menu.Date,
		})
	}

	return &pb.UserFoodMenus{
		Food_Menus: foodMenus,
	}, nil
}
