package service

import "github.com/arjunteli/banking/domain"

// The above code defines an interface for a customer service that provides a method to retrieve all
// customers.
// @property getAllCustomer - A method that returns a slice of domain.Customer objects and an error.
// This method is used to retrieve all customers from a customer service system.
type CustomerService interface {
	getAllCustomer() ([]domain.Customer, error)
}

// The `DefaultCustomerService` struct is defining a new type that represents a default implementation
// of the `CustomerService` interface. It has a single field `repo` of type
// `domain.CustomerRepository`. This field is used to interact with the data storage layer to retrieve
// customer data.
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// The `func (s DefaultCustomerService) getAllCustomer() ([]domain.Customer, error)` is a method of the
// `DefaultCustomerService` struct that implements the `CustomerService` interface.
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// The NewCustomerService function returns a new instance of the DefaultCustomerService struct with the
// given CustomerRepository.
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
