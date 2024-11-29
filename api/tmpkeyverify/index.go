package handler

import (
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var adminkey string = os.Getenv("TMP_ADMIN_KEY")
	license := r.URL.Query()["license"][0]
	if license == adminkey {
		w.WriteHeader(200);
		return
	} else {
		w.WriteHeader(403)
		return
	}
}
