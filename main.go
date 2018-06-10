package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// set the router to Gin's default
	router := gin.Default()

	// serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// setup route grup for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// start the server
	router.Run(":3000")
}
