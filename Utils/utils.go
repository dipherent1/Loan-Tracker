package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectIdToString(objID primitive.ObjectID) string {
	return primitive.ObjectID.Hex(objID)
}

func StringToObjectId(str string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return objID, nil
}

// is user author of post or admin
func IsAuthorOrAdmin(userID primitive.ObjectID, ApplicantID primitive.ObjectID, role string) bool {
	if userID == ApplicantID || role == "admin" {
		return true
	}
	return false
}
