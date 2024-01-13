package service

import (
	"github.com/Shakezidin/config"
	inter "github.com/Shakezidin/pkg/user/repository/interface"
	SVCinter "github.com/Shakezidin/pkg/user/service/interface"
	"github.com/go-redis/redis/v8"
)

type UserSVC struct {
	Repo   inter.UserRepoInter
	twilio *config.TwilioVerify
	redis  *redis.Client
	cfg    *config.Config
}

func NewUserSVC(repo inter.UserRepoInter, twilio *config.TwilioVerify, redis *redis.Client, cfg *config.Config) SVCinter.UserSVCInter {
	return &UserSVC{
		Repo:   repo,
		twilio: twilio,
		redis:  redis,
		cfg:    cfg,
	}
}
