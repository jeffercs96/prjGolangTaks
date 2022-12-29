package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db sql.DB

const (
	DBtype     = "mysql"
	DBUser     = "root"
	DBPassword = "jeff"
	DBName     = "taskPrj"
)

func CreateCon() *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", DBUser, DBPassword, DBName)
	db, err := sql.Open(DBtype, connectionString)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
