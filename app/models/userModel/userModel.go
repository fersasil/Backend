package usermodel

import (
	dbHelper "backend_ca/app/helpers/connectHelper"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
)

// User ...
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	UserID   int    `json:"userID"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

//SignIn ...
func (u User) SignIn(username string, password string) (string, bool) {
	db := dbHelper.ConnectDatabase()

	hash := sha256.New()

	hash.Write([]byte(password))

	hashedPassword := hex.EncodeToString(hash.Sum(nil))

	sqlStatement := `SELECT id FROM user WHERE user.username=? and password = ?;`

	row := db.QueryRow(sqlStatement, username, hashedPassword)

	var userID string

	err := row.Scan(&userID)

	if err != nil {
		return "", false
	}

	return userID, true
}

// UserIsInUse ...
func UserIsInUse(username string) bool {
	var user User
	connectionDB := dbHelper.ConnectDatabase()
	sqlStatement := `SELECT username FROM user WHERE user.username=?;`
	row := connectionDB.QueryRow(sqlStatement, username)
	connectionDB.Close()
	err := row.Scan(&user.Username)
	switch err {
	case sql.ErrNoRows:
		//fmt.Println("No rows were returned!")
		return false // User isn't in use
	case nil:
		return true // User is in use
	default:
		panic(err)
	}
}

// EmailIsInUse ...
func EmailIsInUse(email string) bool {
	var user User
	connectionDB := dbHelper.ConnectDatabase()
	sqlStatement := `SELECT email FROM user WHERE user.email=?;`
	row := connectionDB.QueryRow(sqlStatement, email)
	connectionDB.Close()
	err := row.Scan(&user.Email)
	switch err {
	case sql.ErrNoRows:
		//fmt.Println("No rows were returned!")
		return false // email isn't in use
	case nil:
		return true // email is in use
	default:
		panic(err)
	}
}

// CreateUser ...
func CreateUser(username, name, password, email string) bool {
	connectionDB := dbHelper.ConnectDatabase()
	insForm, err := connectionDB.Prepare("INSERT INTO user (username, name, password, email) VALUES (?, ?, ?, ?);")
	if err != nil {
		panic(err.Error())
		connectionDB.Close()
		return false
	}
	insForm.Exec(username, name, password, email)
	connectionDB.Close()
	return true
}
