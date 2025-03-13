package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/yemiwebby/go-cloud-function/app/handler"
)

func main() {
	funcframework.RegisterHTTPFunction("/", handler.HelloWorld)

	// Get the port from environment variable or set a default.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	log.Printf("Starting function on port %s...", port)
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
