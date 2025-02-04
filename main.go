package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"encoding/json"
)

type JsonObj struct {
	Message string `json:"message"`
	Author string `json:"author"`
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

	msg := JsonObj{
		Message: "What did you expect, like really what?",
		Author: "Isaac Hayab",
	}

	fmt.Println("Another one gone...")
	json.NewEncoder(w).Encode(msg)
}
