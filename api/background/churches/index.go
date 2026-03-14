package handler

import (
	"math/rand"
	"net/http"
	"time"
)

// Define your array of image URLs
var imageURLs = []string{
	"https://unsplash.com/photos/UEesKwR8ccY/download",
	"https://unsplash.com/photos/zNtmXuF9hLg/download",
	"https://unsplash.com/photos/ikD_o6lSNNs/download",
	"https://unsplash.com/photos/mgLX9vQhxc8/download",
	"https://unsplash.com/photos/0gkNn5i7MOg/download",
	"https://unsplash.com/photos/_FTGcg8_OBs/download",
}

func init() {
	// Seed the random number generator once when the function is initialized
	rand.Seed(time.Now().UnixNano())
}

// Handler is the main serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	// Select a random index
	randomIndex := rand.Intn(len(imageURLs))
	randomURL := imageURLs[randomIndex]

	// Redirect the client to the randomly selected image URL
	// Use http.StatusTemporaryRedirect (307) or http.StatusFound (302) for temporary redirects,
        // or http.StatusPermanentRedirect (308) for permanent redirects if desired
	w.Header().Set("Cache-Control", "no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache") // For older HTTP/1.0 compatibility
	w.Header().Set("Expires", "0")       // Also for older compatibility
	http.Redirect(w, r, randomURL, http.StatusTemporaryRedirect)
}


