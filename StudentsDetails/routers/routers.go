package routers

import (
	"example/StudentsDetails/controllers"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	// "github.com/gin-gonic/contrib/cors"
	// "github.com/gin-gonic/gin"
	// "github.com/rs/cors/wrapper/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{Path: "/",
		MaxAge: 60 * 60 * 24}) // expire in 24 hrs

	fmt.Println(store)
	router.Use(sessions.Sessions("mysession", store))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                                    // Allow your React frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", http.MethodHead}, // Allow your desired HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},        // Allow these headers
		ExposeHeaders:    []string{"content-length"},
		AllowCredentials: true, // Allow cookies or credentials
	}))

	// router.Use(static.Serve("/", static.LocalFile("./studentapp/build/", true)))

	studentRoutes := router.Group("/api")
	{
		studentRoutes.GET("/students/", controllers.AuthRequired, controllers.GetStudentDetails)
		studentRoutes.GET("/students/:id", controllers.GetStudentDetailsbyID)
		studentRoutes.POST("/students/", controllers.CreateStudent)
		studentRoutes.PUT("/students/:id", controllers.UpdateStudent)
		studentRoutes.DELETE("/students/:id", controllers.DeleteStudent)

		//Login & Register User with password hashing
		studentRoutes.POST("/login", controllers.Login)
		studentRoutes.POST("/register", controllers.Register)
		studentRoutes.GET("/auth/google/login", controllers.HandleGoogleLogin)
		studentRoutes.GET("/logout", controllers.LogoutUser)
		studentRoutes.GET("/auth/google/callback", controllers.HandleGoogleCallback)
		studentRoutes.GET("/marks", controllers.GetStudMarks)
		studentRoutes.GET("/marks/:id", controllers.GetStudMarksbyID)

		// studentRoutes.GET("/enable-2fa/:username", controllers.GenerateTOTPSecret)
		studentRoutes.POST("/verify-2fa", controllers.VerifyTOTP)

	}

	return router
}

// router.Use(cors.Default())

// Use CORS middleware
// router.Use(cors.New(cors.Config{
// 	AllowOrigins:     []string{"http://localhost:3000"},
// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 	AllowHeaders:     []string{"Origin", "Content-Type"},
// 	ExposeHeaders:    []string{"Content-Length"},
// 	AllowCredentials: true,
// 	// AllowOriginFunc: func(origin string) bool {
// 	// 	return origin == "http://localhost:3000"
// 	// },
// 	MaxAge: 12 * time.Hour,
// }))
