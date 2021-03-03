package app

import (
	"github.com/gorilla/mux"
	"banking/domain"
	"banking/service"
	"log"
	"net/http"
)

func Start()  {

	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)


	log.Fatal(http.ListenAndServe(":3000", router))

}

