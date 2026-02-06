package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDB() {
	dsn := "host=localhost port=5432 user=x3ea password=12346 dbname=kist_db sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//连接数据库
	if err != nil {
		log.Fatalf("Failed to connect db: %v", err)
	}
	log.Println("database connect success")

	//自动迁移
	if err := db.AutoMigrate(&User{}, &File{}); err != nil {
		log.Fatalf("database migrate failed: %v", err)
	}
	log.Println("database migrate success")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)

	log.Println("Successful to connect DB")
}

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you must upload one file"})
		return
	}

	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	filepath := filepath.Join("uploads", filename)

	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	newFile := File{
		UUID:     "123465798",
		UID:      "1437412880",
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

func main() {
	connectDB()

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/upload", uploadHandler)

	fmt.Println("Service run in http://localhost:8080/ping")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Service run failed: %v", err)
	}
}
