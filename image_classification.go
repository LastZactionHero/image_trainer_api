package main

import "time"

// ImageClassification relation between Images and Classifications
type ImageClassification struct {
	ID               int64
	CreatedAt        time.Time
	ClassificationID uint
	ImageID          uint
}
