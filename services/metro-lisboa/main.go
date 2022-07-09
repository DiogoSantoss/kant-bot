package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	logger := log.New(os.Stdout, "metro-lisboa ", log.LstdFlags | log.Lshortfile)

	s := stations.NewHandlers(logger)

	mux := http.NewServeMux()

	s.SetupRoutes(mux)

	srv := server.New(mux, os.Getenv("SERVICE_ADDR"))

	err = srv.ListenAndServe() // TODO: Create certificate for ListenAndServeTLS
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

