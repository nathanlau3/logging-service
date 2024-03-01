package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *Config) Routes() http.Handler {
	path, _ := os.Getwd()
	publicDir := fmt.Sprintf("%s/../../public", path)
	
	fs := http.FileServer(http.Dir(publicDir))

	mux := chi.NewRouter()

	//specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	// mux.Post("/", app.WriteLog)
	mux.Handle("/*", http.StripPrefix("/", fs))
	mux.Post("/upload", app.UploadFile)
	mux.Get("/asset", app.GetFile)

	return mux
}