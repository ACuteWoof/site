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
		fmt.Fprintf(w, `{"license": "verified"}`)
		return
	} else {
		w.WriteHeader(403)
		fmt.Fprintf(w, `{"license": "bad"}`)
		return
	}
}
