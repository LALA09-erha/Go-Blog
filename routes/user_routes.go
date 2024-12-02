package routes

import (
	"Go-Blog/controllers"
	"Go-Blog/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	adminRoutes := router.Group("/users")
	adminRoutes.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminRoutes.GET("/", controllers.GetAllUsers)
		adminRoutes.GET("/:id", controllers.GetUserByID)
		adminRoutes.POST("/", controllers.CreateUser)
		adminRoutes.PUT("/:id", controllers.UpdateUser)
		adminRoutes.DELETE("/:id", controllers.DeleteUser)
	}
}
