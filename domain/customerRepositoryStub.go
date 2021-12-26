package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"10001", "Altamash", "Bangalore", "560044", "3000-01-01", "1"},
		{"10002", "Ahmad", "Delhi", "5600244", "3000-01-01", "2"},
	}
	return CustomerRepositoryStub{customers}
}
