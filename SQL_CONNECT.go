package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func connectDB() *sql.DB {
	username := os.Getenv("SQL_USERNAME")
	password := os.Getenv("SQL_PASSWORD")
	database := os.Getenv("SQL_DATABASE")
	port := os.Getenv("SQL_PORT")
	hostname := os.Getenv("SQL_HOSTNAME")

	db, err := sql.Open("mysql",
		username+":"+password+"@tcp("+hostname+":"+port+")/"+database)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return db
}
