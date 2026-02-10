package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you must upload one file"})
		return
	}

	filename := file.Filename
	filepath := filepath.Join("/home/0x3ea/Kist/data", filename)

	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	newFile := File{
		UUID:     generateUUID(),
		UID:      generateUID(),
		Filename: file.Filename,
		Filepath: filepath,
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
