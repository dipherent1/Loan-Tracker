package controllers

import (
	"errors"
	domain "loaner/Domain"

	"github.com/gin-gonic/gin"
)

func Getclaim(c *gin.Context) (*domain.AccessClaims, error) {
	claim, exists := c.Get("claim")
	if !exists {
		return nil, errors.New("claim not set")
	}

	userClaims, ok := claim.(*domain.AccessClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return userClaims, nil
}
