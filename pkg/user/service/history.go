package service

import (
	"context"
	"errors"

	cpb "github.com/Shakezidin/pkg/user/client/pb"
	pb "github.com/Shakezidin/pkg/user/userpb"
)

// ViewHistorySvc handles the user request to view booking history.
func (c *UserSVC) ViewHistorySvc(p *pb.UserView) (*pb.UserHistories, error) {
	// Create a new context
	ctx := context.Background()

	// Call Coordinator service to get booking history
	result, err := c.codClient.ViewHistory(ctx, &cpb.View{
		Id:     p.Id,
		Page:   p.Page,
		Status: "false",
	})
	if err != nil {
		return nil, err
	}

	// Map Coordinator response to UserHistories
	var history []*pb.UserHistory
	for _, hstry := range result.History {
		userHistory := &pb.UserHistory{
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
		}
		history = append(history, userHistory)
	}

	// Create and return UserHistories response
	return &pb.UserHistories{
		History: history,
	}, nil
}

// ViewBookingSvc handles the user request to view a specific booking.
func (c *UserSVC) ViewBookingSvc(p *pb.UserView) (*pb.UserHistory, error) {
	// Create a new context
	ctx := context.Background()

	// Call Coordinator service to get booking details
	result, err := c.codClient.ViewBooking(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return nil, err
	}

	// Map Coordinator response to UserHistory
	var travellers []*pb.UserTravellerDetails
	for _, traveller := range result.Travellers {
		travellerDetails := &pb.UserTravellerDetails{
			Name:   traveller.Name,
			Age:    traveller.Age,
			Gender: traveller.Gender,
		}
		travellers = append(travellers, travellerDetails)
	}

	// Create and return UserHistory response
	return &pb.UserHistory{
		Id:              result.Id,
		PaymentMode:     result.PaymentMode,
		BookingStatus:   result.BookingStatus,
		CancelledStatus: result.CancelledStatus,
		TotalPrice:      int64(result.TotalPrice),
		UserId:          int64(result.UserId),
		BookingId:       result.BookDate,
		StartDate:       result.StartDate,
		Travellers:      travellers,
		PaidAmount:      result.PaidAmount,
	}, nil
}

// CancelBookingSvc handles the user request to cancel a booking.
func (c *UserSVC) CancelBookingSvc(p *pb.UserView) (*pb.UserResponce, error) {
	// Create a new context
	ctx := context.Background()

	// Find user by ID
	user, err := c.Repo.FindUserById(uint(p.UserId))
	if err != nil {
		return nil, errors.New("user not found with this ID")
	}

	// Call Coordinator service to cancel booking
	result, err := c.codClient.CancelBooking(ctx, &cpb.View{
		Id: p.Id,
	})
	if err != nil {
		return nil, err
	}

	// Update user's wallet
	user.Wallet += int(result.Amount)
	err = c.Repo.UpdateUser(user)
	if err != nil {
		return nil, errors.New("error while updating user wallet")
	}

	// Create and return UserResponce
	return &pb.UserResponce{
		Status:  result.Status,
		Message: result.Message,
	}, nil
}
