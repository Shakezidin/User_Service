package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/userpb"
)

// ViewDestinationSvc handles the user request to view destination.
func (c *UserSVC) ViewDestinationSvc(p *pb.UserView) (*pb.UserDestination, error) {
	// Create a new context
	ctx := context.Background()

	// Call Coordinator service to get destination details
	result, err := c.codClient.CoordinatorViewDestination(ctx, &cpb.View{
		ID: p.ID,
	})
	if err != nil {
		return nil, err
	}

	// Map Coordinator response to UserDestination
	var activities []*pb.UserActivity
	for _, act := range result.Activity {
		activity := pb.UserActivity{
			Description:    act.Description,
			Activity_ID:    act.Activity_ID,
			Activity_Name:  act.Activity_Name,
			Amount:         act.Amount,
			Date:           act.Date,
			Destination_ID: act.Destination_ID,
			Location:       act.Location,
			Time:           act.Time,
		}
		activities = append(activities, &activity)
	}

	// Create and return UserDestination response
	return &pb.UserDestination{
		Destination_Name: result.Destination_Name,
		Destination_ID:   result.Destination_ID,
		Description:      result.Description,
		Max_Capacity:     result.Max_Capacity,
		Package_ID:       result.Package_ID,
		Min_Price:        result.Min_Price,
		Image:            result.Image,
		Activity:         activities,
	}, nil
}

// ViewActivitySvc handles the user request to view activity.
func (c *UserSVC) ViewActivitySvc(p *pb.UserView) (*pb.UserActivity, error) {
	// Create a new context
	ctx := context.Background()

	// Call Coordinator service to get activity details
	result, err := c.codClient.CoordinatorViewActivity(ctx, &cpb.View{
		ID: p.ID,
	})
	if err != nil {
		return nil, err
	}

	// Create and return UserActivity response
	return &pb.UserActivity{
		Activity_ID:    result.Activity_ID,
		Activity_Name:  result.Activity_Name,
		Description:    result.Description,
		Location:       result.Location,
		Activity_Type:  result.Activity_Type,
		Amount:         result.Amount,
		Date:           result.Date,
		Time:           result.Time,
		Destination_ID: result.Destination_ID,
	}, nil
}
