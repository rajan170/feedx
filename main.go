package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not found")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	fmt.Print(v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Failed to start server", err)
	} else {
		log.Printf("Server started at port: {%v}", port)
	}

	fmt.Printf("Server started at port: {%v}", port)

	fmt.Println("Port: ", port)

}
