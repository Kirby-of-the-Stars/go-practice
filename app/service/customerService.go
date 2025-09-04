package service

import (
	"github.com/TrueRou/practice/app/domain"
	"github.com/TrueRou/practice/app/dto"
)

type CustomerService interface {
	GetAllCustomers() (*[]dto.CustomerResponse, error)
	GetCustomerByName(name string) (*dto.CustomerResponse, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (service DefaultCustomerService) GetAllCustomers() (*[]dto.CustomerResponse, error) {
	customers, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var customerResponses []dto.CustomerResponse
	for _, customer := range *customers {
		customerResponses = append(customerResponses, *customer.ToDto())
	}
	return &customerResponses, nil
}

func (service DefaultCustomerService) GetCustomerByName(name string) (*dto.CustomerResponse, error) {
	customer, err := service.repo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return customer.ToDto(), nil
}

func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repository}
}
