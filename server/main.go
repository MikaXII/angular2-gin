package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {

	router := gin.Default()

	gapi := router.Group("/api")
	{

		gapi.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	}

	/*
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	*/

	router.Run("0.0.0.0:4050")
}
