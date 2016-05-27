package main

import (
	"encoding/json"
	"net/http"
)

// ImagesNextFileHandler get the next image file to classify
func ImagesNextFileHandler(writer http.ResponseWriter, request *http.Request) {
	image := NextImage()
	if image == nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("No more images to classify"))
	}

	buffer, contentType := ReadImage(image)
	writer.Header().Set("Content-Type", contentType)
	writer.Write(buffer)
}

// ImagesNextDataHandler get the next image data to classify
func ImagesNextDataHandler(writer http.ResponseWriter, request *http.Request) {
	image := NextImage()
	if image == nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("No more images to classify"))
	}

	json, err := json.Marshal(image)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.Write(json)
}
