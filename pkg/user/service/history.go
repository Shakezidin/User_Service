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
		Id:     p.ID,
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
			ID:              hstry.Id,
			Payment_Mode:     hstry.PaymentMode,
			Booking_Status:   hstry.BookingStatus,
			Cancelled_Status: hstry.CancelledStatus,
			Total_Price:      int64(hstry.TotalPrice),
			User_ID:          int64(hstry.UserId),
			Booking_ID:       hstry.BookingId,
			Book_Date:        hstry.BookDate,
			Start_Date:       hstry.StartDate,
			Paid_Amount:      hstry.PaidAmount,
		}
		history = append(history, userHistory)
	}

	// Create and return UserHistories response
	return &pb.UserHistories{
		Histories: history,
	}, nil
}

// ViewBookingSvc handles the user request to view a specific booking.
func (c *UserSVC) ViewBookingSvc(p *pb.UserView) (*pb.UserHistory, error) {
	// Create a new context
	ctx := context.Background()

	// Call Coordinator service to get booking details
	result, err := c.codClient.ViewBooking(ctx, &cpb.View{
		Id: p.ID,
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
		ID:              result.Id,
		Payment_Mode:     result.PaymentMode,
		Booking_Status:   result.BookingStatus,
		Cancelled_Status: result.CancelledStatus,
		Total_Price:      int64(result.TotalPrice),
		User_ID:          int64(result.UserId),
		Booking_ID:      result.BookDate,
		Start_Date:      result.StartDate,
		Travellers:      travellers,
		Paid_Amount:     result.PaidAmount,
	}, nil
}

// CancelBookingSvc handles the user request to cancel a booking.
func (c *UserSVC) CancelBookingSvc(p *pb.UserView) (*pb.UserResponse, error) {
	// Create a new context
	ctx := context.Background()

	// Find user by ID
	user, err := c.Repo.FindUserById(uint(p.User_ID))
	if err != nil {
		return nil, errors.New("user not found with this ID")
	}

	// Call Coordinator service to cancel booking
	result, err := c.codClient.CancelBooking(ctx, &cpb.View{
		Id: p.ID,
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
	return &pb.UserResponse{
		Status:  result.Status,
		Message: result.Message,
	}, nil
}
