package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/dinosaurs", func(c *gin.Context) {
		name := "Success"
		c.String(http.StatusOK, name)
	})

	router.Run(":9999") // listen and serve on 0.0.0.0:9999
}
