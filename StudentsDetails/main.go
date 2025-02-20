package main

import (
	config "example/StudentsDetails/Config"
	"example/StudentsDetails/routers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Dbconnection()
	// config.LoginConnection()
	router := routers.SetupRouter()

	router.Use(static.Serve("/", static.LocalFile("./studentapp/build", false)))

	// Serve the React app at the root URL
	router.NoRoute(func(c *gin.Context) {
		c.File("./studentapp/build/index.html")
	})

	router.Run(":8080")

}
