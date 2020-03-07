package route

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func addSignHandler(r *mux.Router) {
	r.HandleFunc("/signup", signupGetHandler).Methods("GET")
	// r.HandleFunc("/signup", signupPostHandler).Methods("POST")
	// r.HandleFunc("/signin", signinGetHandler).Methods("GET")
	// r.HandleFunc("/signin", signinPostHandler).Methods("POST")
	// r.HandleFunc("/signout", signoutGetHandler).Methods("GET")
}

func signupGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
