package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/internal/shortener"
)

func main() {
	http.HandleFunc("/shorten", shortener.ShortenHandler)
	http.HandleFunc("/", shortener.RedirectHandler)

	fmt.Println("Server running on: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
