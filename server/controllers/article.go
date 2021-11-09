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

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	models.DB.First(&article, id)
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func UpdateArticle(c *gin.Context) {
	var input models.CreateArticleInput
	var article models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&article).Updates(&input)
}

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

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	models.DB.Delete(&article, id)
	c.JSON(http.StatusOK, gin.H{"article": article})

}
