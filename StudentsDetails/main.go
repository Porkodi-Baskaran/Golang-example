package main

import (
	config "example/StudentsDetails/Config"
	"example/StudentsDetails/routers"
)

func main() {
	config.Dbconnection()
	router := routers.SetupRouter()
	router.Run(":8080")

}
