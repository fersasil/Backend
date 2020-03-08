package authcontroller

import (
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
	Password string `json:"username`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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

// func createJWT(string userId) {
// 	jwtKey := []byte(secretJWT)
// }
