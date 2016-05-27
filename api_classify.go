package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ClassifyCreateHandler create classifications for an image
func ClassifyCreateHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Classify Create Handler")

	// Parse Request Body
	// key:
	// classifications: [name, name, name]
	type classifyJSON struct {
		Key             string   `json:"key"`
		Classifications []string `json:"classifications"`
	}
	var requestParams classifyJSON
	jsonBuffer := make([]byte, request.ContentLength)
	request.Body.Read(jsonBuffer)
	err := json.Unmarshal(jsonBuffer, &requestParams)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the image by key
	image := FindImageByKey(requestParams.Key)
	if image == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if image.Classified {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Image already classified"))
		return
	}

	// Find the classifications
	var classifications []*Classification
	for _, name := range requestParams.Classifications {
		classification := FindClassificationByName(name)
		if classification == nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("Could not find classification"))
			return
		}
		classifications = append(classifications, classification)
	}

	// Create Classifications
	for _, classification := range classifications {
		imageClassification := ImageClassification{ClassificationID: uint(classification.ID), ImageID: uint(image.ID)}
		db.Create(&imageClassification)
	}

	// Mark image as classified
	image.Classified = true
	db.Save(image)
}
