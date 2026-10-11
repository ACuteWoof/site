package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_code")
	if err != nil || cookie.Value == "" || cookie.Value != os.Getenv("ACCESS_CODE") {
		http.Error(w, "Access denied. Valid cookie required.", http.StatusForbidden)
		return
	}

	requestedFile := strings.TrimPrefix(r.URL.Path, "/api/file/")
	if requestedFile == "" || requestedFile == "/" {
		http.Error(w, "Filename missing.", http.StatusBadRequest)
		return
	}

	for _, segment := range strings.Split(requestedFile, "/") {
		if segment == "." || segment == ".." || segment == "" {
			http.Error(w, "Invalid filename.", http.StatusBadRequest)
			return
		}
	}

	blobBaseURL := strings.TrimSuffix(
		os.Getenv("CONTENT_BLOB_BASE_URL"), "/",
	)
	token := os.Getenv("BLOB_READ_WRITE_TOKEN")

	if blobBaseURL == "" || token == "" {
		http.Error(w, "Blob configuration missing.", http.StatusInternalServerError)
		return
	}

	targetURL := blobBaseURL + "/" + requestedFile

	req, err := http.NewRequestWithContext(
		r.Context(),
		http.MethodGet,
		targetURL,
		nil,
	)
	if err != nil {
		http.Error(w, "Failed to create request.", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch blob.", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(
			w,
			fmt.Sprintf("Blob returned: %s", resp.Status),
			http.StatusBadGateway,
		)
		return
	}

	// Forward relevant response headers.
	for _, name := range []string{
		"Content-Type",
		"Content-Length",
		"Content-Disposition",
		"Content-Encoding",
		"Last-Modified",
		"ETag",
		"Accept-Ranges",
		"Content-Range",
	} {
		if value := resp.Header.Get(name); value != "" {
			w.Header().Set(name, value)
		}
	}

	w.Header().Set("Cache-Control", "private, no-store")

	_, _ = io.Copy(w, resp.Body)
}
