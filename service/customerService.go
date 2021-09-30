package service

import "github.com/JyotsnaGorle/banking/domain"

// primary port
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

// implementation of the interface CutomerService
type DefaultCustomerService struct {
	// has a dependency of the repository
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// helper func to instantiate the default customer service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}