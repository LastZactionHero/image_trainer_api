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
	DownloadBucket()

	responseWriter.WriteHeader(http.StatusOK)
}
