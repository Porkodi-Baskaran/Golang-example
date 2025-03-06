package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Dbconnection() {
	var db *sql.DB
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "test123",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "student_db",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic("failed to connect database")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to MySQL Database!")
	DB = db

	// Create table if it doesn't exist
	query := `CREATE TABLE IF NOT EXISTS student (
        ID INT AUTO_INCREMENT PRIMARY KEY,
        Name VARCHAR(100) NOT NULL,
        Class INT NOT NULL,
        Address VARCHAR(255) NOT NULL
    );`
	if _, err := DB.Exec(query); err != nil {
		panic(fmt.Sprintf("Failed to create table: %s", err.Error()))
	} else {
		fmt.Println("Table created or already Exists")
	}

}
