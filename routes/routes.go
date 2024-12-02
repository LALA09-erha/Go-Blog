package routes

import (
	"Go-Blog/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	category := r.Group("/categories")
	{
		category.GET("/", controllers.GetAllCategories)
		category.POST("/", controllers.CreateCategory) // Auth required
	}

	article := r.Group("/articles")
	{
		article.GET("/", controllers.GetAllArticles)
		article.POST("/", controllers.CreateArticle) // Auth required
	}
}
