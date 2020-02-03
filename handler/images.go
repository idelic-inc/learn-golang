package handler

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

func ImageUpload(c *gin.Context, uploader *s3manager.Uploader) {

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("Cannot upload file: %s", err),
		})
	} else {
		err := uploadFile(uploader, file)

		if err != nil {
			c.JSON(400, gin.H{
				"response": fmt.Sprintf("Cannot upload file: %s.", err),
			})
		} else {
			c.JSON(200, gin.H{
				"response": fmt.Sprintf("File uploaded successfully: %s.", file.Filename),
			})
		}

	}
}

func uploadFile(uploader *s3manager.Uploader, fileHeader *multipart.FileHeader) error {

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	bucketName := "golang-learn"
	keyName := fileHeader.Filename

	upParams := &s3manager.UploadInput{
		Bucket: &bucketName,
		Key:    &keyName,
		Body:   file,
	}

	_, err = uploader.Upload(upParams)
	if err != nil {
		return err
	}

	return nil
}
