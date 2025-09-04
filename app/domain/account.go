package domain

import (
	"time"

	"github.com/TrueRou/practice/app/dto"
)

type Account struct {
	Id          string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate int64   `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      int     `db:"status"`
}

type AccountRepository interface {
	Create(account *Account) (int64, error)
	FindAll() (*[]Account, error)
	FindByCustomerId(id int64) (*[]Account, error)
}

func (account *Account) ToDto() *dto.AccountResponse {
	openingDateStr := time.Unix(account.OpeningDate, 0).Format("2006-01-02 15:04:05")
	return &dto.AccountResponse{
		Id:          account.Id,
		OpeningDate: openingDateStr,
		AccountType: account.AccountType,
		Amount:      account.Amount,
	}
}

func (account *Account) FromDto(request dto.NewAccountRequest) *Account {
	account.CustomerId = request.CustomerId
	account.AccountType = request.AccountType
	account.Amount = request.Amount
	account.OpeningDate = time.Now().Unix()
	account.Status = 1
	return account
}
