package DB

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/mattn/go-sqlite3"
)

type AdminUser struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func Dbconnection() (db *sql.DB) {

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
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Create table if it doesn't exist
	query := `CREATE TABLE IF NOT EXISTS student (
        ID INT AUTO_INCREMENT PRIMARY KEY,
        Name VARCHAR(100) NOT NULL,
        Class INT NOT NULL,
        Address VARCHAR(255) NOT NULL
    );`
	if _, err := db.Exec(query); err != nil {
		panic(fmt.Sprintf("Failed to create table: %s", err.Error()))
	}
	return db

}
