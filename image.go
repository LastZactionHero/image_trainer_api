package main

import (
	"bytes"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Image represents an image file on S3
type Image struct {
	ID         int64
	CreatedAt  time.Time
	Key        string
	Classified bool `gorm:"default:0"`
}

// FindImageByKey find image by key
func FindImageByKey(key string) *Image {
	var image Image
	db.Where("key = ?", key).First(&image)
	if image.ID == 0 {
		return nil
	}
	return &image
}

// NextImage return the next image to classify
func NextImage() *Image {
	var image Image
	db.Where("classified = ?", 0).First(&image)
	if image.ID == 0 {
		return nil
	}
	return &image
}

// ReadImage from S3
func ReadImage(image *Image) ([]byte, string) {
	ApplyBucketAccess()
	bucket := CurrentBucket(db)

	svc := s3.New(session.New(), &aws.Config{Region: aws.String("us-west-1")})

	params := &s3.GetObjectInput{
		Bucket: &bucket.Bucket,
		Key:    &image.Key,
	}
	resp, err := svc.GetObject(params)

	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer

	for {
		readBuffer := make([]byte, *resp.ContentLength)
		bytesRead, _ := resp.Body.Read(readBuffer)
		if bytesRead == 0 {
			break
		}
		buffer.Write(readBuffer[0:bytesRead])
	}

	return buffer.Bytes(), *resp.ContentType
}
