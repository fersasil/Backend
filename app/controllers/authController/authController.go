package authcontroller

import (
	jwt "backend_ca/app/helpers/jwtHelper"
	user "backend_ca/app/models/userModel"
	"encoding/json"
	"net/http"
	"strings"
)

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

	var u user.User

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
