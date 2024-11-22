package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Message sent to Woof!")

	var websocket_uri string = os.Getenv("WEBSOCKET_URI")

	postBody, _ := json.Marshal(map[string]string{
		"content": "@acutewoof, " + r.RemoteAddr + " says: " + r.URL.Query()["message"][0],
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(websocket_uri, "application/json", responseBody)
	fmt.Println(resp)

	if err != nil {
		log.Fatalln(err)
	}

}
