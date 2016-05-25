package main

import "github.com/jinzhu/gorm"

// S3Bucket S3 credentials
type S3Bucket struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
	Bucket string `json:"bucket"`
}

// Valid does the bucket have valid data
func (b S3Bucket) Valid() bool {
	return (len(b.Bucket) > 0 && len(b.Secret) > 0 && len(b.Token) > 0)
}

// CurrentBucket in the database
func CurrentBucket(db *gorm.DB) *S3Bucket {
	var bucket S3Bucket
	db.First(&bucket)
	return &bucket
}

// DeleteBucket deletes the current bucket, if present
func DeleteBucket(db *gorm.DB) bool {
	bucket := CurrentBucket(db)
	if bucket != nil {
		db.Delete(bucket)
		return true
	}
	return false
}
