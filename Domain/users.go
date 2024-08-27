package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	Email      string    `json:"email" validate:"required" bson:"email"`
	Password   string    `json:"password" validate:"required" bson:"password"`
	Name       string    `json:"name" bson:"name"`
	UserName   string    `json:"username" bson:"username"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
	Role       string    `json:"role" bson:"role"`
	IsVerified bool      `json:"is_verified" bson:"is_verified"`
}
