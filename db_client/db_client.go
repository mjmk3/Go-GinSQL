package db_client

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DBClient *sql.DB

func InitializeDBConnect() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_connect_db?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	//pinging DB to make sure it's work correct
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db
}
