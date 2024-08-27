package controllers

import (
	domain "loaner/Domain"
	utils "loaner/Utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// LoanController struct
type LoanController struct {
	loanUsecase domain.LoanUsecase
	validator   *validator.Validate
}

// NewLoanController creates a new loan controller
func NewLoanController(loanUsecase domain.LoanUsecase) *LoanController {
	return &LoanController{
		loanUsecase: loanUsecase,
		validator:   validator.New(),
	}
}

// Apply for a loan
func (l *LoanController) Apply(c *gin.Context) {
	// get claim
	claim, err := Getclaim(c)
	if err != nil {
		c.JSON(400, domain.Response{
			Status:  400,
			Message: "Invalid request",
		})
		return
	}

	var loan *domain.Loan
	err = c.ShouldBindJSON(&loan)
	if err != nil {
		c.JSON(400, domain.Response{
			Status:  400,
			Message: "Invalid request",
		})
		return
	}

	err = l.validator.Struct(loan)
	if err != nil {
		c.JSON(400, domain.Response{
			Status:  400,
			Message: "Invalid request missing fields",
		})
		return
	}

	loan.ApplicantID = claim.ID
	loan.Status = "pending"
	loan.ApplicationDate = time.Now()

	response := l.loanUsecase.Apply(c, loan)
	c.JSON(response.Status, gin.H{
		"message": response.Message,
	})
}

// GetLoanById get loan by id
func (l *LoanController) GetLoanById(c *gin.Context) {
	// get claim
	claim, err := Getclaim(c)
	if err != nil {
		c.JSON(400, domain.Response{
			Status:  400,
			Message: "error getting claim",
		})
		return
	}

	loanid := c.Param("loanID")
	if loanid == "" {
		c.JSON(400, domain.Response{
			Status:  400,
			Message: "error getting loan id",
		})
		return
	}

	loanID, err := utils.StringToObjectId(loanid)
	if err != nil {
		c.JSON(400, domain.Response{
			Status:  400,
			Message: "error converting loan id",
		})
		return
	}

	response := l.loanUsecase.GetLoanById(c, loanID, claim.ID)
	
	c.JSON(response.Status, gin.H{
		"message": response.Message,
		"data":    response.Data,
	})
}
