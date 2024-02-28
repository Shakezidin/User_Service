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
		ID: p.ID,
	})
	if err != nil {
		return nil, err
	}

	var destinations []*pb.UserDestination
	for _, pkg := range result.Destinations {
		destinations = append(destinations, &pb.UserDestination{
			Description:      pkg.Description,
			Destination_ID:   pkg.Destination_ID,
			Destination_Name: pkg.Destination_Name,
			Image:            pkg.Image,
			Max_Capacity:     pkg.Max_Capacity,
			Min_Price:        int64(pkg.Min_Price),
		})
	}

	category := &pb.UserCategory{
		Category_Name: result.Category.Category_Name,
	}

	return &pb.UserPackage{
		Package_Name:      result.Packagename,
		Start_Location:    result.Start_Location,
		Start_Date:        result.Start_Date,
		Start_Time:        result.Start_Time,
		End_Date:          result.End_Date,
		Price:             result.Price,
		Image:             result.Image,
		Available_Space:   result.Available_Space,
		Destination_Count: result.Destination_Count,
		Destination:       result.Destination,
		Package_ID:        result.Package_ID,
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
			Package_ID:        pkg.Package_ID,
			Destination:       pkg.Destination,
			Destination_Count: int64(pkg.Destination_Count),
			End_Date:          pkg.End_Date,
			Image:             pkg.Image,
			Package_Name:      pkg.Packagename,
			Available_Space:   pkg.Available_Space,
			Price:             int64(pkg.Price),
			Start_Date:        pkg.Start_Date,
			Start_Time:        pkg.Start_Time,
			Start_Location:    pkg.Start_Location,
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
	result, err := c.codClient.Viewcategories(ctx, &cpb.View{
		Page: p.Page,
	})
	if err != nil {
		return nil, err
	}

	var categories []*pb.UserCategory
	for _, category := range result.Categories {
		categories = append(categories, &pb.UserCategory{
			Category_ID:   category.Category_ID,
			Category_Name: category.Category_Name,
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
		Status:         p.Status,
		Page:           p.Page,
		Departure_Time: p.Departure_Time,
		Min_Price:      p.Min_Price,
		Max_Price:      p.Max_Price,
		Category_ID:    p.Category_ID,
		Order_By:       p.Order_By,
	})
	if err != nil {
		return nil, err
	}

	var packages []*pb.UserPackage
	for _, pkg := range result.Packages {
		packages = append(packages, &pb.UserPackage{
			Package_ID:        pkg.Package_ID,
			Destination:       pkg.Destination,
			Destination_Count: int64(pkg.Destination_Count),
			End_Date:          pkg.End_Date,
			Image:             pkg.Image,
			Package_Name:      pkg.Packagename,
			Available_Space:   pkg.Available_Space,
			Price:             int64(pkg.Price),
			Start_Date:        pkg.Start_Date,
			Start_Time:        pkg.Start_Time,
			Start_Location:    pkg.Start_Location,
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
		ID:   p.ID,
		Page: p.Page,
	})
	if err != nil {
		return nil, err
	}

	var foodMenus []*pb.UserFoodMenu
	for _, menu := range result.FoodMenus {
		foodMenus = append(foodMenus, &pb.UserFoodMenu{
			Food_Menu_ID: menu.Food_Menu_ID,
			Package_ID:   menu.Package_ID,
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
