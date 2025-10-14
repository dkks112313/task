package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Metadata struct {
	Path string `json:"page"`
}

type Event struct {
	UserID   uint     `json:"user_id"`
	Action   string   `json:"action"`
	Metadata Metadata `json:"metadata"`
}

var savePlace []Event

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	http.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
				var event Event
				decoder := json.NewDecoder(r.Body)
				err := decoder.Decode(&event)
				if err != nil {
					fmt.Errorf("Error, when decode json")
				}

				println(event.UserID)
				println(event.Action)
				println(event.Metadata.Path)
			}
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
