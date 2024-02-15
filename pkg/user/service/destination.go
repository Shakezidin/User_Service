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
		Id: p.Id,
	})
	if err != nil {
		return nil, err
	}

	// Map Coordinator response to UserDestination
	var activities []*pb.UserActivity
	for _, act := range result.Activity {
		activity := pb.UserActivity{
			Description:   act.Description,
			ActivityId:    act.ActivityId,
			Activityname:  act.Activityname,
			Amount:        act.Amount,
			Date:          act.Date,
			DestinationId: act.DestinationId,
			Location:      act.Location,
			Time:          act.Time,
		}
		activities = append(activities, &activity)
	}

	// Create and return UserDestination response
	return &pb.UserDestination{
		DestinationName: result.DestinationName,
		DestinationId:   result.DestinationId,
		Description:     result.Description,
		MaxCapacity:     result.MaxCapacity,
		PackageID:       result.PackageID,
		Minprice:        result.Minprice,
		Image:           result.Image,
		Activity:        activities,
	}, nil
}

// ViewActivitySvc handles the user request to view activity.
func (c *UserSVC) ViewActivitySvc(p *pb.UserView) (*pb.UserActivity, error) {
	// Create a new context
	ctx := context.Background()

	// Call Coordinator service to get activity details
	result, err := c.codClient.CoordinatorViewActivity(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return nil, err
	}

	// Create and return UserActivity response
	return &pb.UserActivity{
		ActivityId:    result.ActivityId,
		Activityname:  result.Activityname,
		Description:   result.Description,
		Location:      result.Location,
		ActivityType:  result.ActivityType,
		Amount:        result.Amount,
		Date:          result.Date,
		Time:          result.Time,
		DestinationId: result.DestinationId,
	}, nil
}
