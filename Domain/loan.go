package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Loan struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ApplicantID     primitive.ObjectID `bson:"applicant_id" json:"applicant_id"`
	Amount          float64            `bson:"amount" json:"amount" validate:"required"`
	LoanTerm        int                `bson:"loan_term" json:"loan_term" validate:"required"` // in months
	InterestRate    float64            `bson:"interest_rate" json:"interest_rate" validate:"required"`
	ApplicationDate time.Time          `bson:"application_date" json:"application_date"`
	Status          string             `bson:"status" json:"status"` // e.g., "pending", "approved", "rejected"
}
