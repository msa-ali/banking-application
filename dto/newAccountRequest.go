package dto

import (
	"strings"

	"github.com/Altamashattari/banking-application/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit atleast 5000.00")
	}
	accountType := strings.ToLower(r.AccountType)
	if accountType != "saving" && accountType != "checking" {
		return errs.NewValidationError("Account type should be checking or saving ")
	}
	return nil
}
