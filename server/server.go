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

	r.GET("/article/${id}", controllers.GetArticle)

	r.POST("/article", controllers.CreateArticle)

	r.Run()
}
