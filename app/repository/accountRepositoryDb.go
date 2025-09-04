package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/TrueRou/practice/app/domain"
	"github.com/TrueRou/practice/app/errs"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (repo AccountRepositoryDb) Create(account *domain.Account) (int64, error) {
	queryStr := "INSERT INTO accounts (account_id, customer_id, opening_date, account_type, amount, status) VALUES (:account_id, :customer_id, :opening_date, :account_type, :amount, :status)"
	result, err := repo.client.NamedExec(queryStr, &account)
	if err != nil {
		fmt.Println(err)
		return -1, errs.NewStatusInternalServerError("Unexpected database error.")
	}
	return result.LastInsertId()
}

func (repo AccountRepositoryDb) FindAll() (*[]domain.Account, error) {
	queryStr := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts"
	var accounts []domain.Account
	err := repo.client.Select(&accounts, queryStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewStatusNotFoundError(err.Error())
		}
		log.Println("Unexpected database error.", err.Error())
		return nil, errs.NewStatusInternalServerError("Unexpected database error.")
	}
	return &accounts, nil
}

func (repo AccountRepositoryDb) FindByCustomerId(id int64) (*[]domain.Account, error) {
	queryStr := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts WHERE customer_id = ?"
	var accounts []domain.Account
	err := repo.client.Select(&accounts, queryStr, id)
	if err != nil {
		return nil, errs.NewStatusInternalServerError("Unexpected database error.")
	}
	return &accounts, nil
}

func NewAccountRepositoryDb() domain.AccountRepository {
	db, err := sqlx.Open("mysql", "root:nekomeow@/practice")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return AccountRepositoryDb{client: db}
}
