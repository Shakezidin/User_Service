package service

import (
	"context"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/userpb"
)

// OnlinePaymentSVC handles the user request for online payment.
func (c *UserSVC) OnlinePaymentSVC(ctx context.Context, p *pb.UserBooking) (*pb.UserOnlinePaymentResponse, error) {
	// Call Coordinator service to process online payment
	result, err := c.codClient.OnlinePayment(ctx, &cpb.Booking{
		RefId: p.RefId,
		Typ:   p.Typ,
	})
	if err != nil {
		return nil, err
	}

	// Create and return UserOnlinePaymentResponse
	return &pb.UserOnlinePaymentResponse{
		UserId:           result.UserId,
		TotalFare:        result.TotalFare,
		BookingReference: result.BookingReference,
		Email:            result.Email,
		OrderId:          result.OrderId,
	}, nil
}

// PaymentConfirmedSVC handles the user request for confirming payment.
func (c *UserSVC) PaymentConfirmedSVC(ctx context.Context, p *pb.UserPaymentConfirmedRequest) (*pb.UserBookingResponce, error) {
	// Call Coordinator service to confirm payment
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
		return nil, err
	}

	// Create and return UserBookingResponce
	return &pb.UserBookingResponce{
		Status:     result.Status,
		Booking_Id: result.Booking_Id,
	}, nil
}
