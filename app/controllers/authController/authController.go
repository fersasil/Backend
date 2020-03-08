package authcontroller

import (
	jwt "backend_ca/app/helpers/jwtHelper"
	usermodel "backend_ca/app/models/userModel"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

// Returns ...
type Returns struct {
	Title       string
	Status      int
	Description string
}

func sendUnauthorizedError(w http.ResponseWriter, statusCode int) {
	res := make(map[string]string)

	res["error"] = "User or password wrong"
	resJSON, err := json.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(resJSON)
}

//SigninHandler ...
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	if queryParams["username"] == nil || queryParams["password"] == nil {
		sendUnauthorizedError(w, http.StatusUnauthorized)
		return
	}

	var u usermodel.User

	//convert []string to string
	password := strings.Join(queryParams["password"], " ")
	username := strings.Join(queryParams["username"], " ")

	userID, ok := u.SignIn(username, password)

	if !ok {
		sendUnauthorizedError(w, http.StatusUnauthorized)
		return
	}

	token, ok := jwt.CreateJWT(userID, 60)

	res := make(map[string]string)

	res["token"] = token

	resJSON, err := json.Marshal(res)

	if err != nil {
		sendUnauthorizedError(w, http.StatusInternalServerError)
		return
	}

	w.Write(resJSON)
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
