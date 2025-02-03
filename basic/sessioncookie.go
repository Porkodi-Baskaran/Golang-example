package basic

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SessionCookiefunc() {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{Path: "/",
		MaxAge: 60 * 60 * 24}) // expire in a day
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	router.GET("/del", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Signed out!...So session cookies got deleted",
		})
	})

	router.Run(":8080")

}
