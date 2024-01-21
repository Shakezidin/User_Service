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
			Endlocation:      result.Endlocation,
			Startdate:        result.Startdate,
			Enddate:          result.Enddate,
			Price:            result.Price,
			Image:            result.Image,
			DestinationCount: result.DestinationCount,
			Destination:      result.Destination,
			PackageId:        result.PackageId,
			Description:      result.Description,
			Category:         catogry,
			Destinations:     dsts,
		}, nil
	}
}

func (c *UserSVC) ViewCatagoriesSvc() (*pb.UserCategories, error) {
	var ctxt = context.Background()
	result, err := c.codClient.ViewCatagories(ctxt, &cpb.View{})
	if err != nil {
		fmt.Println("fetching available packages error")
		return nil, err
	}

	var ctgry pb.UserCategory
	var ctgries []*pb.UserCategory

	for _, ctgr := range result.Catagories {
		ctgry.CategoryName = ctgr.CategoryName
		ctgry.CategoryId = ctgr.CatagoryId
		ctgries = append(ctgries, &ctgry)
	}

	return &pb.UserCategories{
		Catagory: ctgries,
	}, nil
}
