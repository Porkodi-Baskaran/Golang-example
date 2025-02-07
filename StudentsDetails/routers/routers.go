package routers

import (
	"example/StudentsDetails/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	studentRoutes := router.Group("/students")
	{
		studentRoutes.GET("/", controllers.GetStudentDetails)
		studentRoutes.GET("/:id", controllers.GetStudentDetailsbyID)
		studentRoutes.POST("/", controllers.CreateStudent)
		studentRoutes.PUT("/:id", controllers.UpdateStudent)
		studentRoutes.DELETE("/:id", controllers.DeleteStudent)
	}

	return router
}
