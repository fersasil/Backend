package main

import (
	"backend_ca/app/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	routes.AddSignHandler(r)

	corsObj := handlers.AllowedOrigins([]string{"*"})

	// http.Handle("/", r)
	err := http.ListenAndServe(":3000", handlers.CORS(corsObj)(r))

	fmt.Println(err)

	// r.HandleFunc("/", HomeHandler)

}
