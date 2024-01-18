package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// The CustomerRepository interface defines a method to retrieve all customers.
// @property FindAll - A method that returns a slice of Customer objects and an error. It is used to
// retrieve all customers from a data source.
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

// The CustomerRepositoryStub type is a stub implementation of a customer repository.
// @property {[]Customer} customers - A slice of Customer objects.
type CustomerRepositoryStub struct {
	customers []Customer
}

// The `FindAll()` function is a method of the `CustomerRepositoryStub` struct. It returns a slice of
// `Customer` objects and an error.
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// helper function responsible for creating new customerRepoStub
// The function newCustomerRepositoryStub returns a new instance of CustomerRepositoryStub with a
// predefined list of customers.
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Arjun Teli", "Bari-Sadri", "312403", "11-06-1999", "1"},
		{"1002", "Karan Teli", "Bari-Sadri", "312403", "11-06-1999", "1"},
	}

	return CustomerRepositoryStub{customers}
}
