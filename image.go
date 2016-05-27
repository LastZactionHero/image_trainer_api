package main

import "time"

// Image represents an image file on S3
type Image struct {
	ID        int64
	CreatedAt time.Time
	Key       string
}
