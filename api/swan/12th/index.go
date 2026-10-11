package handler

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var authorized bool
	cookie, err := r.Cookie("access_code")
	if err == nil && cookie.Value == os.Getenv("ACCESS_CODE") {
		authorized = true
	}

	if !authorized {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/me/swan", http.StatusTemporaryRedirect)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Redirect(w, r, "/me/swan", http.StatusTemporaryRedirect)
			return
		}

		inputCode := r.PostFormValue("access-code")
		if inputCode != os.Getenv("ACCESS_CODE") {
			http.Error(w, "Access denied.", http.StatusForbidden)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_code",
			Value:    inputCode,
			Path:     "/",
			Expires:  time.Now().Add(30 * 24 * time.Hour),
			HttpOnly: true,                               
			Secure:   true,                              
			SameSite: http.SameSiteLaxMode,
		})
	}

	blobBaseURL := strings.TrimSuffix(
		os.Getenv("CONTENT_BLOB_BASE_URL"), "/",
	)
	url := blobBaseURL + "/anniversary.html"
	token := os.Getenv("BLOB_READ_WRITE_TOKEN")

	if url == "" || token == "" {
		http.Error(w, "Blob configuration missing.", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
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
		http.Error(w, "Blob returned: "+resp.Status, http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	if _, err := io.Copy(w, resp.Body); err != nil {
		return
	}
}

