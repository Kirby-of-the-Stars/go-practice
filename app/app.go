package app

import (
	"log"
	"net/http"

	"github.com/TrueRou/practice/app/handlers"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	handler := handlers.NewCustomerHandlers()
	accountHandlers := handlers.NewAccountHandlers()
	router.HandleFunc("/customers", handler.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{name}", handler.GetCustomerByName).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_name}/accounts", accountHandlers.GetAccountsByOwner).Methods(http.MethodGet)

	router.HandleFunc("/accounts", accountHandlers.GetAllAccounts).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
