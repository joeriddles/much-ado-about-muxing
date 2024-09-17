package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			handler(w, r, "GET /")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	request(mux, "GET", "/")
	request(mux, "GET", "/contact")
	request(mux, "GET", "/hello-world")
	request(mux, "GET", "/hello-world/admin")
	request(mux, "POST", "/")
}

func handler(w http.ResponseWriter, _ *http.Request, route string) {
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Route: %v", route)))
}

func request(mux *http.ServeMux, method, url string) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, &bytes.Buffer{})
	fmt.Printf("Path : %v %v\n", request.Method, request.URL.Path)

	mux.ServeHTTP(recorder, request)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", recorder.Body)
	fmt.Println()
}
