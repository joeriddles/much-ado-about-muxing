package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
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
		fmt.Printf("Path : %v %v\n", r.Method, r.URL.Path)

		switch r.Method {
		case http.MethodGet:
			segment, found := strings.CutPrefix(r.URL.Path, "/")
			if !found || segment == "" {
				handler(w, r, "GET /")
				return
			} else {
				fmt.Printf("Segmt: %v\n", segment)
			}

			segments := strings.Split(segment, "/")
			if len(segments) == 1 {
				handler(w, r, "GET /:slug")
				return
			}

			switch segments[1] {
			case "admin":
				handler(w, r, "GET /:slug/admin")
				return
			case "image":
				handler(w, r, "GET /:slug/image")
				return
			}

			http.Error(w, "Not found", http.StatusNotFound)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Path : %v %v\n", r.Method, r.URL.Path)

		switch r.Method {
		case http.MethodGet:
			handler(w, r, "GET /contact")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/widgets", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Path : %v %v\n", r.Method, r.URL.Path)

		switch r.Method {
		case http.MethodGet:
			handler(w, r, "GET /api/widgets")
		case http.MethodPost:
			handler(w, r, "POST /api/widgets")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/widgets/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Path : %v %v\n", r.Method, r.URL.Path)

		switch r.Method {
		case http.MethodPost:
			segment, _ := strings.CutPrefix(r.URL.Path, "/api/widgets/")
			fmt.Printf("Segmt: %v\n", segment)

			handler(w, r, "POST /api/widgets/")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
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
}

func handler(w http.ResponseWriter, _ *http.Request, route string) {
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Route: %v", route)))
}

func request(mux *http.ServeMux, method, url string) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, &bytes.Buffer{})
	mux.ServeHTTP(recorder, request)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", recorder.Body)
	fmt.Println()
}
