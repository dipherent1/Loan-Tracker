package authmiddleware

import (
	"context"
	"fmt"
	"loaner/Config"
	domain "loaner/Domain"
	dtos "loaner/Dtos"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT validation logic

		// JWT validation logic
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return Config.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid JWT user must login"})
			c.Abort()
			return
		}

		// Extract the user data from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(500, gin.H{"error": "Failed to parse token claims"})
			c.Abort()
			return
		}

		userID, err := primitive.ObjectIDFromHex(claims["id"].(string))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort() // Stop further handlers from being executed
			return
		}

		if !ok {
			c.JSON(500, gin.H{"error": "Failed to parse user ID"})
			c.Abort()
			return

		}

		//set claims to the context
		c.Set("claim", &domain.AccessClaims{
			ID: userID,
		})

		c.Next()
	}
}

func IsAdminMiddleware(userRepo domain.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {

		claim := c.MustGet("claim").(domain.AccessClaims)

		//get contec with timeout
		ctx, cancel := context.WithTimeout(c, time.Second*10)
		defer cancel()
		response := userRepo.GetUserById(ctx, claim.ID)
		
		user:= response.Data.(dtos.RegisterUserDto)
		if user.Role == "admin" {
			c.Next()
		} else {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

	}
}
