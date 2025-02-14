package routers

import (
	"example/StudentsDetails/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	// "github.com/gin-gonic/contrib/cors"
	// "github.com/gin-gonic/gin"
	// "github.com/rs/cors/wrapper/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// router.Use(cors.Default())

	// Use CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "http://localhost:3000"
		// },
		MaxAge: 12 * time.Hour,
	}))

	studentRoutes := router.Group("/students")
	{
		studentRoutes.GET("/", controllers.GetStudentDetails)
		studentRoutes.GET("/:id", controllers.GetStudentDetailsbyID)
		studentRoutes.POST("/", controllers.CreateStudent)
		studentRoutes.PUT("/:id", controllers.UpdateStudent)
		studentRoutes.DELETE("/:id", controllers.DeleteStudent)

		//Login & Register User with password hashing
		studentRoutes.POST("/login", controllers.Login)
		studentRoutes.POST("/register", controllers.Register)
	}

	return router
}

// authorized := studentRoutes.Group("/")

// authorized.Use(controllers.AuthMiddleware())
// {
// 	authorized.GET("/users", controllers.GetUsers())
// 	authorized.POST("/users", controllers.CreateUsers())
// }
