package repository

import (
	inter "github.com/Shakezidin/pkg/user/repository/interface"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) inter.UserRepoInter {
	return &UserRepo{
		db: db,
	}
}
