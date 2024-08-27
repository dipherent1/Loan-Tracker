package controllers

import (
	domain "loaner/Domain"
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
	// extract the data from the response, it should be a dtos.UserRegistrstion type
	// then return the data to the user

	if response.Status != http.StatusOK {
		c.IndentedJSON(response.Status,
			gin.H{"message": response.Message})
		return
	}
}

// Login is a function that logs in a user
func (a *AuthController) Login(c *gin.Context) {
	// code to login a user
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data", "error": err.Error()})
		return
	}

	if err := a.validator.Struct(user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": err.Error()})
		return
	}

	response := a.authUsecase.Login(c, user)
	// extract the data from the response, it should be a domain.token type

	if response.Status != http.StatusOK {
		c.IndentedJSON(response.Status,
			gin.H{"message": response.Message})
		return
	}
	c.IndentedJSON(response.Status, gin.H{"access_token": response.AccessToken})
}

// activate is a function that activates a user
func (a *AuthController) Activate(c *gin.Context) {
	// code to activate a user
	token := c.Param("token")
	// if activation token doesnt exist return an error
	if token == "" {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"message": "invalid token"})
		return
	}
	response := a.authUsecase.Activate(c, token)

	c.IndentedJSON(response.Status,
		gin.H{"message": response.Message})
}
