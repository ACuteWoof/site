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
	var websocket_uri string = os.Getenv("WEBSOCKET_URI")

	message := r.FormValue("message")

	postBody, _ := json.Marshal(map[string]string{
		"content": string("To: <@618114750827462660> From:" + r.RemoteAddr + "\n---\n\n" + message),
	})

	fmt.Println(r.URL.Query()["message"][0])
	fmt.Println(message)

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(websocket_uri, "application/json", responseBody)
	fmt.Println(resp)

	if err != nil {
		log.Fatalln(err)

		postBody, _ := json.Marshal(map[string]string{
			"content": string("To: <@618114750827462660> From:" + r.RemoteAddr + "\n---\n\n" + "CHECK VERCEL LOGS"),
		})

		responseBody := bytes.NewBuffer(postBody)

		resp, err = http.Post(websocket_uri, "application/json", responseBody)
		fmt.Println(resp)

	} else {
		fmt.Fprintf(w, "Message sent!")
	}
}
