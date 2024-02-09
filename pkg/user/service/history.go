package service

import (
	"context"
	"errors"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/userpb"
)

func (c *UserSVC) ViewHistorySvc(p *pb.UserView) (*pb.UserHistories, error) {
	var ctx = context.Background()
	result, err := c.codClient.ViewHistory(ctx, &cpb.View{
		Id:     p.Id,
		Page:   p.Page,
		Status: "false",
	})

	if err != nil {
		return &pb.UserHistories{}, err
	}

	var history []*pb.UserHistory
	for _, hstry := range result.History {
		history = append(history, &pb.UserHistory{
			Id:              hstry.Id,
			PaymentMode:     hstry.PaymentMode,
			BookingStatus:   hstry.BookingStatus,
			CancelledStatus: hstry.CancelledStatus,
			TotalPrice:      int64(hstry.TotalPrice),
			UserId:          int64(hstry.UserId),
			BookingId:       hstry.BookingId,
			BookDate:        hstry.BookDate,
			StartDate:       hstry.StartDate,
			PaidAmount:      hstry.PaidAmount,
		})
	}

	return &pb.UserHistories{
		History: history,
	}, nil

}

func (c *UserSVC) ViewBookingSvc(p *pb.UserView) (*pb.UserHistory, error) {
	var ctx = context.Background()
	result, err := c.codClient.ViewBooking(ctx, &cpb.View{
		Id: p.Id,
	})

	if err != nil {
		return &pb.UserHistory{}, errors.New("error while fetching booking")
	}

	var traveller []*pb.UserTravellerDetails
	for _, trvlr := range result.Travellers {
		traveller = append(traveller, &pb.UserTravellerDetails{
			Name:   trvlr.Name,
			Age:    trvlr.Age,
			Gender: trvlr.Gender,
		})
	}

	return &pb.UserHistory{
		Id:              result.Id,
		PaymentMode:     result.PaymentMode,
		BookingStatus:   result.BookingStatus,
		CancelledStatus: result.CancelledStatus,
		TotalPrice:      int64(result.TotalPrice),
		UserId:          int64(result.UserId),
		BookingId:       result.BookDate,
		StartDate:       result.StartDate,
		Travellers:      traveller,
		PaidAmount:      result.PaidAmount,
	}, nil
}

func (c *UserSVC) CancelBookingSvc(p *pb.UserView) (*pb.UserResponce, error) {
	var ctx = context.Background()
	user, err := c.Repo.FindUserById(uint(p.UserId))
	if err != nil {
		return &pb.UserResponce{}, errors.New("user not found in this id")
	}
	result, err := c.codClient.CancelBooking(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return &pb.UserResponce{}, errors.New("error while cancelling booking")
	}

	user.Wallet += int(result.Amount)

	err = c.Repo.UpdateUser(user)
	if err != nil {
		return &pb.UserResponce{}, errors.New("error while updating user wallet")
	}

	return &pb.UserResponce{
		Status:  result.Status,
		Message: result.Message,
	}, nil
}
