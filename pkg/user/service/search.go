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
		CatagoryId:       p.Category_ID,
		Travelercount:    p.Traveler_Count,
		PickupPlace:      p.Pickup_Place,
		Finaldestination: p.Final_Destination,
		Date:             p.Date,
		Page:             p.Page,
		Enddate:          p.End_Date,
		MaxDestination:   p.Max_Destination,
		Destination:      destinations,
	})
	if err != nil {
		return nil, err
	}

	// Process the retrieved packages
	var pkgs []*pb.UserPackage
	for _, pakg := range result.Packages {
		pkgs = append(pkgs, &pb.UserPackage{
			Package_ID:        pakg.PackageId,
			Destination:       pakg.Destination,
			Destination_Count: int64(pakg.DestinationCount),
			End_Date:          pakg.Enddate,
			Image:             pakg.Image,
			Package_Name:      pakg.Packagename,
			Available_Space:   pakg.AvailableSpace,
			Price:             int64(pakg.Price),
			Start_Date:        pakg.Startdate,
			Start_Time:        pakg.Starttime,
			Start_Location:    pakg.Startlocation,
			Description:       pakg.Description,
			Coordinator_ID:    pakg.CoorinatorId,
		})
	}

	// Return the packages
	return &pb.UserPackages{
		Packages: pkgs,
	}, nil
}
