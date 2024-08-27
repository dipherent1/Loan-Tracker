package repositories

import (
	"context"
	custommongo "loaner/CustomMongo"
	domain "loaner/Domain"

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
func (l *LoanRepo) Apply(ctx context.Context, loan *domain.Loan) domain.Respose {
	_, err := l.loanCollections.InsertOne(ctx, loan)
	if err != nil {
		return domain.Respose{
			Status:  500,
			Message: "Failed to apply for loan",
		}
	}
	return domain.Respose{
		Status:  200,
		Message: "Loan application successful",
	}
}
