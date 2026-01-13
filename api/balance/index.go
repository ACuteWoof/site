package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers first
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 25 * time.Second,
	}

	// Make the request
	res, err := client.Get("http://70.34.201.17/log.csv")
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		http.Error(w, fmt.Sprintf("Failed to fetch balance data: %v", err), http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	// Check status code
	if res.StatusCode != http.StatusOK {
		log.Printf("Bad status code: %d", res.StatusCode)
		http.Error(w, fmt.Sprintf("Server returned status %d", res.StatusCode), http.StatusBadGateway)
		return
	}

	// Read response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		http.Error(w, fmt.Sprintf("Failed to read balance data: %v", err), http.StatusInternalServerError)
		return
	}

	// Set content type and write response
	w.Header().Set("Content-Type", "text/plain")
	w.Write(bodyBytes)
}
