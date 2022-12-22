package db_client

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

func InitializeDBConnect() {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/go_connect_db?parseTime=true")
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
