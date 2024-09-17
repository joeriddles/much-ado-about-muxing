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

	handle(mux, "GET /products")
	handle(mux, "GET /products/{id}")
	handle(mux, "GET /products/{id}/details")
	handle(mux, "GET /products/{id}/reviews")
	handle(mux, "GET /products/{id}/reviews/{review_id}")
	handle(mux, "GET /products/{category}/{id}")
	handle(mux, "GET /products/{category}/{id}/details")
	handle(mux, "GET /products/{category}/{id}/reviews")
	handle(mux, "GET /products/{category}/{id}/reviews/{review_id}")
	handle(mux, "GET /search")
	handle(mux, "GET /search/{query}")
	handle(mux, "GET /search/{query}/filter")
	handle(mux, "GET /search/{query}/filter/{filter_id}")
	handle(mux, "GET /search/{query}/filter/{filter_id}/sort")
	handle(mux, "GET /user/{username}")
	handle(mux, "GET /user/{username}/settings")
	handle(mux, "GET /user/{username}/settings/{section}")
	handle(mux, "GET /user/{username}/settings/{section}/edit")
	handle(mux, "GET /user/{username}/posts")
	handle(mux, "GET /user/{username}/posts/{post_id}")
	handle(mux, "GET /user/{username}/posts/{post_id}/comments")
	handle(mux, "GET /user/{username}/posts/{post_id}/comments/{comment_id}")
	handle(mux, "GET /admin")
	handle(mux, "GET /admin/{section}")
	handle(mux, "GET /admin/{section}/{id}")
	handle(mux, "GET /admin/{section}/{id}/edit")
	handle(mux, "GET /admin/{section}/{id}/delete")
	handle(mux, "GET /admin/{section}/{id}/view")
	handle(mux, "GET /admin/{section}/create")
	handle(mux, "GET /admin/{section}/list")

	request(mux, "GET", "/user/johndoe/posts/987/comments/654")
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
