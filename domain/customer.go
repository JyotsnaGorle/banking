package domain

import (
	"github.com/JyotsnaGorle/banking/app/errs"
	"github.com/JyotsnaGorle/banking/dto"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `json:"full_name" db:"name"`
	City        string `json:"city" db:"city"`
	Zipcode     string `json:"zip_code" db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// method to the domain object
func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

// secondary port - protocol
type CustomerRepository interface {
	// status == 1, status == 0, status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
