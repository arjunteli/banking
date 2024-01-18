package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//define multiplexer
	//mux:= http.NewServeMux()
	mux := mux.NewRouter()
	//define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomer)
	//starting http server
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
