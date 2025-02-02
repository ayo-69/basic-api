package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"encoding/json"
)

type Message struct {
	Content string `json:"content"`
}

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Connecting to server on PORT : 8080")

	//Setup PORT for heroku
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	msg := Message{
		Content: "What did you expect, like really what",
	}

	fmt.Println("User connected")
	json.NewEncoder(w).Encode(msg)
}
