package connecthelper

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //
)

// ConnectDatabase ...
func ConnectDatabase() *sql.DB {
	//sql.Open("mysql", "user:password@tcp(host:port)/database")
	connectionDB, err := sql.Open("mysql", "teste:123@tcp(127.0.0.1:3306)/dbTeste")
	if err != nil {
		panic(err.Error())
	}
	return connectionDB
}
