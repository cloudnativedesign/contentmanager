package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world!"})
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Some test data"})
	})

	r.Run()
}
