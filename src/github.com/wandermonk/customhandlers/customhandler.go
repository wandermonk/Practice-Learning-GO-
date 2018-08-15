package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	th := &TextHandler{"this is the responseString"}
	mux.Handle("/", th)
	http.ListenAndServe(":8080", mux)
}

type TextHandler struct {
	responseText string
}

func (th *TextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, th.responseText)
}
