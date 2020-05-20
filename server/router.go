package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter makes routes
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	return router
}
