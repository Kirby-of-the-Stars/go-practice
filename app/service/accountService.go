package service

import (
	"github.com/TrueRou/practice/app/domain"
	"github.com/TrueRou/practice/app/dto"
	"github.com/TrueRou/practice/app/errs"
)

type AccountService interface {
	GetAllAccounts() (*[]dto.AccountResponse, error)
	GetAccountsByCustomerName(name string) (*[]dto.AccountResponse, error)
	CreateAccount(request dto.NewAccountRequest) (int64, error)
}

type DefaultAccountService struct {
	repo         domain.AccountRepository
	customerRepo domain.CustomerRepository
}

func (service DefaultAccountService) GetAllAccounts() (*[]dto.AccountResponse, error) {
	accounts, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var accountResponses []dto.AccountResponse
	for _, account := range *accounts {
		accountResponses = append(accountResponses, *account.ToDto())
	}
	return &accountResponses, nil
}

func (service DefaultAccountService) GetAccountsByCustomerName(name string) (*[]dto.AccountResponse, error) {
	customer, err := service.customerRepo.FindByName(name)
	if err != nil {
		return nil, err
	}
	accounts, err := service.repo.FindByCustomerId(customer.Id)
	if err != nil {
		return nil, err
	}
	if len(*accounts) == 0 {
		return nil, errs.NewStatusNotFoundError("No accounts found for customer " + name)
	}
	var accountResponses []dto.AccountResponse
	for _, account := range *accounts {
		accountResponses = append(accountResponses, *account.ToDto())
	}
	return &accountResponses, nil
}

func (service DefaultAccountService) CreateAccount(request dto.NewAccountRequest) (int64, error) {
	var account = (&domain.Account{}).FromDto(request)
	return service.repo.Create(account)
}

func NewAccountService(repo domain.AccountRepository, customerRepo domain.CustomerRepository) AccountService {
	return DefaultAccountService{repo, customerRepo}
}
