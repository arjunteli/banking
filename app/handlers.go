package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/arjunteli/banking/service"
)

type Customer struct {
	Name    string `json:"cutomer_name" xml:"customer_name"`
	City    string `json:"cutomer_city" xml:"city"`
	ZipCode string `json:"cutomer_zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	
	//customers
	
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	} else /*if r.Header.Get("Content-Type") == "application/xml"*/ {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}
}
