package main

// ImageClassification relation between Images and Classifications
type ImageClassification struct {
	ID               int64
	ClassificationID uint
	ImageID          uint
}
