package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
)

// GET  /
// GET  /contact
// GET  /api/widgets
// POST /api/widgets
// POST /api/widgets/:slug
// POST /api/widgets/:slug/parts
// POST /api/widgets/:slug/parts/:id/update
// POST /api/widgets/:slug/parts/:id/delete
// GET  /:slug
// GET  /:slug/admin
// POST /:slug/image

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "GET /{$}")
	})

	mux.HandleFunc("GET /contact", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "GET /contact")
	})

	mux.HandleFunc("GET /api/widgets", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "GET /api/widgets")
	})

	mux.HandleFunc("POST /api/widgets", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "POST /api/widgets")
	})

	mux.HandleFunc("POST /api/widgets/{slug}", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "POST /api/widgets/{slug}")
	})

	mux.HandleFunc("POST /api/widgets/{slug}/parts", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "POST /api/widgets/{slug}/parts")
	})

	mux.HandleFunc("POST /api/widgets/{slug}/parts/{id}/update", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "POST /api/widgets/{slug}/parts/{id}/update")
	})

	mux.HandleFunc("POST /api/widgets/{slug}/parts/{id}/delete", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "POST /api/widgets/{slug}/parts/{id}/delete")
	})

	mux.HandleFunc("GET /{slug}", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "GET /{slug}")
	})

	mux.HandleFunc("GET /{slug}/admin", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "GET /{slug}/admin")
	})
	mux.HandleFunc("POST /{slug}/image", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "POST /{slug}/image")
	})

	request(mux, "GET", "/")
	request(mux, "GET", "/contact")
	request(mux, "GET", "/api/widgets")
	request(mux, "POST", "/api/widgets")
	request(mux, "POST", "/api/widgets/slug")
	request(mux, "POST", "/api/widgets/slug/parts")
	request(mux, "POST", "/api/widgets/slug/parts/1/update")
	request(mux, "POST", "/api/widgets/slug/parts/1/delete")
	request(mux, "GET", "/slug")
	request(mux, "GET", "/slug/admin")
	request(mux, "POST", "/slug/image")

	request(mux, "POST", "/")
}

var pattern regexp.Regexp = *regexp.MustCompile(`{(\w+)}`)

func handler(w http.ResponseWriter, r *http.Request, route string) {
	w.WriteHeader(200)
	fmt.Printf("Route: %v\n", route)

	for _, match := range pattern.FindAllStringSubmatch(route, -1) {
		segment := match[1]
		// fmt.Printf("%v\n", strings.Join(match, ","))
		fmt.Printf("%v: %v\n", segment, r.PathValue(segment))
	}

	fmt.Println()
}

func request(mux *http.ServeMux, method, url string) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, &bytes.Buffer{})
	fmt.Printf("Path : %v %v\n", request.Method, request.URL.Path)
	mux.ServeHTTP(recorder, request)

	if err != nil {
		log.Fatal(err)
	}
}
