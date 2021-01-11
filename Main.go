package main

import (
	"net/http"
	"fmt"
)

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1> bimeze gte its ur boy <h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}