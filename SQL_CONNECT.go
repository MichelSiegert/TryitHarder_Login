package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func connectDB() (*sql.DB, error) {
	user := os.Getenv("SQL_USERNAME")
	password := os.Getenv("SQL_PASSWORD")
	hostname := os.Getenv("SQL_DATABASE")
	port := os.Getenv("SQL_PORT")
	dbname := os.Getenv("SQL_HOSTNAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, hostname, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("sql.open: " + err.Error())
		return nil, err
	}
	return db, nil
}
