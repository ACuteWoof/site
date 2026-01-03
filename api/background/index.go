package handler

import (
	"math/rand"
	"net/http"
	"time"
)

// Define your array of image URLs
var imageURLs = []string{
	"https://unsplash.com/photos/8-eG7YJ0gEY/download?w=1920",
	"https://unsplash.com/photos/9DedxUSwLg4/download?w=1920",
	"https://unsplash.com/photos/jSH-7IjJVsQ/download?w=1920",
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
	http.Redirect(w, r, randomURL, http.StatusTemporaryRedirect)
}

