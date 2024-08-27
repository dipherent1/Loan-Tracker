package repositories

import (
	"context"
	"fmt"
	custommongo "loaner/CustomMongo"
	domain "loaner/Domain"
	dtos "loaner/Dtos"
	passwordservice "loaner/Infrastructure/passwordService"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// generate a new repository that takes a registered, unregistered, and refreshtokn collection
// and returns a new repository
type AuthRepo struct {
	verified     custommongo.Collection
	unverified   custommongo.Collection
	refreshtoken custommongo.Collection
}

// generate a new repository that takes a registered, unregistered, and refreshtokn collection
// and returns a new repository
func NewAuthRepo(database custommongo.Database) *AuthRepo {

	return &AuthRepo{
		verified:     database.Collection("verified"),
		unverified:   database.Collection("unverified"),
		refreshtoken: database.Collection("refreshtoken"),
	}
}

// register a new user
func (a *AuthRepo) Register(ctx context.Context, newUser *domain.User) domain.Respose {
	err := passwordservice.CheckPasswordStrength(newUser.Password)
	if err != nil {
		return domain.Respose{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	hashedPassword, err := passwordservice.GenerateFromPasswordCustom(newUser.Password)
	if err != nil {
		return domain.Respose{
			Status:  http.StatusInternalServerError,
			Message: "Error hashing password",
		}
	}

	newUser.Password = hashedPassword
	if newUser.UserName != "" {
		newUser.UserName = newUser.Email + "_user"
	}

	newUser.ID = primitive.NewObjectID()
	newUser.Role = "user"
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	InsertedID, err := a.unverified.InsertOne(ctx, newUser)
	if err != nil {
		return domain.Respose{
			Status:  http.StatusInternalServerError,
			Message: "Error inserting user",
		}
	}

	fmt.Println("InsertedID")
	fmt.Println(InsertedID)
	fmt.Println("InsertedID")

	var userDto dtos.RegisterUserDto
	// get user from database
	insertedID := InsertedID.(primitive.ObjectID)
	err = a.unverified.FindOne(ctx, bson.D{{"_id", insertedID}}).Decode(&userDto)
	if err != nil {
		return domain.Respose{
			Status:  http.StatusInternalServerError,
			Message: "Error getting user",
		}
	}
	return domain.Respose{
		Status:  http.StatusOK,
		Message: "User registered successfully",
		Data:    userDto,
	}
}
