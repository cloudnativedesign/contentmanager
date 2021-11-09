package server

import (
	"github.com/cloudnativedesign/contentmanager/server/controllers"
	"github.com/cloudnativedesign/contentmanager/server/models"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/articles", controllers.GetArticles)

	r.GET("/article/:id", controllers.GetArticleById)

	r.POST("/article", controllers.CreateArticle)

	r.POST("/article/:id", controllers.UpdateArticle)

	r.DELETE("/article/:id", controllers.DeleteArticle)

	r.Run()
}
