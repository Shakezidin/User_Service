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
		Category_ID:       p.Category_ID,
		Traveler_Count:    p.Traveler_Count,
		Pickup_Place:      p.Pickup_Place,
		Final_Destination: p.Final_Destination,
		Date:              p.Date,
		Page:              p.Page,
		End_Date:          p.End_Date,
		Max_Destination:   p.Max_Destination,
		Destination:       destinations,
	})
	if err != nil {
		return nil, err
	}

	// Process the retrieved packages
	var pkgs []*pb.UserPackage
	for _, pakg := range result.Packages {
		pkgs = append(pkgs, &pb.UserPackage{
			Package_ID:        pakg.Package_ID,
			Destination:       pakg.Destination,
			Destination_Count: int64(pakg.Destination_Count),
			End_Date:          pakg.End_Date,
			Image:             pakg.Image,
			Package_Name:      pakg.Packagename,
			Available_Space:   pakg.Available_Space,
			Price:             int64(pakg.Price),
			Start_Date:        pakg.Start_Date,
			Start_Time:        pakg.Start_Time,
			Start_Location:    pakg.Start_Location,
			Description:       pakg.Description,
			Coordinator_ID:    pakg.Coordinator_ID,
		})
	}

	// Return the packages
	return &pb.UserPackages{
		Packages: pkgs,
	}, nil
}
