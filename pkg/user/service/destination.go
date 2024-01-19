package service

import (
	"context"
	"fmt"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/pb"
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
			DestinationId:   result.DestinationId,
			DestinationName: result.DestinationName,
			Description:     result.Description,
			Minprice:        result.Minprice,
			MaxCapacity:     result.MaxCapacity,
			Image:           result.Image,
			Activity:        actvty,
			PackageID:       result.PackageID,
		}, nil
	}
}