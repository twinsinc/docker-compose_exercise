package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	imgURL := os.Getenv("IMAGES_URL")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    	root(w, r, imgURL)
	})
	log.Fatalf("error: %s", http.ListenAndServe(":8080", nil))
}
func root(w http.ResponseWriter, r *http.Request, imgURL string) {
	log.Printf("new request: %s:%s", r.Method, r.URL.String())

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<img src='%s/the-cloud.jpg' alt='some meme'>", imgURL) // this is very ugly, don't do this yourself
}
