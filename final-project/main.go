package main

import (
	"final-project/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Sayhello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
