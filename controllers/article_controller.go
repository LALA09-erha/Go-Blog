package controllers

import (
	"Go-Blog/config"
	"Go-Blog/models"
	"Go-Blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input struct {
		Title      string `json:"title" binding:"required"`
		Content    string `json:"content" binding:"required"`
		CategoryID string `json:"category_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate slug
	slug := utils.GenerateSlug(input.Title)

	// Create article
	article := models.Article{
		CategoryID: input.CategoryID,
		Title:      input.Title,
		Slug:       slug,
		Content:    input.Content,
		AuthorID:   userID.(string),
	}

	if err := config.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article created successfully"})
}

func GetAllArticles(c *gin.Context) {
	var articles []models.Article
	if err := config.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch articles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"articles": articles})
}
