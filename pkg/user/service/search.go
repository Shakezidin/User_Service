package service

import (
	"context"
	"fmt"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/pb"
)

func (c *UserSVC) SearchPackageSvc(p *pb.UserSearch) (*pb.UserPackages, error) {
	var ctxt = context.Background()
	destinations := make([]string, len(p.Destination))
	for _, dstn := range p.Destination {
		destinations = append(destinations, dstn)
	}
	result, err := c.codClient.PackageSearch(ctxt, &cpb.Search{
		CatagoryId:       p.CatagoryId,
		Travelercount:    p.Travelercount,
		PickupPlace:      p.PickupPlace,
		Finaldestination: p.Finaldestination,
		Date:             p.Date,
		Page:             p.Page,
		Enddate:          p.Enddate,
		MaxDestination:   p.MaxDestination,
		Destination:      destinations,
	})
	if err != nil {
		fmt.Println("fetching available packages error")
		return nil, err
	}

	var pkgs []*pb.UserPackage
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
		pkg.Startlocation = pakg.Startlocation
		pkg.Description = pakg.Description
		pkg.CoorinatorId = pakg.CoorinatorId
		pkgs = append(pkgs, &pkg)
	}

	return &pb.UserPackages{
		Packages: pkgs,
	}, nil

}
