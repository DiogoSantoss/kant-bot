package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/DiogoSantoss/kant-bot/services/metro-lisboa/handlers"
	"github.com/DiogoSantoss/kant-bot/services/metro-lisboa/metro"
	"github.com/DiogoSantoss/kant-bot/services/metro-lisboa/server"
)

func main() {

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file,", err)
	}

	// Dependencies
	client := &http.Client{Timeout: 10 * time.Second}
	logger := log.New(os.Stdout, "metro-lisboa ", log.LstdFlags | log.Lshortfile)

	// Load destinations
	err = metro.LoadDestinations(client)
	if err != nil {
		logger.Printf("error loading destinations: %v", err)
		return
	}

	mux := http.NewServeMux()

	// Routes
	handlers := handlers.NewHandlers(client, logger)
	handlers.SetupRoutes(mux)

	srv := server.New(mux, os.Getenv("SERVICE_ADDR"))

	logger.Println("Server started at " + os.Getenv("SERVICE_ADDR"))

	err = srv.ListenAndServe() // TODO: Create certificate for ListenAndServeTLS
	if err != nil {
		logger.Printf("server failed to start: %v", err)
	}
}

