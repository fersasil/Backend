package main

import (
	_ "github.com/fersasil/backend_ca/app/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	route.addSignHandler(r)

	// r.HandleFunc("/", HomeHandler)

}
