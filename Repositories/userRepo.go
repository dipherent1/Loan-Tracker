package repositories

import (
	"context"
	custommongo "loaner/CustomMongo"
	domain "loaner/Domain"
	dtos "loaner/Dtos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// user repository struct
type UserRepo struct {
	verifiedCollections domain.Collection
}

// create new user repository
func NewUserRepository(database *mongo.Database) *UserRepo {
	return &UserRepo{
		verifiedCollections: custommongo.NewMongoCollection(database.Collection("verified")),
	}
}

// get user by id
func (u *UserRepo) GetUserById(ctx context.Context, id primitive.ObjectID) domain.Response {
	var user dtos.RegisterUserDto
	err := u.verifiedCollections.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return domain.Response{
			Status:  500,
			Message: "user not found",
		}
	}

	return domain.Response{
		Status:  200,
		Message: "User found",
		Data:    user,
	}
}
