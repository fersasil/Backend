package authcontroller

import (
	usermodel "backend_ca/app/models/user"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

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

// SigninHandler ...
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
		errors = append(errors, Returns{Title: "Username: `" + createUser.Username + "` is already in use.", Status: 409, Description: "Someone is using this user."})
	}

	// If email is Valid
	if EmailIsValid(createUser.Email) {
		// Verify if email already is in database
		if usermodel.EmailIsInUse(createUser.Email) {
			errors = append(errors, Returns{Title: "Email: `" + createUser.Email + "` is already in use.", Status: 409, Description: "Someone is using this email."})
		}
	} else {
		errors = append(errors, Returns{Title: "Email: `" + createUser.Email + "` isn't a valid email.", Status: 400, Description: "Verify if you typed your email address correctly."})
	}

	if IsStrongPassword(createUser.Password) {
		// If password is strong, encrypt
		hash := sha256.New()
		hash.Write([]byte(createUser.Password))
		createUser.Password = hex.EncodeToString(hash.Sum(nil))
	} else {
		errors = append(errors, Returns{Title: "Password Invalid", Status: 417, Description: "Password need have at least one uppercase and on lowercase letter and a number, and need have 8 characters"})
	}

	if errors == nil {
		if usermodel.CreateUser(createUser.Username, createUser.Name, createUser.Password, createUser.Email) {
			errors = append(errors, Returns{Title: "Succefully User Created.", Status: 201, Description: "User was created succefully"})
		} else {
			errors = append(errors, Returns{Title: "User create fail.", Status: 500, Description: "Something gones wrong on data insert in database."})
		}
		JSON, _ := json.MarshalIndent(errors, "", "\t")
		w.Write(JSON)
		return
	}

	JSON, err := json.MarshalIndent(errors, "", "\t")
	w.Write(JSON)
}

// EmailIsValid ...
func EmailIsValid(email string) bool {
	Re := regexp.MustCompile(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`)
	return Re.MatchString(email)
}

// IsStrongPassword ...
func IsStrongPassword(password string) bool {
	Re := regexp.MustCompile(`(^[a-zA-Z0-9]+$)`)
	if Re.MatchString(password) && len(password) >= 8 {
		return true
	}
	return false
}
