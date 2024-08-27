package domain

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// generate interface for the AuthRepo
type AuthRepo interface {
	Register(ctx context.Context, newUser *User) Respose
	Login(ctx context.Context, user User) Respose
	Activate(ctx context.Context, token string) Respose

}

// generate interface for the AuthUsecase
type AuthUsecase interface {
	Register(c *gin.Context, newUser *User) Respose
	Login(c *gin.Context, user User) Respose
	Activate (c *gin.Context, token string) Respose
}

type RefreshRepository interface {
	UpdateToken(ctx context.Context, refreshToken string, userid primitive.ObjectID) (error, int)
	DeleteToken(ctx context.Context, userid primitive.ObjectID) (error, int)
	FindToken(ctx context.Context, userid primitive.ObjectID) (string, error, int)
	StoreToken(ctx context.Context, userid primitive.ObjectID, refreshToken string) (error, int)
}

type RefreshUseCase interface {
	// UpdateToken(c *gin.Context, refreshToken string, userid primitive.ObjectID) (error, int)
	DeleteToken(c *gin.Context, userid primitive.ObjectID) (error, int)
	FindToken(c *gin.Context, userid primitive.ObjectID) (string, error, int)
	StoreToken(c *gin.Context, userid primitive.ObjectID, refreshToken string) (error, int)
}