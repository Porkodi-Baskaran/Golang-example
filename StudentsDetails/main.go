package main

import (
	"database/sql"
	config "example/StudentsDetails/Config"
	"example/StudentsDetails/routers"

	"github.com/gin-gonic/contrib/static"
)

var db *sql.DB

func main() {
	config.Dbconnection()
	// config.LoginConnection()
	router := routers.SetupRouter()
	router.Use(static.Serve("/", static.LocalFile("./studentapp/build/", true)))

	// router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.Run(":8080")

}
