package http

import (
	"github.com/gin-gonic/gin"
)

func GinAPIServer() {

	router := gin.Default()

	router.GET("/", getHandler)
	router.POST("/", postHandler)

	router.Run(":8080")
}

func getHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Gin!",
	})
}

type User struct {
	Name string `json:"name"`
	Age  int    `json: "age"`
	Pass bool   `json: "pass"`
}
