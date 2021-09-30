package domain

import "github.com/JyotsnaGorle/banking/app/errs"

type Customer struct {
	Id          string
	Name        string `json:"full_name" xml:"full_name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zip_code"`
	DateOfBirth string
	Status      string
}

// secondary port - protocol
type CustomerRepository interface {
	// status == 1, status == 0, status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
