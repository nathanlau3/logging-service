package main

import (
	"fmt"
	"log"
	"minio-service/data"
	"net/http"

	"github.com/joho/godotenv"
)

const webPort = "81"

type Config struct {
	Models data.MinioEntry
}

func main() {
	godotenv.Load(".env")

	app := Config{
		Models: data.New(),
	}

	minioClient := data.New()
	log.Printf("%#v\n", minioClient) // minioClient is now set up
	log.Println("minio is up")

	log.Println("Starting service on port", webPort)
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
