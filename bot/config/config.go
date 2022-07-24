package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Token  string
	Prefix string
	Client *http.Client
	Logger *log.Logger
}

// Global variable
var config *Config

func Setup() error {

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file,", err)
		return err
	}

	config = &Config{
		Token:  os.Getenv("DISCORD_TOKEN"),
		Prefix: os.Getenv("BOT_PREFIX"),
		Client: &http.Client{Timeout: 10 * time.Second},
		Logger: log.New(os.Stdout, "bot ", log.LstdFlags|log.Lshortfile),
	}

	return nil
}

func GetToken() string {
	return config.Token
}

func GetPrefix() string {
	return config.Prefix
}

func GetClient() *http.Client {
	return config.Client
}

func GetLogger() *log.Logger {
	return config.Logger
}
