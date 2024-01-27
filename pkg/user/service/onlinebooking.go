package service

import (
	"context"
	"fmt"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/pb"
)

func (c *UserSVC) OnlinePaymentSVC(ctx context.Context, p *pb.UserBooking) (*pb.UserOnlinePaymentResponse, error) {
	result, err := c.codClient.OnlinePayment(ctx, &cpb.Booking{
		RefId:  p.RefId,
		UserId: p.UserId,
	})
	if err != nil {
		fmt.Println("fetching available packages error")
		return &pb.UserOnlinePaymentResponse{}, err
	}

	return &pb.UserOnlinePaymentResponse{
		UserId:           result.UserId,
		TotalFare:        result.TotalFare,
		BookingReference: result.BookingReference,
		Email:            result.Email,
		OrderId:          result.OrderId,
	}, nil
}
