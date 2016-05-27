package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func main() {
	fmt.Println("Image Trainer")

	db = dbConnect()
	dbInit()

	r := mux.NewRouter()
	r.HandleFunc("/s3/bucket", S3BucketCreateHandler).Methods("POST")
	r.HandleFunc("/s3/bucket/refresh", S3BucketRefreshHandler).Methods("POST")
	r.HandleFunc("/classifications", ClassificationsCreateHandler).Methods("POST")
	r.HandleFunc("/classifications", ClassificationsIndexHandler).Methods("GET")
	r.HandleFunc("/images/next_file", ImagesNextFileHandler).Methods("GET")
	r.HandleFunc("/images/next_data", ImagesNextDataHandler).Methods("GET")
	r.HandleFunc("/images/remaining", ImagesRemainingHandler).Methods("GET")
	r.HandleFunc("/classify", ClassifyCreateHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("IMAGE_TRAINER_PORT")), nil)
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
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&Classification{})
	db.AutoMigrate(&ImageClassification{})
}
