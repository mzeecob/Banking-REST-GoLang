package main

import (
	"net/http"
	"fmt"
)

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1> bimeze gte its ur boy mzee <h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "please send email to get in touch <a href=\"mzeecob@gmail.com\"> mzeecob@gmail.com</a>")

	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}