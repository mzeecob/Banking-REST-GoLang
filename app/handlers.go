package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"banking/service"
	"net/http"
)

type Customer struct{
	Name string	`json:"full_name"`
	City string	`json:"city"`
	Zipcode string	`json:"zip_code"`
}

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1> bimeze gte its ur boy mzee <h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "please send email to get in touch <a href=\"mzeecob@gmail.com\"> mzeecob@gmail.com</a>")
	}else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "this page doesn't exit")
	}

}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request)  {

	//customers := []Customer{
	//	{ Name: "Regis", City: "Kigali", Zipcode: "045415" },
	//	{ Name: "Regis", City: "Kigali", Zipcode: "045415" },
	//}

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("content-type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)

	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

