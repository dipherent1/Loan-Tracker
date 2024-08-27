package controllers

import (
	domain "loaner/Domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (a *UserController) Profile(c *gin.Context) {
	// get claims
	claim, err := Getclaim(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// get user by id
	response := a.userUsecase.GetUserById(c, claim.ID)
	
	c.JSON(response.Status, gin.H{
		"message": response.Message,
		"data":    response.Data,
	})
}
