package main

import (
	"backend_ca/app/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	routes.AddSignHandler(r)

	// http.Handle("/", r)
	err := http.ListenAndServe(":3000", r)

	fmt.Println(err)

	// r.HandleFunc("/", HomeHandler)

}
