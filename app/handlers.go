package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/go-delve/delve/service"
)

type Customer struct {
	Name    string `json:"cutomer_name" xml:"customer_name"`
	City    string `json:"cutomer_city" xml:"city"`
	ZipCode string `json:"cutomer_zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers)getAllCustomer(w http.ResponseWriter, r *http.Request) {
	Customers := []Customer{
		{"Arjun Teli", "Bari-Sadri", "312403"},
		{"Karan Teli", "Udaipur", "312403"},
		{"Bhailu Teli", "Ahmedabad", "312403"},
	}
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Customers)

	} else if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(Customers)
	}
}
