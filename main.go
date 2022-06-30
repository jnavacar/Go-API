package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Title   string `json:"Title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Messages []Message

func allMessages(w http.ResponseWriter, r *http.Request) {
	messages := Message{Title: "First Go endpoint", Content: "This is the first endpoint I create that returns some info", Author: "Jose"}
	json.NewEncoder(w).Encode(messages)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allMessages)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
