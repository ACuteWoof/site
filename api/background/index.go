package handler

import (
	"math/rand"
	"net/http"
	"time"
)

// Define your array of image URLs
var imageURLs = []string{
	"https://unsplash.com/photos/e0wBK0xJXYQ/download",
	"https://unsplash.com/photos/9DedxUSwLg4/download",
	"https://unsplash.com/photos/jSH-7IjJVsQ/download",
	"https://unsplash.com/photos/J3KMRjNlCps/download",
	"https://unsplash.com/photos/8E31onOcVXg/download",
	"https://unsplash.com/photos/lSMpjzAOXgs/download",
	"https://unsplash.com/photos/emJByvG4l88/download",
	"https://unsplash.com/photos/HhMt_nUAIEc/download",
	"https://unsplash.com/photos/EQTZi6bae-Q/download",
	"https://unsplash.com/photos/XBvR0B6uzmA/download",
	"https://unsplash.com/photos/XnoOrTIjgjw/download",
	"https://unsplash.com/photos/YIfFVwDcgu8/download",
	"https://unsplash.com/photos/n-vnWQmmVoY/download",
	"https://unsplash.com/photos/Agwv1mKDUnc/download",
	"https://unsplash.com/photos/4BXWIQoS8Fo/download",
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

