package app

import "net/http"

func Start()  {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/customers", getAllCustomer)
	http.ListenAndServe(":3000", nil)
}
