package controllers

import (
	domain "loaner/Domain"
	dtos "loaner/Dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// AuthController is a struct that contains the AuthUsecase
type AuthController struct {
	authUsecase domain.AuthUsecase
	validator   *validator.Validate
}

// NewAuthController is a function that returns a new AuthController
func NewAuthController(authUsecase domain.AuthUsecase) *AuthController {
	return &AuthController{
		authUsecase: authUsecase,
		validator:   validator.New(),
	}
}

// Register is a function that registers a new user
func (a *AuthController) Register(c *gin.Context) {
	// code to register a new user
	var newUser *domain.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data", "error": err.Error()})
		return
	}

	if err := a.validator.Struct(newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": err.Error()})
		return
	}

	response := a.authUsecase.Register(c, newUser)
	data := response.Data.(dtos.RegisterUserDto)

	if response.Status != http.StatusOK {
		c.IndentedJSON(response.Status,
			gin.H{"message": response.Message,
				"data": data})
		return
	}
}
