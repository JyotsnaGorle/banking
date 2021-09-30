package domain

// implementation of the interface CutomerInterface
// adapter of the secondary port
type CustomerRepositoryStub struct {
	customers []Customer
}

// duck typing
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "asdlkhj", City: "Neew delin", Zipcode: "11211", DateOfBirth: "2011-90-09", Status: "1"},
		{Id: "1002", Name: "rob", City: "Neew delin", Zipcode: "11211", DateOfBirth: "2011-90-09", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
