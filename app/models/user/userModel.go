package usermodel

import (
	db "backend_ca/app/helpers/connecthelper"
	"database/sql"
)

// User ...
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	UserID   int    `json:"userID"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserIsInUse ...
func UserIsInUse(username string) bool {
	var user User
	connectionDB := db.ConnectDatabase()
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
	connectionDB := db.ConnectDatabase()
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
	connectionDB := db.ConnectDatabase()
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
