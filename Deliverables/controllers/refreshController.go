package controllers

import (
	"fmt"
	domain "loaner/Domain"
	jwtservice "loaner/Infrastructure/jwtService"

	"github.com/gin-gonic/gin"
)

// RefreshController struct
type RefreshController struct {
	RefreshUseCase domain.RefreshUseCase
}

// NewRefreshController function
func NewRefreshController(usecase domain.RefreshUseCase) *RefreshController {
	return &RefreshController{
		RefreshUseCase: usecase,
	}
}

// Refresh function
func (r *RefreshController) Refresh(c *gin.Context) {
	accessClaims, err := Getclaim(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	// get the refresh token
	refreshToken, err, statuscode := r.RefreshUseCase.FindToken(c, accessClaims.ID)

	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}

	// check if the refresh token is valid
	if refreshToken == "" {
		c.JSON(401, gin.H{"error": "refresh token not found"})
		return
	}

	// verify the refresh token
	err = jwtservice.VerifyRefreshToken(refreshToken, accessClaims.ID)

	fmt.Println("refresh token", refreshToken)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	user := domain.User{
		ID:   accessClaims.ID,
	}

	newAccessToken, err := jwtservice.CreateAccessToken(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// newRefreshToken, err := jwtservice.CreateRefreshToken(user)

	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// store the refresh token
	// err, statuscode = r.RefreshUseCase.StoreToken(c,  accessClaims.ID, newRefreshToken)

	// if err != nil {
	// 	c.JSON(statuscode, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(200, gin.H{"access_token": newAccessToken, "refresh_token": refreshToken})

}
