package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Id    string
	Email string
	Role  string
	jwt.StandardClaims
}

func GenerateToken(email, role, id string, cfg string) (string, error) {
	expireTime := time.Now().Add(time.Minute * 20).Unix()
	claims := &Claims{
		Id:    id,
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Subject:   email,
			IssuedAt:  time.Now().Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(cfg))
	if err != nil {
		log.Printf("unable to generate jwt token for user %v, err: %v", email, err.Error())
		return "", err
	}

	return signedToken, nil
}
