package handler

import (
	"fmt"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var adminkey string = os.Getenv("TMP_ADMIN_KEY")
	license := r.URL.Query()["license"][0]
	if license == adminkey {
		fmt.Fprintf(w, "Verified")
	} else {
		w.WriteHeader(403)
		return
	}
}
