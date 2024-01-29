package service

import (
	"context"
	"fmt"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/pb"
)

func (c *UserSVC) OnlinePaymentSVC(ctx context.Context, p *pb.UserBooking) (*pb.UserOnlinePaymentResponse, error) {
	result, err := c.codClient.OnlinePayment(ctx, &cpb.Booking{
		RefId: p.RefId,
		Typ:   p.Typ,
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

func (c *UserSVC) PaymentConfirmedSVC(ctx context.Context, p *pb.UserPaymentConfirmedRequest) (*pb.UserBookingResponce, error) {
	result, err := c.codClient.PaymentConfirmed(ctx, &cpb.PaymentConfirmedRequest{
		Email:       p.Email,
		ReferenceID: p.ReferenceID,
		PaymentId:   p.PaymentId,
		UserID:      p.UserID,
		OrderID:     p.OrderID,
		Signature:   p.Signature,
		Total:       p.Total,
	})
	if err != nil {
		fmt.Println("fetching available packages error")
		return &pb.UserBookingResponce{}, err
	}

	return &pb.UserBookingResponce{
		Status:     result.Status,
		Booking_Id: result.Booking_Id,
	}, nil
}
