package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ClassificationsCreateHandler create a Classification
func ClassificationsCreateHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Classifications Create Handler")
	apiApplyCorsHeaders(writer, request)

	body, _ := ioutil.ReadAll(request.Body)
	var classification Classification
	err := json.Unmarshal(body, &classification)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if !classification.Valid() {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Incomplete request"))
		return
	}

	if FindClassificationByName(classification.Name) != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("A Classification with this name already exists"))
		return
	}
	if FindClassificationByHotkey(classification.Hotkey) != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("A Classification with this hotkey already exists"))
		return
	}

	db.Create(&classification)
}

// ClassificationsIndexHandler list all Classifications
func ClassificationsIndexHandler(writer http.ResponseWriter, request *http.Request) {
	apiApplyCorsHeaders(writer, request)

	var classifications []Classification
	db.Find(&classifications)

	json, err := json.Marshal(classifications)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write(json)
}
