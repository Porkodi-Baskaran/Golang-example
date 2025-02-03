package basic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CookieExample() {
	router := gin.Default()
	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)

		}

		fmt.Printf("Cookie value: %s \n", cookie)
		if cookie != "" {
			c.SetCookie("gin_cookie", "", -1, "/", "localhost", false, true)
			c.String(http.StatusOK, "Cookie has been deleted")
			fmt.Println("gin Cookie deleted")
			fmt.Printf("Cookie value: %s \n", cookie)

		}
	})

	router.Run()
}

// func DeleteCookieHandler(var c){
// 	c.SetCookie("user", "", -1, "/", "localhost", false, true)
// 	c.String(http.StatusOK, "Cookie has been deleted")
// }
