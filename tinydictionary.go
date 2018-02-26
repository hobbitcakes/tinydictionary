package main

import "fmt"
import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	// Create a gin router with the default middleware
	router := gin.Default()

	// Simulate a ping request
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Simulate a ping request
	router.GET("/version", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"version": "archaeopteryx",
		})
	})

	// Simulate a file upload
	// Example: curl -X POST -F "file=@/full/path/to/file.txt" -H "Content-Type: multipart/form-data" http://<hostname>:9999/dinosaurs
	router.POST("/dinosaurs", func(context *gin.Context) {
		// Source
		file, err := context.FormFile("file")
		// Verify request
		if err != nil {
			context.String(http.StatusBadRequest, fmt.Sprint("get form err: %s", err.Error()))
			return
		}

		// Save the file uploaded to /dev/null !note, process saves a temp file to /tmp/ before removing it
		// Change "/dev/null" to file.Filename to save the file to the local directory
		if err := context.SaveUploadedFile(file, "/dev/null"); err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		context.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully!\n", file.Filename))
	})

	router.Run(":9999") // listen and serve on 0.0.0.0:9999
}
