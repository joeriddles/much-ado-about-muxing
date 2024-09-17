package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
)

func main() {
	mux := http.NewServeMux()

	handle(mux, "GET /{$}")
	handle(mux, "GET /{wc...}")

	request(mux, "GET", "/hello")
}

// Ignore everything below this...

// Register the default handler at the [route].
func handle(mux *http.ServeMux, route string) {
	mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, route)
	})
}

// Request and print metadata to stdout.
func request(mux *http.ServeMux, method, url string) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, &bytes.Buffer{})
	fmt.Printf("Path : %v %v\n", request.Method, request.URL.Path)
	mux.ServeHTTP(recorder, request)

	if err != nil {
		log.Fatal(err)
	}
}

var pattern regexp.Regexp = *regexp.MustCompile(`{(\w+)}`)

func handler(w http.ResponseWriter, r *http.Request, route string) {
	w.WriteHeader(200)
	fmt.Printf("Route: %v\n", route)

	for _, match := range pattern.FindAllStringSubmatch(route, -1) {
		segment := match[1]
		fmt.Printf("%v: %v\n", segment, r.PathValue(segment))
	}

	fmt.Println()
}
