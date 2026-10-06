package handler

import (
	"io"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request.", http.StatusBadRequest)
		return
	}

	if r.PostFormValue("access-code") != os.Getenv("ACCESS_CODE") {
		http.Error(w, "Access denied.", http.StatusForbidden)
		return
	}

	url := os.Getenv("CONTENT_BLOB_URL")
	token := os.Getenv("BLOB_READ_WRITE_TOKEN")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "Internal error.", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch content.", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch content.", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if _, err := io.Copy(w, resp.Body); err != nil {
		return
	}
}
