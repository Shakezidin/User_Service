package handler

import (
	"context"
	"log"

	pb "github.com/Shakezidin/pkg/user/userpb"
)

func (c *UserHandler) UserViewHistory(ctx context.Context, p *pb.UserView) (*pb.UserHistories, error) {
	resp, err := c.SVC.ViewHistorySvc(p)
	if err != nil {
		log.Printf("fetching history error", err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler) UserViewBooking(ctx context.Context, p *pb.UserView) (*pb.UserHistory, error) {
	resp, err := c.SVC.ViewBookingSvc(p)
	if err != nil {
		log.Printf("fetching Booking error", err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *UserHandler)UserCancelBooking(ctx context.Context,p *pb.UserView)(*pb.UserResponce,error){
	resp, err := c.SVC.CancelBookingSvc(p)
	if err != nil {
		log.Printf("fetching Booking error", err.Error())
		return nil, err
	}
	return resp, nil
}