package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	Customers := []Customer{
		{"1001", "Arjun Teli", "Bari-Sadri", "312403", "11-06-1999", "1"},
		{"1002", "Karan Teli", "Bari-Sadri", "312403", "11-06-1999", "1"},
	}

	return CustomerRepositoryStub{Customers}
}
