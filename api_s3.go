package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// S3BucketCreateHandler API create bucket
func S3BucketCreateHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("S3BucketCreateHandler")

	body, _ := ioutil.ReadAll(request.Body)
	var s3Bucket S3Bucket
	err := json.Unmarshal(body, &s3Bucket)

	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	if !s3Bucket.Valid() {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Incomplete request"))
		return
	}

	DeleteBucket(db)
	db.Create(s3Bucket)
	go DownloadBucket() // download bucket files in a go routine

	responseWriter.WriteHeader(http.StatusOK)
}

// S3BucketRefreshHandler refresh files from bucket
func S3BucketRefreshHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("S3BucketRefreshHandler")

	currentBucket := CurrentBucket(db)
	if len(currentBucket.Bucket) == 0 {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("No bucket to refresh"))
	}

	ClearImages(db)
	go DownloadBucket() // download bucket files in a go routine
}
