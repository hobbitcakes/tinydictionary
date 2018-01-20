package main

import "fmt"
import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()
	// Simulate a ping/pong request
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Simulate a file upload
	// Example: curl -X POST -F "file=@/full/path/to/file.txt" -H "Content-Type: multipart/form-data" http://<hostname>:9999/dinosaurs
	router.POST("/dinosaurs", func(c *gin.Context) {

		// Source
		file, err := c.FormFile("file")
		// Verify request
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprint("get form err: %s", err.Error()))
			return
		}

		// Save the file uploaded to /dev/null
		// Change "/dev/null" to file.Filename to save the file to the local directory
		if err := c.SaveUploadedFile(file, "/dev/null"); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprint("File %s uploaded successfully!", file.Filename))
	})

	router.Run(":9999") // listen and serve on 0.0.0.0:9999
}
