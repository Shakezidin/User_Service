package service

import (
	"context"
	"fmt"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/pb"
)

func (c *UserSVC) TraverllerDetailSVC(p *pb.UserTravellerRequest) (*pb.UserTravellerResponse, error) {
	var ctx = context.Background()
	var travellerDetails []*cpb.TravellerDetails

	for _, dtls := range p.TravellerDetails {
		var travellerDetail cpb.TravellerDetails
		travellerDetail.ActivityId = dtls.ActivityId
		travellerDetail.Age = dtls.Age
		travellerDetail.Gender = dtls.Gender
		travellerDetail.Name = dtls.Name

		travellerDetails = append(travellerDetails, &travellerDetail)
	}
	resp, err := c.codClient.TravellerDetails(ctx, &cpb.TravellerRequest{
		TravellerDetails: travellerDetails,
		UserId:           p.UserId,
		PackageId:        p.PackageId,
	})

	if err != nil {
		fmt.Println("error while adding traveller details")
		return nil, err
	}

	return &pb.UserTravellerResponse{
		Status:             resp.Status,
		PackagePrice:       resp.PackagePrice,
		ActivityTotalPrice: resp.ActivityTotalPrice,
		TotalPrice:         resp.TotalPrice,
		RefId:              resp.RefId,
	}, nil
}

func (c *UserSVC) OfflineBookingSVC(ctx context.Context, p *pb.UserBooking) (*pb.UserBookingResponce, error) {
	result, err := c.codClient.OfflineBooking(ctx, &cpb.Booking{
		RefId:  p.RefId,
		UserId: p.UserId,
	})
	if err != nil {
		fmt.Println("fetching available packages error")
		return nil, err
	}

	return &pb.UserBookingResponce{
		Status:     result.Status,
		Booking_Id: result.Booking_Id,
	}, nil
}
