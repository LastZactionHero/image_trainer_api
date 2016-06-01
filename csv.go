package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

// ClassificationCsv create a CSV from ImageClassifications
func ClassificationCsv() string {
	// Filename, Strcat, Malloc, Cody, Night, Skip
	// x.jpg,y,y,,,

	// Get all Classifications
	var classifications []Classification
	db.Find(&classifications)

	csvHeaders := []string{"Filename"}
	columnMap := make(map[int64]int)

	// Assign each Classification a position in the CSV
	for idx, classification := range classifications {
		csvHeaders = append(csvHeaders, classification.Name)
		columnMap[classification.ID] = idx
	}

	var sheet [][]string
	sheet = append(sheet, csvHeaders)

	// For each Image
	var images []Image
	db.Find(&images)

	fmt.Println("Column Map:")
	fmt.Println(columnMap)

	// For each Image
	for _, image := range images {
		row := make([]string, len(csvHeaders))
		row[0] = image.Key

		var imageClassifications []ImageClassification
		db.Where(&ImageClassification{ImageID: uint(image.ID)}).Find(&imageClassifications)

		// For each ImageClassification
		for _, imageClassification := range imageClassifications {
			colIdx := columnMap[int64(imageClassification.ClassificationID)]
			row[colIdx+1] = "Y"
		}

		sheet = append(sheet, row)
	}

	b := &bytes.Buffer{}
	w := csv.NewWriter(b)
	for _, row := range sheet {
		w.Write(row)
	}
	w.Flush()

	return b.String()
}
