

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

	postBody, _ := json.Marshal(map[string]string{
		"content": string("To: <@618114750827462660> From:" + r.RemoteAddr + "\n>>> " + r.URL.Query()["message"][0]),
	})

	fmt.Println(r.URL.Query()["message"][0])

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(websocket_uri, "application/json", responseBody)
	fmt.Println(resp)

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Fprintf(w, "Message sent to Woof!")
	}
}
