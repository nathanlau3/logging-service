package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)

func (c *Config) UploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Access the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a destination file
	path, err := os.Getwd()
	dst, err := os.Create(filepath.Join(fmt.Sprintf("%s/temp", path), handler.Filename))
	if err != nil {
		http.Error(w, "Unable to create destination file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Unable to copy file to destination", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)

	// Make a new bucket called testbucket.
	bucketName := "nathan-coba"
	location := "local"
	err = c.Models.Minio.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := c.Models.Minio.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the test file
	// Change the value of filePath if the file is in another location
	objectName := handler.Filename
	filePath := fmt.Sprintf("%s/temp/%s", path, handler.Filename)
	contentType := "application/octet-stream"

	// Upload the test file with FPutObject
	info, err := c.Models.Minio.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}

func (c *Config) GetFile(w http.ResponseWriter, r *http.Request) {
	// Replace "your_file_path" with the actual path of the file you want to serve
	path, _ := os.Getwd()

	filePath := fmt.Sprintf("%s/temp/sampel.txt", path)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set the appropriate headers
	w.Header().Set("Content-Disposition: inline", "attachment; filename="+"peakpx (1).jpg")
	// w.Header().Set("Content-Type", "application/octet-stream")

	// Copy the file content to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Unable to copy file content to response", http.StatusInternalServerError)
		return
	}
}
