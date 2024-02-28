package service

import (
	"errors"
	"strconv"

	"github.com/Shakezidin/config"
	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/Shakezidin/utils"
)

func (c *UserSVC) RefreshTokenSVC(p *pb.TokenData) (*pb.UserLoginResponse, error) {
	// Convert string id to integer
	id, err := strconv.Atoi(p.ID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	// Find user by id
	user, err := c.Repo.FindUserById(uint(id))
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Generate new tokens
	token, refreshToken, err := utils.GenerateToken(user.Email, p.Role, p.ID, config.LoadConfig().SECRETKEY)
	if err != nil {
		return nil, errors.New("error generating tokens")
	}

	// Return response with new tokens
	return &pb.UserLoginResponse{
		Token:         token,
		Refresh_Token: refreshToken,
	}, nil
}
