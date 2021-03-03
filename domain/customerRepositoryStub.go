package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error)  {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Regis", "Kigali", "0000", "2000/01/01", "1"},
		{"1002", "Mzee", "Kigali", "0000", "2000/01/01", "2"},
	}
	return CustomerRepositoryStub{customers}
}
