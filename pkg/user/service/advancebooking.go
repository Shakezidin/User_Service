package service

import (
	"context"
	"errors"
	"strconv"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/userpb"
)

// TraverllerDetailSVC handles the request to add traveller details.
func (c *UserSVC) TraverllerDetailSVC(p *pb.UserTravellerRequest) (*pb.UserTravellerResponse, error) {
	// Create a background context
	ctx := context.Background()

	// Convert UserId to integer
	userID, err := strconv.Atoi(p.UserId)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	// Find user by ID
	user, err := c.Repo.FindUserById(uint(userID))
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Initialize a slice to store traveller details
	var travellerDetails []*cpb.TravellerDetails

	// Iterate over traveller details and populate the slice
	for _, dtls := range p.TravellerDetails {
		travellerDetail := &cpb.TravellerDetails{
			ActivityId: dtls.ActivityId,
			Age:        dtls.Age,
			Gender:     dtls.Gender,
			Name:       dtls.Name,
		}
		travellerDetails = append(travellerDetails, travellerDetail)
	}

	// Call the client to add traveller details
	resp, err := c.codClient.UserTravellerDetails(ctx, &cpb.TravellerRequest{
		UserId:           p.UserId,
		TravellerDetails: travellerDetails,
		PackageId:        p.PackageId,
		Email:            user.Email,
		Username:         user.Name,
	})

	if err != nil {
		return nil, err
	}

	// Return the response
	return &pb.UserTravellerResponse{
		Status:             resp.Status,
		PackagePrice:       resp.PackagePrice,
		ActivityTotalPrice: resp.ActivityTotalPrice,
		TotalPrice:         resp.TotalPrice,
		AdvanceAmount:      resp.AdvanceAmount,
		RefId:              resp.RefId,
	}, nil
}
