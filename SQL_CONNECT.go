package main

import (
	"database/sql"
	"fmt"
)

func connectDB() (*sql.DB, error) {

	user := "tryituserbase"
	password := "voJ<kCK4~)T.-H,k"
	hostname := "34.159.221.111"
	port := "3306"
	dbname := "TryItHarder"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, hostname, port, dbname)
	dbPool, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("sql.open: " + err.Error())
		return nil, err
	}
	return dbPool, nil
}
