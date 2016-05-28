package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// S3BucketStatusHandler current status
func S3BucketStatusHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("S3BucketStatusHandler")
	apiApplyCorsHeaders(writer, request)

	type bucketResponse struct {
		Bucket string `json:"bucket"`
	}
	currentBucket := CurrentBucket(db)
	response := bucketResponse{Bucket: currentBucket.Bucket}

	json, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write(json)
}

// S3BucketCreateHandler API create bucket
func S3BucketCreateHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("S3BucketCreateHandler")
	apiApplyCorsHeaders(responseWriter, request)

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
	apiApplyCorsHeaders(responseWriter, request)

	currentBucket := CurrentBucket(db)
	if len(currentBucket.Bucket) == 0 {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("No bucket to refresh"))
	}

	ClearImages(db)
	go DownloadBucket() // download bucket files in a go routine
}
