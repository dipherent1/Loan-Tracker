package controllers

import (
	domain "loaner/Domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// LoanController struct
type LoanController struct {
	loanUsecase domain.LoanUsecase
	validator  *validator.Validate
}

// NewLoanController creates a new loan controller
func NewLoanController(loanUsecase domain.LoanUsecase) *LoanController {
	return &LoanController{
		loanUsecase: loanUsecase,
		validator:  validator.New(),
	}
}

// Apply for a loan
func (l *LoanController) Apply(c *gin.Context) {
	var loan domain.Loan
	err := c.ShouldBindJSON(&loan)
	if err != nil {
		c.JSON(400, domain.Respose{
			Status:  400,
			Message: "Invalid request",
		})
		return
	}

	err = l.validator.Struct(&loan)
	if err != nil {
		c.JSON(400, domain.Respose{
			Status:  400,
			Message: "Invalid request missing fields",
		})
		return
	}

	response := l.loanUsecase.Apply(c, &loan)
	c.JSON(response.Status, gin.H{
		"message": response.Message,
	})
}
