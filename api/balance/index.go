package handler

import (
	"net/http"
	"io"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("https://70.34.201.17/log.csv")
	defer res.Body.Close()
	bodyBytes, _ := io.ReadAll(res.Body)

	w.Header().Set("Content-Type", "text/plain")
	w.Write(bodyBytes)
}
