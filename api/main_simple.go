package main

import (
	"context"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type HealthResponse struct {
	APIStatus string `json:"api_status"`
	Message   string `json:"message"`
	Version   string `json:"version"`
}

func main() {
	// Create Chi router
	router := chi.NewMux()

	// Add CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Add other middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Create Huma API
	config := huma.DefaultConfig("AIW3 NFT API", "1.0.0")
	config.Info.Description = "API for AIW3 NFT system with Solana integration"

	api := humachi.New(router, config)

	// Add a simple health check endpoint
	huma.Register(api, huma.Operation{
		OperationID: "health-check",
		Method:      "GET",
		Path:        "/health",
		Summary:     "Health check endpoint",
		Tags:        []string{"Health"},
	}, func(ctx context.Context, input *struct{}) (*HealthResponse, error) {
		return &HealthResponse{
			APIStatus: "ok",
			Message:   "AIW3 NFT API is running",
			Version:   "1.0.0",
		}, nil
	})

	// Start server
	log.Printf("Starting AIW3 NFT API server on :3000")
	log.Printf("OpenAPI docs available at: http://localhost:3000/docs")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
