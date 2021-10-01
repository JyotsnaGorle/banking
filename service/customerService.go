package service

import (
	"github.com/JyotsnaGorle/banking/domain"
	"github.com/JyotsnaGorle/banking/dto"
	"github.com/JyotsnaGorle/banking/errs"
)

// primary port
type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// implementation of the interface CutomerService
type DefaultCustomerService struct {
	// has a dependency of the repository
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	allCustomers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var allResponses = make([]dto.CustomerResponse, 0)

	for _, eachCustomer := range allCustomers {
		allResponses = append(allResponses, eachCustomer.ToDto())
	}

	return allResponses, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

// helper func to instantiate the default customer service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
