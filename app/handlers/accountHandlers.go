package handlers

import (
	"errors"
	"net/http"

	"github.com/TrueRou/practice/app/errs"
	"github.com/TrueRou/practice/app/repository"
	"github.com/TrueRou/practice/app/service"
	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service service.AccountService
}

func (handlers AccountHandlers) GetAllAccounts(w http.ResponseWriter, _ *http.Request) {
	accounts, err := handlers.service.GetAllAccounts()
	if err != nil {
		var appError errs.AppError
		errors.As(err, &appError)
		writeResponse(w, appError.Code, appError.AsText())
	} else {
		writeResponse(w, http.StatusOK, accounts)
	}
}

func (handlers AccountHandlers) GetAccountsByOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer, err := handlers.service.GetAccountsByCustomerName(vars["customer_name"])
	if err != nil {
		var appError errs.AppError
		errors.As(err, &appError)
		writeResponse(w, appError.Code, appError.AsText())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func NewAccountHandlers() AccountHandlers {
	return AccountHandlers{service.NewAccountService(repository.NewAccountRepositoryDb(), repository.NewCustomerRepositoryDb())}
}
