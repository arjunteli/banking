package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/arjunteli/banking/domain"
	"github.com/arjunteli/banking/service"
)

func Start() {
	//define multiplexer
	router := mux.NewRouter()

	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)

	//starting http server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
