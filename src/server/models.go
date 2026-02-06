package main

import (
	"time"
)

type User struct {
	UID          string    `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"uniqueIndex;not null" json:"username"`
	Password     string    `gorm:"not null" json:"-"`
	RegisterTime time.Time `json:"register_time"`
}

type File struct {
	UUID       string    `gorm:"primaryKey" json:"uuid"`
	UID        string    `gorm:"not null" json:"uid"`
	Filename   string    `gorm:"not null" json:"file_name"`
	Filepath   string    `gorm:"not null" json:"file_path"`
	Filesize   int64     `gorm:"not null" json:"file_size"`
	Uploadtime time.Time `json:"upload_time"`
}
