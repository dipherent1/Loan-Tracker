package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

// generate interface for the AuthRepo
type AuthRepo interface {
	Register(ctx context.Context, newUser *User) Respose
}

// generate interface for the AuthUsecase
type AuthUsecase interface {
	Register(c *gin.Context, newUser *User) Respose
}
