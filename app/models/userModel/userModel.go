package usermodel

import (
	db "backend_ca/app/helpers/dbHelper"
	"crypto/sha256"
	"encoding/hex"
)

//User ...
type User struct {
	Name     string
	UserID   int
	Password string
	Email    string
	Username string
}

//SignIn ...
func (u User) SignIn(username string, password string) (string, bool) {
	db := db.ConnectDatabase()

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
