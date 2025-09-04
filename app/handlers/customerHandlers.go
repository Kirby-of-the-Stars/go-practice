package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/TrueRou/practice/app/errs"
	"github.com/TrueRou/practice/app/repository"
	"github.com/TrueRou/practice/app/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (handlers CustomerHandlers) GetAllCustomers(w http.ResponseWriter, _ *http.Request) {
	customers, _ := handlers.service.GetAllCustomers()
	writeResponse(w, http.StatusOK, customers)
}

func (handlers CustomerHandlers) GetCustomerByName(w http.ResponseWriter, r *http.Request) {
	customer, err := handlers.service.GetCustomerByName(mux.Vars(r)["name"])
	if err != nil {
		var appError errs.AppError
		errors.As(err, &appError)
		writeResponse(w, appError.Code, appError.AsText())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}

func NewCustomerHandlers() CustomerHandlers {
	return CustomerHandlers{service.NewCustomerService(repository.NewCustomerRepositoryDb())}
}
