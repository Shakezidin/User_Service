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
		Ref_ID: p.Ref_ID,
		Typ:    p.Typ,
	})
	if err != nil {
		return nil, err
	}

	// Create and return UserOnlinePaymentResponse
	return &pb.UserOnlinePaymentResponse{
		User_ID:           result.User_ID,
		Total_Fare:        result.Total_Fare,
		Booking_Reference: result.Booking_Reference,
		Email:             result.Email,
		Order_ID:          result.Order_ID,
	}, nil
}

// PaymentConfirmedSVC handles the user request for confirming payment.
func (c *UserSVC) PaymentConfirmedSVC(ctx context.Context, p *pb.UserPaymentConfirmedRequest) (*pb.UserBookingResponse, error) {
	// Call Coordinator service to confirm payment
	result, err := c.codClient.PaymentConfirmed(ctx, &cpb.PaymentConfirmedRequest{
		Email:        p.Email,
		Reference_ID: p.Reference_ID,
		Payment_ID:   p.Payment_ID,
		User_ID:      p.User_ID,
		Order_ID:     p.Order_ID,
		Signature:    p.Signature,
		Total:        p.Total,
	})
	if err != nil {
		return nil, err
	}

	// Create and return UserBookingResponce
	return &pb.UserBookingResponse{
		Status:     result.Status,
		Booking_ID: result.Booking_ID,
	}, nil
}
