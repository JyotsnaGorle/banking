package domain

import "github.com/JyotsnaGorle/banking/app/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `json:"full_name" db:"name"`
	City        string `json:"city" db:"city"`
	Zipcode     string `json:"zip_code" db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// secondary port - protocol
type CustomerRepository interface {
	// status == 1, status == 0, status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
