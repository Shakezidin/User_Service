package interfaces

import DOM "github.com/Shakezidin/pkg/dom"

type UserRepoInter interface {
	SignupRepo(user *DOM.User) error
	FindUserByEmail(email string) (*DOM.User, error)
	FindUserByPhone(number int) (*DOM.User, error)
	CreateUser(user *DOM.User) error
	UpdatePassword(id uint, newpassword string) error 
}
