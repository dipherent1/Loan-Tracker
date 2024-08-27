package repositories

import (
	"context"
	"fmt"
	"loaner/Config"
	custommongo "loaner/CustomMongo"
	domain "loaner/Domain"
	dtos "loaner/Dtos"
	emailservice "loaner/Infrastructure/emailService"
	jwtservice "loaner/Infrastructure/jwtService"
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
	emailservice emailservice.MailTrapService
}

// generate a new repository that takes a registered, unregistered, and refreshtokn collection
// and returns a new repository
func NewAuthRepo(database custommongo.Database) *AuthRepo {

	return &AuthRepo{
		verified:     database.Collection("verified"),
		unverified:   database.Collection("unverified"),
		refreshtoken: database.Collection("refreshtoken"),
		emailservice: emailservice.NewMailTrapService(),
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

	fmt.Println(userDto)

	err, status := a.SendActivationEmail(newUser.Email)
	fmt.Println(err)
	if err != nil {
		return domain.Respose{
			Status:  status,
			Message: "Error sending activation email",
		}
	}

	return domain.Respose{
		Status:  http.StatusOK,
		Message: "User registered successfully",
		Data:    userDto,
	}
}

// login a user
func (a *AuthRepo) Login(ctx context.Context, user domain.User) domain.Respose {
	filter := bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "username", Value: user.UserName}},
			bson.D{{Key: "email", Value: user.Email}},
		}},
	}

	var existingUser domain.User
	err := a.verified.FindOne(ctx, filter).Decode(&existingUser)
	if err != nil {
		return domain.Respose{
			Status:  http.StatusNotFound,
			Message: "User not found user may not be activated",
		}
	}

	if !passwordservice.CompareHashAndPasswordCustom(existingUser.Password, user.Password) {
		return domain.Respose{
			Status:  http.StatusUnauthorized,
			Message: "Invalid password",
		}
	}

	return domain.Respose{}
	
}


// func (a *AuthRepo) GenerateTokenFromUser(ctx context.Context, existingUser domain.User) (domain.Tokens, error, int) {

// 	// filter := bson.D{{Key: "email", Value: existingUser.Email}}
// 	// Generate JWT access
// 	jwtAccessToken, err := jwtservice.CreateAccessToken(existingUser)
// 	if err != nil {
// 		return domain.Tokens{}, err, 500
// 	}
// 	refreshToken, err := jwtservice.CreateRefreshToken(existingUser)
// 	if err != nil {
// 		return domain.Tokens{}, err, 500
// 	}

// 	// filter := primitive.D{{"_id", existingUser.ID}}
// 	existingToken, err, statusCode := a.TokenRepository.FindToken(ctx, existingUser.ID)
// 	if err != nil && err.Error() != "mongo: no documents in result" {
// 		fmt.Println("error at count", err)
// 		return domain.Tokens{}, err, statusCode
// 	}

// 	if existingToken != "" {
// 		// update the refresh token
// 		err, statusCode := a.TokenRepository.UpdateToken(ctx, refreshToken, existingUser.ID)
// 		if err != nil {
// 			return domain.Tokens{}, err, statusCode
// 		}

// 	} else {
// 		err, statusCode := a.TokenRepository.StoreToken(ctx, existingUser.ID, refreshToken)
// 		if err != nil {
// 			return domain.Tokens{}, err, statusCode
// 		}
// 	}

// 	return domain.Tokens{
// 		AccessToken:  jwtAccessToken,
// 		RefreshToken: refreshToken,
// 	}, nil, 200
// }

func (a *AuthRepo) ActivateAccount(ctx context.Context, token string) domain.Respose {
	email, err := jwtservice.VerifyToken(token)
	if err != nil {
		return domain.Respose{
			Status:  http.StatusBadRequest,
			Message: "Invalid token",
		}
	}
	fmt.Println("email:", email, "token:", token)

	var user domain.User
	err = a.unverified.FindOne(ctx, bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return	domain.Respose{
			Status:  http.StatusInternalServerError,
			Message: "Error getting user",
		}
	}

	_, err = a.verified.InsertOne(ctx, user)
	if err != nil {
		return domain.Respose{
			Status:  http.StatusInternalServerError,
			Message: "Error inserting user",
		}
	}

	_, err = a.unverified.DeleteOne(ctx, bson.D{{"email", email}})
	if err != nil {
		return domain.Respose{
			Status:  http.StatusInternalServerError,
			Message: "Error deleting user",
		}
	}

	return domain.Respose{
		Status:  http.StatusOK,
		Message: "User activated successfully",
	}

}

func (a *AuthRepo) SendActivationEmail(email string) (error, int) {

	activationToken, err := jwtservice.GenerateToken(email)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = a.emailservice.SendEmail(email, "Verify Email", `Click "`+Config.BASE_URL+`/auth/activate/`+activationToken+`"here to verify email.
`, "reset")
	if err != nil {
		fmt.Println("in activation email 2")
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
