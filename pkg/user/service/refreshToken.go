package service

import (
	"errors"
	"log"
	"strconv"

	"github.com/Shakezidin/config"
	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/Shakezidin/utils"
)

func (c *UserSVC) RefreshTokenSVC(p *pb.TokenData) (*pb.UserLoginResponce, error) {
	id, _ := strconv.Atoi(p.Id)
	user, err := c.Repo.FindUserById(uint(id))
	if err != nil {
		return nil, errors.New("User not found")
	}
	token, refreshToken, err := utils.GenerateToken(user.Email, p.Role, p.Id, config.LoadConfig().SECRETKEY)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", p.Email, err.Error())
		return nil, err
	}

	return &pb.UserLoginResponce{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
