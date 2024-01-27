package handler

import (
	"context"
	"log"

	pb "github.com/Shakezidin/pkg/user/pb"
)

func (c *UserHandler) UserOnlinePayment(ctx context.Context, p *pb.UserBooking) (*pb.UserOnlinePaymentResponse, error) {
	resp, err := c.SVC.OnlinePaymentSVC(ctx,p)
	if err != nil {
		log.Printf("error while adding traveller details", err.Error())
		return nil, err
	}
	return resp, nil
}
