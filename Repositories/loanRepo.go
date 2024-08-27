package repositories

import (
	"context"
	"fmt"
	custommongo "loaner/CustomMongo"
	domain "loaner/Domain"
	dtos "loaner/Dtos"
	utils "loaner/Utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// loan repository struct
type LoanRepo struct {
	loanCollections domain.Collection
	UserRepo        domain.UserRepo
}

// create new loan repository
func NewLoanRepository(database *mongo.Database) *LoanRepo {
	return &LoanRepo{
		loanCollections: custommongo.NewMongoCollection(database.Collection("loans")),
		UserRepo:        NewUserRepository(database),
	}
}

// apply
func (l *LoanRepo) Apply(ctx context.Context, loan *domain.Loan) domain.Response {
	_, err := l.loanCollections.InsertOne(ctx, loan)
	if err != nil {
		return domain.Response{
			Status:  500,
			Message: "Failed to apply for loan",
		}
	}
	return domain.Response{
		Status:  200,
		Message: "Loan application successful",
	}
}

// get loan by id
func (l *LoanRepo) GetLoanById(ctx context.Context, loanID primitive.ObjectID, userID primitive.ObjectID) domain.Response {
	var loan domain.Loan
	// check if user is the owner of the loan or admin
	response := l.UserRepo.GetUserById(ctx, userID)
	if response.Status != 200 {
		return response
	}
	// check if user is the owner of the loan or admin
	user := response.Data.(dtos.RegisterUserDto)
	fmt.Println(user)

	err := l.loanCollections.FindOne(ctx, primitive.M{"_id": loanID}).Decode(&loan)
	if err != nil {
		return domain.Response{
			Status:  404,
			Message: "Loan not found",
		}
	}

	autherized := utils.IsAuthorOrAdmin(userID, loan.ApplicantID, user.Role)
	if !autherized {
		return domain.Response{
			Status:  403,
			Message: "You are not authorized to view this loan",
		}
	}

	return domain.Response{
		Status:  200,
		Message: "Loan found",
		Data:    loan,
	}
}

// get all loans
func (l *LoanRepo) GetAllLoans(ctx context.Context, filter domain.Filter) domain.Response {
	var loans []domain.Loan

	pipeline := mongo.Pipeline{
		bson.D{{"$set", bson.D{
			{"statusOrder", bson.D{
				{"$indexOfArray", bson.A{filter.StatusOrder, "$status"}},
			}},
		}}},
		bson.D{{"$sort", bson.D{{"statusOrder", 1}}}},
		bson.D{{"$unset", "statusOrder"}},
	}

	cur, err := l.loanCollections.Aggregate(ctx, pipeline)
	if err != nil {
		return domain.Response{
			Status:  500,
			Message: "Failed to get loans",
		}
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var loan domain.Loan
		err := cur.Decode(&loan)
		if err != nil {
			return domain.Response{
				Status:  500,
				Message: "Failed to get loans",
			}
		}
		loans = append(loans, loan)
	}
	

	return domain.Response{
		Status:  200,
		Message: "Loans found",
		Data:    loans,
	}
}
