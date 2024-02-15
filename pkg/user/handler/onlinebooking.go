package handler

import (
	"context"

	pb "github.com/Shakezidin/pkg/user/userpb"
)

func (c *UserHandler) UserOnlinePayment(ctx context.Context, p *pb.UserBooking) (*pb.UserOnlinePaymentResponse, error) {
	resp, err := c.SVC.OnlinePaymentSVC(ctx, p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *UserHandler) UserPaymentConfirmed(ctx context.Context, p *pb.UserPaymentConfirmedRequest) (*pb.UserBookingResponce, error) {
	resp, err := c.SVC.PaymentConfirmedSVC(ctx, p)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
