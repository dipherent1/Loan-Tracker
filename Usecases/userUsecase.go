package usecases

import (
	"context"
	domain "loaner/Domain"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (a *UserUsecase) GetUserById(c *gin.Context, id primitive.ObjectID) domain.Response {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.authRepo.GetUserById(ctx, id)
}
