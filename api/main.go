package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"
)

// stringPtr returns a pointer to a string
func stringPtr(s string) *string {
	return &s
}

func main() {
	// Create service with OpenAPI documentation
	service := web.NewService(openapi3.NewReflector())

	// Configure OpenAPI metadata
	service.OpenAPISchema().SetTitle("AIW3 NFT API")
	service.OpenAPISchema().SetDescription("Mock API for AIW3 NFT Solana system - provides same data structures as Node.js API")
	service.OpenAPISchema().SetVersion("1.0.0")

	// Note: Error responses follow original API pattern { code, message, data }
	// All endpoints return consistent 3-field structure matching lastmemefi-api

	// Add server info
	spec := service.OpenAPISchema().(*openapi3.Spec)
	spec.WithServers(openapi3.Server{
		URL:         "http://localhost:8080",
		Description: stringPtr("Development server"),
	})

	// Setup middlewares
	service.Wrap(
		gzip.Middleware, // Response compression
		// Add CORS middleware
		func(handler http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")

				if r.Method == "OPTIONS" {
					w.WriteHeader(http.StatusOK)
					return
				}

				handler.ServeHTTP(w, r)
			})
		},
	)

	// Register NFT and Badge endpoints
	setupAPIRoutes(service)

	// Swagger UI endpoint at /docs
	service.Docs("/docs", swgui.New)

	// Start server
	fmt.Println("🚀 Starting AIW3 NFT API server on :8080")
	fmt.Println("📚 API Documentation: http://localhost:8080/docs")

	if err := http.ListenAndServe(":8080", service); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
