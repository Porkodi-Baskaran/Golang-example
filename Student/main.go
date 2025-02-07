package main

import (
	"database/sql"
	"example/Student/DB"
	"fmt"
)

var db *sql.DB

func main() {

	fmt.Println("DB connection starts here")
	db = DB.Dbconnection()

	fmt.Println(DB.RetrieveDatas(db, "Tom"))

}
