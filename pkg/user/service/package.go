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
		var dst pb.UserDestination
		var dsts []*pb.UserDestination
		for _, pakg := range result.Destinations {
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
			Startdatetime:    result.Startdatetime,
			Enddatetime:      result.Enddatetime,
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
