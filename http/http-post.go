package http

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func postHandler(c *gin.Context) {
	var user = new(User) // User struct with appropriate fields

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User: ", user)

	if user.Pass {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// Validate and process user data
	// Create new user in the database
	c.JSON(200, gin.H{
		"message": "User created successfully!",
	})
}
