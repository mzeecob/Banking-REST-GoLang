package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Customer struct{
	Name string
	City string
	Zipcode string
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

func getAllCustomer(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	customers := []Customer {
		{ Name: "Regis", City: "Kigali", Zipcode: "045415" },
		{ Name: "Regis", City: "Kigali", Zipcode: "045415" },
	}

	json.NewEncoder(w).Encode(customers)

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/customers", getAllCustomer)
	http.ListenAndServe(":3000", nil)
}