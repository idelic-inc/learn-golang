package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/idelic-inc/learn-golang/aws"
	"github.com/idelic-inc/learn-golang/handler"
)

func main() {
	r := gin.Default()

	uploader, err := aws.NewS3Uploader()

	if err != nil {
		log.Printf("Could not connect to S3: %s\n", err)
		os.Exit(1)
	}

	// Routes
	r.POST("/images/upload", func(c *gin.Context) {
		handler.ImageUpload(c, uploader)
	})

	// Listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}
