package app

import (
	"log"
	"net/http"

	"github.com/arjunteli/banking/domain"
	"github.com/arjunteli/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	//define multiplexer
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.newCustomerRepositoryStub())}

	//define routes
	router.HandleFunc("/customers", getAllCustomer).Methods(http.MethodGet)

	//starting http server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
