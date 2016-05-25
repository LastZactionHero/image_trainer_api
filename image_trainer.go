package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

// S3Bucket S3 credentials
type S3Bucket struct {
	Token  string
	Secret string
	Bucket string
}

func main() {
	fmt.Println("Image Trainer")

	db = dbConnect()
	dbInit()
}

func dbConnect() *gorm.DB {
	dbPath := os.Getenv("IMAGE_TRAINER_DB_PATH")
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("failed to connect to database")
	}
	return db
}

func dbInit() {
	db.AutoMigrate(&S3Bucket{})
}
