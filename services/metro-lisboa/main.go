package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/DiogoSantoss/kant-bot/services/metro-lisboa/server"
	"github.com/DiogoSantoss/kant-bot/services/metro-lisboa/stations"
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

	mux := http.NewServeMux()

	// Routes
	stations := stations.NewHandlers(client, logger)
	stations.SetupRoutes(mux)

	srv := server.New(mux, os.Getenv("SERVICE_ADDR"))

	err = srv.ListenAndServe() // TODO: Create certificate for ListenAndServeTLS
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

