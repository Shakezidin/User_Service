package service

import (
	"context"
	"fmt"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/userpb"
)

func (c *UserSVC) ViewDestinationSvc(p *pb.UserView) (*pb.UserDestination, error) {
	var ctxt = context.Background()
	result, err := c.codClient.CoordinatorViewDestination(ctxt, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		fmt.Println("fetching available packages error")
		return nil, err
	} else {
		var actvty []*pb.UserActivity
		for _, act := range result.Activity {
			var actv pb.UserActivity
			actv.Description = act.Description
			actv.ActivityId = act.ActivityId
			actv.Activityname = act.Activityname
			actv.Amount = act.Amount
			actv.Date = actv.Date
			actv.DestinationId = actv.DestinationId
			actv.Location = actv.Location
			actv.Time = actv.Time

			actvty = append(actvty, &actv)

		}

		return &pb.UserDestination{
			DestinationName: result.DestinationName,
			DestinationId:   result.DestinationId,
			Description:     result.Description,
			MaxCapacity:     result.MaxCapacity,
			PackageID:       result.PackageID,
			Minprice:        result.Minprice,
			Image:           result.Image,
			Activity:        actvty,
		}, nil
	}
}

func (c *UserSVC) ViewActivitySvc(p *pb.UserView) (*pb.UserActivity, error) {
	var ctxt = context.Background()
	result, err := c.codClient.CoordinatorViewActivity(ctxt, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		fmt.Println("fetching activity error")
		return nil, err
	}

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
