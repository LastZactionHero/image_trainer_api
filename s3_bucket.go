package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jinzhu/gorm"
)

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
		ClearImages(db)
		db.Delete(bucket)
		return true
	}
	return false
}

// ClearImages delete all Images and ImageClassifications
func ClearImages(db *gorm.DB) {
	db.Where("id > 0", 0).Delete(&Image{})
	db.Where("id > ?", 0).Delete(&ImageClassification{})
}

// ApplyBucketAccess token and secret to env
func ApplyBucketAccess() {
	bucket := CurrentBucket(db)
	os.Setenv("AWS_ACCESS_KEY_ID", bucket.Token)
	os.Setenv("AWS_SECRET_ACCESS_KEY", bucket.Secret)
}

// DownloadBucket download all bucket files
func DownloadBucket() {
	ApplyBucketAccess()
	bucket := CurrentBucket(db)

	sess := session.New()

	svc := s3.New(sess, &aws.Config{Region: aws.String("us-west-1")})

	err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: &bucket.Bucket,
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		for _, obj := range p.Contents {
			image := Image{Key: *obj.Key}
			db.Create(&image)
		}
		return true
	})
	if err != nil {
		panic(err)
	}
}
