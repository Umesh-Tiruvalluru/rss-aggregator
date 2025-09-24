package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port is not found")
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handleReadiness)
	v1Router.Get("/err", handlerErr)

	v1Router.Mount("/v1", v1Router)

	log.Printf("Server running at port %v\n", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
