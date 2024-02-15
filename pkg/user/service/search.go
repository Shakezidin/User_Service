package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/userpb"
)

// SearchPackageSvc searches for packages based on the provided search parameters.
func (c *UserSVC) SearchPackageSvc(p *pb.UserSearch) (*pb.UserPackages, error) {
	// Create a context
	ctx := context.Background()

	destinations := make([]string, len(p.Destination))
	copy(destinations, p.Destination)

	// Call the coordinator service to search for packages
	result, err := c.codClient.PackageSearch(ctx, &cpb.Search{
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
		return nil, err
	}

	// Process the retrieved packages
	var pkgs []*pb.UserPackage
	for _, pakg := range result.Packages {
		pkgs = append(pkgs, &pb.UserPackage{
			PackageId:        pakg.PackageId,
			Destination:      pakg.Destination,
			DestinationCount: int64(pakg.DestinationCount),
			Enddate:          pakg.Enddate,
			Image:            pakg.Image,
			Packagename:      pakg.Packagename,
			AvailableSpace:   pakg.AvailableSpace,
			Price:            int64(pakg.Price),
			Startdate:        pakg.Startdate,
			Starttime:        pakg.Starttime,
			Startlocation:    pakg.Startlocation,
			Description:      pakg.Description,
			CoorinatorId:     pakg.CoorinatorId,
		})
	}

	// Return the packages
	return &pb.UserPackages{
		Packages: pkgs,
	}, nil
}
