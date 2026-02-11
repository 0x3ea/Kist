package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func downloadHandler(c *gin.Context) {
	filename := c.Query("filename")

	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please offer filename"})
		return
	}

	baseDir := "/home/0x3ea/Kist/data"

	downloadPath, err := secureJoin(baseDir, filename)

	if err != nil {
		log.Printf("illegal uploadPath, input: %s,error: %v", filename, err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
		return
	}

	if _, err := os.Stat(downloadPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not exit"})
		return
	}

	c.FileAttachment(downloadPath, filename)
}
