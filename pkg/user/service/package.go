package service

import (
	"context"
	"fmt"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/pb"
)

func (c *UserSVC) ViewPackageSvc(p *pb.ViewPackage) (*pb.ViewPacakgeResponce, error) {
	var ctxt = context.Background()
	result, err := c.codClient.CoordinatorViewPackage(ctxt, &cpb.CoodinatorViewPackage{
		PackageId: p.PackageId,
	})
	if err != nil {
		fmt.Println("fetching available packages error")
		return &pb.ViewPacakgeResponce{}, err
	} else {
		var dst pb.UsrDestinations
		var dsts []*pb.UsrDestinations
		for _, pakg := range result.Destinations {
			dst.Description = pakg.Description
			dst.DestinationId = pakg.DestinationId
			dst.DestinationName = pakg.DestinationName
			dst.Image = pakg.Image
			dst.MaxCapacity = pakg.MaxCapacity
			dst.MinPrice = pakg.MinPrice
			dsts = append(dsts, &dst)

		}
		var catogry = &pb.UsrCategory{
			CategoryName: result.Category.CategoryName,
		}
		return &pb.ViewPacakgeResponce{
			Name:             result.Name,
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
