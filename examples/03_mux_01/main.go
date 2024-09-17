package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello MUX"))
	})

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/", &bytes.Buffer{})
	mux.ServeHTTP(recorder, request)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Code: %v\n", recorder.Code)
	fmt.Printf("Body: %v\n", recorder.Body)
	fmt.Printf("Pattern: %v\n", request.Pattern)
}
