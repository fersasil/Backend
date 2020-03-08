package main

import (
	"github.com/gorilla/mux"
	"backend_ca/app/routes"
	"net/http"
	"fmt"
)

func main() {

	r := mux.NewRouter()

	routes.AddSignHandler(r)

	// http.Handle("/", r)
	err := http.ListenAndServe(":3000", r)

	fmt.Println(err)

	// r.HandleFunc("/", HomeHandler)

}
