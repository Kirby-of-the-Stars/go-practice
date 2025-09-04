package repository

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/TrueRou/practice/app/domain"
	"github.com/TrueRou/practice/app/errs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (repo CustomerRepositoryDb) FindAll() (*[]domain.Customer, error) {
	var customers []domain.Customer
	queryStr := "SELECT id, name, city, zipcode FROM customers"
	err := repo.client.Select(&customers, queryStr)
	if err != nil {
		log.Print("Failed to scan the record.")
		return nil, errs.NewStatusInternalServerError("Unexpected database error.")
	}

	return &customers, nil
}

func (repo CustomerRepositoryDb) FindByName(name string) (*domain.Customer, error) {
	var customer domain.Customer
	queryStr := "SELECT id, name, city, zipcode FROM customers WHERE name = ?"
	err := repo.client.Get(&customer, queryStr, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewStatusNotFoundError(err.Error())
		}
		log.Println("Failed to scan the record.")
		return nil, errs.NewStatusInternalServerError("Unexpected database error.")
	}
	return &customer, nil
}

func NewCustomerRepositoryDb() domain.CustomerRepository {
	db, err := sqlx.Open("mysql", "root:nekomeow@/practice")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDb{db}
}
