package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello ")
	})

	r.GET("/haha", func(c *gin.Context) {
		c.String(http.StatusOK, "hahahahha")
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}