package handler

import (
	"math/rand"
	"net/http"
	"time"
)

// Define your array of image URLs
var imageURLs = []string{
	// misc
	"https://unsplash.com/photos/9DedxUSwLg4/download",
	"https://unsplash.com/photos/8E31onOcVXg/download",
	"https://unsplash.com/photos/Ame-KJjR5ZY/download",
	// "https://unsplash.com/photos/emJByvG4l88/download",
	"https://unsplash.com/photos/XBvR0B6uzmA/download",
	"https://unsplash.com/photos/n-vnWQmmVoY/download",
	"https://unsplash.com/photos/8XHXFfC6y0c/download",
	"https://unsplash.com/photos/MQ4cKv5d2G0/download",
	// "https://unsplash.com/photos/wWNz7wlv8zk/download",
	"https://unsplash.com/photos/qlJiJw7bcuU/download",
	// "https://unsplash.com/photos/CWzixmX3RY8/download",

	// flowers
	"https://unsplash.com/photos/YIfFVwDcgu8/download",
	"https://unsplash.com/photos/qi9jveT9X6A/download",
	"https://unsplash.com/photos/wPbZENgLO-4/download",
	"https://unsplash.com/photos/bg__Z7AJgEQ/download",
	"https://unsplash.com/photos/nnqLZGI_bpg/download",
	"https://unsplash.com/photos/jVIGXRapAO4/download",
	"https://unsplash.com/photos/TeNqlqOEMMI/download",
	"https://unsplash.com/photos/NQbZu_Y0ewE/download",
	"https://unsplash.com/photos/5TK1F5VfdIk/download",
	"https://unsplash.com/photos/w6PYXHtNls8/download",
	"https://unsplash.com/photos/1WlbWTo-a_A/download",
	"https://unsplash.com/photos/FwL9UMosJCg/download",
	"https://unsplash.com/photos/rLOAyD1nSns/download",
	"https://unsplash.com/photos/FjM4EHEzQe4/download",

	// landscapes
	"https://unsplash.com/photos/mdy0v3HnS1w/download",
	"https://unsplash.com/photos/hoI200JXjJY/download",

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

