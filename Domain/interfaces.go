package domain

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// generate interface for the AuthRepo
type AuthRepo interface {
	Register(ctx context.Context, newUser *User) Response
	Login(ctx context.Context, user User) Response
	Activate(ctx context.Context, token string) Response
}

// generate interface for the AuthUsecase
type AuthUsecase interface {
	Register(c *gin.Context, newUser *User) Response
	Login(c *gin.Context, user User) Response
	Activate(c *gin.Context, token string) Response
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

type UserRepo interface {
	GetUserById(ctx context.Context, id primitive.ObjectID) Response
}

type UserUsecase interface {
	GetUserById(c *gin.Context, id primitive.ObjectID) Response
}

type LoanRepo interface {
	Apply(ctx context.Context, loan *Loan) Response
	GetLoanById(ctx context.Context, loanID primitive.ObjectID, userID primitive.ObjectID) Response
}

type LoanUsecase interface {
	Apply(c *gin.Context, loan *Loan) Response
	GetLoanById(c *gin.Context, loanID primitive.ObjectID, userID primitive.ObjectID) Response
}
