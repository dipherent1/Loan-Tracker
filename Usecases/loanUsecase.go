package usecases

import (
	"context"
	domain "loaner/Domain"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LoanUsecase struct

type LoanUsecase struct {
	loanRepo       domain.LoanRepo
	contextTimeout time.Duration
}

// NewLoanUsecase creates a new loan usecase
func NewLoanUsecase(loanRepo domain.LoanRepo) *LoanUsecase {
	return &LoanUsecase{
		loanRepo:       loanRepo,
		contextTimeout: time.Second * 10,
	}
}

// Apply for a loan
func (l *LoanUsecase) Apply(c *gin.Context, loan *domain.Loan) domain.Response {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()

	return l.loanRepo.Apply(ctx, loan)
}

// GetLoanById get loan by id
func (l *LoanUsecase) GetLoanById(c *gin.Context, loanID  primitive.ObjectID, userID primitive.ObjectID) domain.Response {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()

	return l.loanRepo.GetLoanById(ctx, loanID, userID)
}
