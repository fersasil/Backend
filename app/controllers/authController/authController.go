package authcontroller

import (
	usermodel "backend_ca/app/models/user"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"password"`
	Password string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Returns struct {
	Title       string
	Status      int
	Description string
}

const secretJWT = "ASDJADHASDKFHADS"

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]
	fmt.Println(expectedPassword, ok)

	return
	os.Exit(0)

	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

// SignupHandler ...
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var errors []Returns

	var createUser usermodel.User
	err := json.NewDecoder(r.Body).Decode(&createUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// If user is in use
	if usermodel.UserIsInUse(createUser.Username) {
		errors = append(errors, Returns{Title: "Username already in use.", Status: 400, Description: "Someone is using this user."})
	}

	if errors == nil {
		errors = append(errors, Returns{Title: "Succefully User Created", Status: 200, Description: "User was created succefully"})
		JSON, _ := json.MarshalIndent(errors, "", "\t")
		w.Write(JSON)
		return
	}

	JSON, err := json.MarshalIndent(errors, "", "\t")
	w.Write(JSON)
}
