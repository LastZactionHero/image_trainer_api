# Image Trainer API

API for classifying images for use with a neural network. Handles coordination with S3 and classification interface.

## Environmental Variables

- IMAGE_TRAINER_DB_USER - Mysql DB User
- IMAGE_TRAINER_DB_PASS - Mysql DB Password
- IMAGE_TRAINER_DB_NAME - Mysql DB Name
- IMAGE_TRAINER_PORT - Server Port

## API

#### GET /s3/bucket/status
Get current status of S3 Bucket

Body: None

Response: 200
```
{
  bucket: "my_bucket_name"
}
```

#### POST /s3/bucket
Create the S3 Bucket and create a record for all Images

Body:
```
{
  "token": "<AWS IAM KEY >"
  "secret": "<AWS IAM SECRET>",
  "bucket": "<AWS BUCKET NAME >",
}
```

Response: 200

#### POST /s3/bucket/refresh
Refresh Images from the S3 Bucket

Body: None

Response: 200

#### POST /classifications
Create a new Classifications

Body:
```
{
  "name": "Malloc",
  "hotkey": "m"
}
```

Response: 200

#### GET /classifications
Get a list of all classifications

Body: None

Response: 200
```
[
{"ID":1,"name":"strcat","hotkey":"c"},
{"ID":2,"name":"cody",  "hotkey":"d"},
{"ID":3,"name":"malloc","hotkey":"m"}
]
```

#### GET /images/next_data
Get metadata for the next Image to classify

Body: None

Response: 200
```
{
  "id":3,
  "key":"image_1463849668.jpg",
  "classified":false
}
```
#### GET /images/next_file
Get image for the next Image to classify

Body: None

Response: 200
Image File

#### GET /images/remaining
Number of images remaining to classify

Body: None

Response: 200
```
{
  count: 1025
}
```

#### POST /classify
Classify the image

Body:
```
{
  "key": "image_1463849652.jpg",
  "classifications": [
    "malloc", "strcat"
  ]
}
```

Response: 200

#### GET /csv
Get CSV output of all image classifications

Body: None

Response: 200
CSV File