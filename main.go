package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
}

// getEnvVars returns a formatted string of all environment variables
func getEnvVars() string {
	envVars := os.Environ()
	sort.Strings(envVars)
	return strings.Join(envVars, ", ")
}

// getEnvVar gets an environment variable with a default value
func getEnvVar(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthResponse{Status: "ok"})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello from Go API")
}

func main() {
	port := getEnvVar("PORT", "8080")

	// Log environment variables at startup
	log.Printf("Server starting on port %s", port)
	log.Printf("Environment variables: %s", getEnvVars())

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", rootHandler)

	// Start logging goroutine
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				log.Printf("Server is running on port %s | Env vars: %s", port, getEnvVars())
			}
		}
	}()

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

