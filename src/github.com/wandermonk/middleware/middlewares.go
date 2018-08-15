package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s , %s\n", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
		log.Printf("Completed %s, %s\n", r.URL.Path, time.Since(start))
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index")
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the welcome page")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/welcome", loggingMiddleware(http.HandlerFunc(welcomeHandler)))
	http.ListenAndServe(":8080", nil)
}
