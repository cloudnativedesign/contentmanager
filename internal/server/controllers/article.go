package controllers

import (
	"net/http"

	"github.com/cloudnativedesign/contentmanager/server/models"

	"github.com/gin-gonic/gin"
)

// GET /articles
func GetArticles(c *gin.Context) {
	var articles []models.Article
	models.DB.Find(&articles)

	c.JSON(http.StatusOK, gin.H{"data": articles})
}

func GetArticle(c *gin.Context) {
	var articles []models.Article
	models.DB.Find(&articles)

	c.JSON(http.StatusOK, gin.H{"data": articles})
}

// POST /article
func CreateArticle(c *gin.Context) {
	var input models.CreateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create book after successful verification of input
	article := models.Article{Title: input.Title, Author: input.Author}
	models.DB.Create(&article)

	c.JSON(http.StatusOK, gin.H{"data": article})
}
