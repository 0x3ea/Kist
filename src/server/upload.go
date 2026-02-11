package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you must upload one file"})
		return
	}
	baseDir := "/home/0x3ea/Kist/data"

	uploadPath, err := secureJoin(baseDir, file.Filename)

	if err != nil {
		log.Printf("illegal uploadPath, input: %s,error: %v", file.Filename, err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
		return
	}

	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	newFile := File{
		UUID:     generateUUID(),
		UID:      generateUID(),
		Filename: file.Filename,
		Filepath: uploadPath,
		Filesize: file.Size,
	}

	if result := db.Create(&newFile); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload file success",
		"file":    newFile,
	})
}
