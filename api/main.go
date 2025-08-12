package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"aiw3-nft-api/handlers"
	"aiw3-nft-api/models"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"3000"`
}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Create Chi router
		router := chi.NewMux()

		// Add CORS middleware
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))

		// Add other middleware
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		router.Use(middleware.RequestID)

		// Create Huma API
		config := huma.DefaultConfig("AIW3 NFT API", "1.0.0")
		config.Info.Description = "API for AIW3 NFT system with Solana integration"

		api := humachi.New(router, config)

		// Add OpenAPI info
		api.OpenAPI().Info.Contact = &huma.Contact{
			Name:  "AIW3 Team",
			Email: "dev@aiw3.com",
		}

		// Register all handlers
		handlers.RegisterUserHandlers(api)
		handlers.RegisterNFTHandlers(api)
		handlers.RegisterBadgeHandlers(api)
		handlers.RegisterFeesHandlers(api)
		handlers.RegisterTradingHandlers(api)

		// Add health check endpoint
		huma.Register(api, huma.Operation{
			OperationID: "health-check",
			Method:      "GET",
			Path:        "/health",
			Summary:     "Health check endpoint",
			Tags:        []string{"Health"},
		}, func(ctx context.Context, input *struct{}) (*models.HealthResponse, error) {
			return &models.HealthResponse{
				APIStatus: "ok",
				Message:   "AIW3 NFT API is running",
				Version:   "1.0.0",
			}, nil
		})

		// Setup server
		addr := fmt.Sprintf(":%d", options.Port)
		server := &http.Server{
			Addr:    addr,
			Handler: router,
		}

		hooks.OnStart(func() {
			log.Printf("Starting AIW3 NFT API server on :%d", options.Port)
			log.Printf("OpenAPI docs available at: http://localhost:%d/docs", options.Port)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Server failed to start: %v", err)
			}
		})
	})

	cli.Run()
}
