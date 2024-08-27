package jwtservice

import (
	"loaner/Config"
	domain "loaner/Domain"
	"time"

	"github.com/golang-jwt/jwt"
)

var CreateAccessToken = func(existingUser domain.User) (string, error) {
	userclaims := &domain.AccessClaims{
		ID:   existingUser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
		},
	}

	// Create a new JWT token with the user claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userclaims)

	// Ensure Config.JwtSecret is of type []byte
	jwtToken, err := token.SignedString([]byte(Config.JwtSecret))
	return jwtToken, err
}
