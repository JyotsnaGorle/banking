package domain

type Customer struct {
	Id          string
	Name        string `json:"full_name" xml:"full_name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zip_code"`
	DateOfBirth string
	Status      string
}

// secondary port
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
