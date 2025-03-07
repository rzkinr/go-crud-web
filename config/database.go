package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectedDB() {
	// Code to connect to the database
	db, err := sql.Open("mysql", "root:root@/go_products?parseTime=true")
	if err != nil {
		panic(err)
	}

	DB = db
	log.Println("Connected to the database")
}
