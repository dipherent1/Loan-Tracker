package usecases

import (
	"context"
	domain "loaner/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
)

type UserUsecase struct {
	authRepo       domain.UserRepo
	contextTimeout time.Duration
}

func NewUserUsecase(authRepo domain.UserRepo) *UserUsecase {
	return &UserUsecase{
		authRepo:       authRepo,
		contextTimeout: time.Second * 10,
	}
}

func (a *UserUsecase) GetUserById(c *gin.Context, id primitive.ObjectID) domain.Respose {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.authRepo.GetUserById(ctx, id)
}
