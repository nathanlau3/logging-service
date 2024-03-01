package data

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioEntry struct {
	Minio *minio.Client
}

type Models struct {
	MinioEntry MinioEntry
}

func New() MinioEntry {
	
	endpoint := "localhost:9000"
	accessKeyID := "masWO39Xkj7wQ1QzgPaQ"
	secretAccessKey := "IQlV1nEvMTol7O9wNaEjtoHQNdsFjoIqcZHdXzvI"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return MinioEntry{
		Minio: minioClient,
	}
}

func (m *MinioEntry) UploadFile () {
	// // Make a new bucket called testbucket.
	// bucketName := "nathan-coba"
	// location := "local"
	// err := minioClient.Minio.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	// if err != nil {
	// 	// Check to see if we already own this bucket (which happens if you run this twice)
	// 	exists, errBucketExists := minioClient.Minio.BucketExists(ctx, bucketName)
	// 	if errBucketExists == nil && exists {
	// 		log.Printf("We already own %s\n", bucketName)
	// 	} else {
	// 		log.Fatalln(err)
	// 	}
	// } else {
	// 	log.Printf("Successfully created %s\n", bucketName)
	// }

	// // Upload the test file
	// // Change the value of filePath if the file is in another location
	// objectName := "sample.txt"
	// filePath := "./cmd/api/sample.txt"
	// contentType := "application/octet-stream"

	// // Upload the test file with FPutObject
	// info, err := minioClient.Minio.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}



