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
	userID, err := strconv.Atoi(p.User_ID)
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
	for _, dtls := range p.Traveller_Details {
		travellerDetail := &cpb.TravellerDetails{
			Activity_ID: dtls.Activity_ID,
			Age:         dtls.Age,
			Gender:      dtls.Gender,
			Name:        dtls.Name,
		}
		travellerDetails = append(travellerDetails, travellerDetail)
	}

	// Call the client to add traveller details
	resp, err := c.codClient.UserTravellerDetails(ctx, &cpb.TravellerRequest{
		User_ID:           p.User_ID,
		Traveller_Details: travellerDetails,
		Package_ID:        p.Package_ID,
		Email:             user.Email,
		Username:          user.Name,
	})

	if err != nil {
		return nil, err
	}

	// Return the response
	return &pb.UserTravellerResponse{
		Status:               resp.Status,
		Package_Price:        resp.Package_Price,
		Activity_Total_Price: resp.Activity_Total_Price,
		Total_Price:          resp.Total_Price,
		Advance_Amount:       resp.Advance_Amount,
		Ref_ID:               resp.Ref_ID,
	}, nil
}
