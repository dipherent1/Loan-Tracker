package domain

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// generate interface for the AuthRepo
type AuthRepo interface {
	Register(ctx context.Context, newUser *User) Respose
}

// generate interface for the AuthUsecase
type AuthUsecase interface {
	Register(c *gin.Context, newUser *User) Respose
}

type RefreshRepository interface {
	UpdateToken(ctx context.Context, refreshToken string, userid primitive.ObjectID) (error, int)
	DeleteToken(ctx context.Context, userid primitive.ObjectID) (error, int)
	FindToken(ctx context.Context, userid primitive.ObjectID) (string, error, int)
	StoreToken(ctx context.Context, userid primitive.ObjectID, refreshToken string) (error, int)
}
