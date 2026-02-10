package main

import (
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

// TODO: add guard for http://localhost:8080/download?filename=../../../../../etc/passwd
func downloadHandler(c *gin.Context) {
	filename := c.Query("filename")

	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please offer filename"})
		return
	}

	baseDir := "/home/0x3ea/Kist/data"
	filepath := path.Join(baseDir, filename)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not exit"})
		return
	}

	// c.File(filepath)

	c.FileAttachment(filepath, filename)
}
