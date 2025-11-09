package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
}

// configurations is a singleton instance
var configurations *Config

// LoadConfig loads env variables and returns Config instance
func LoadConfig() *Config {
	// Return already loaded config if exists
	if configurations != nil {
		return configurations
	}

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, reading from system environment")
	}

	// Read env variables
	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("VERSION env variable is required")
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("SERVICE_NAME env variable is required")
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		log.Fatal("HTTP_PORT env variable is required")
	}

	// Set singleton config
	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
	}

	fmt.Printf(" Loaded Config: %+v\n", configurations)
	return configurations
}
