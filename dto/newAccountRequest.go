package dto

import (
	"strings"

	"github.com/JyotsnaGorle/banking/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit more than 5000.")
	}

	if strings.ToLower(r.AccountType) != "saving" && r.AccountType != "checking" {
		return errs.NewValidationError("Account type should be checking or savings")
	}
	return nil
}
