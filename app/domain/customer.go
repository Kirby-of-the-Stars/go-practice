package domain

import "github.com/TrueRou/practice/app/dto"

type Customer struct {
	Id      int64  `db:"id"`
	Name    string `db:"name"`
	City    string `db:"city"`
	Zipcode string `db:"zipcode"`
}

func (customer *Customer) ToDto() *dto.CustomerResponse {
	return &dto.CustomerResponse{
		Name:    customer.Name,
		City:    customer.City,
		Zipcode: customer.Zipcode,
	}
}

type CustomerRepository interface {
	FindAll() (*[]Customer, error)
	FindByName(name string) (*Customer, error)
}
