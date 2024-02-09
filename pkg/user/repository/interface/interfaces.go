package interfaces

import DOM "github.com/Shakezidin/pkg/dom"

type UserRepoInter interface {
	SignupRepo(user *DOM.User) error
	FindUserByEmail(email string) (*DOM.User, error)
	FindUserByPhone(number int) (*DOM.User, error)
	CreateUser(user *DOM.User) error
	UpdatePassword(id uint, newpassword string) error
	FindUserById(id uint) (*DOM.User, error)
	UpdateUser(user *DOM.User) error
	UserCount() int64
	FindUsers(offset, limit int) ([]*DOM.User, error)
}
