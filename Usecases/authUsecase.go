package usecases

import (
	"context"
	domain "loaner/Domain"
	"time"

	"github.com/gin-gonic/gin"
)

// generate a new usecase that takes an authrepo and returns a new usecase
type AuthUsecase struct {
	authRepo       domain.AuthRepo
	contextTimeout time.Duration
}

// generate a new usecase that takes an authrepo and returns a new usecase
func NewAuthUsecase(authRepo domain.AuthRepo) *AuthUsecase {
	return &AuthUsecase{
		authRepo:       authRepo,
		contextTimeout: time.Second * 10,
	}
}

// register a new user
func (a *AuthUsecase) Register(c *gin.Context, newUser *domain.User) domain.Respose {
	// create a new context
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.authRepo.Register(ctx, newUser)

}
