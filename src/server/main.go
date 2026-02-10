package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDB() {
	dsn := "host=localhost port=5432 user=0x3ea password=123456 dbname=kist_db sslmode=disable"

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

func main() {
	snowflakeInit()
	connectDB()

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/upload", uploadHandler)
	r.GET("/download", downloadHandler)

	fmt.Println("Service run in http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Service run failed: %v", err)
	}
}
