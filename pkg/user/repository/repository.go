package repository

import (
	DOM "github.com/Shakezidin/pkg/dom"
)

func (c *UserRepo) SignupRepo(user *DOM.User) error {
	if err := c.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (c *UserRepo) FindUserByEmail(email string) (*DOM.User, error) {
	var user DOM.User
	if err := c.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *UserRepo) FindUserByPhone(number int) (*DOM.User, error) {
	var user DOM.User
	if err := c.db.Where("phone = ?", number).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *UserRepo) CreateUser(user *DOM.User) error {
	if err := c.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (c *UserRepo) UpdatePassword(id uint, newpassword string) error {
	user := DOM.User{}
	user.ID = id

	result := c.db.Model(&user).Update("password", newpassword)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
