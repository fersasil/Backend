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
	sqlStatement := `SELECT * FROM user WHERE user.username=?;`
	row := connectionDB.QueryRow(sqlStatement, username)
	err := row.Scan(&user.UserID, &user.Name, &user.Username, &user.Password, &user.Email)
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
