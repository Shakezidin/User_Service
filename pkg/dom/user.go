package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null,unique"`
	Phone    int    `gorm:"not null,unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
	Isblock  bool   `gorm:"default:false"`
}

type EmailMessage struct {
	Email   string
	Subject string
	Content string
}

type OtpData struct {
	Otp        int       `json:"otp"`
	Email      string    `json:"email"`
	ExpireTime time.Time `json:"time"`
}
