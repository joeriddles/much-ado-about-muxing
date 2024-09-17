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

	handle(mux, "GET /a/b/c/d")
	handle(mux, "POST /x/y/z")
	handle(mux, "GET /abc/{id}")
	handle(mux, "PUT /abc/{id}/edit")
	handle(mux, "GET /a/{b}/{c}/d/{e}")
	handle(mux, "DELETE /alpha/{beta}/gamma")
	handle(mux, "GET /a/{b}/c/{d}/e/{f}")
	handle(mux, "POST /123/{xyz}/456")
	handle(mux, "GET /a/{foo}/{bar}/{baz}")
	handle(mux, "PUT /α/{β}/γ")
	handle(mux, "GET /foo/bar/{id}/baz")
	handle(mux, "DELETE /a/{b}/c/d/{e}")
	handle(mux, "GET /x/y/{z}/w")
	handle(mux, "POST /m/n/o/p")
	handle(mux, "GET /hello/{world}")
	handle(mux, "PATCH /a/{b}/c/{d}/e/{f}/g/{h}")
	handle(mux, "GET /data/{key}/value")
	handle(mux, "POST /test/{param1}/sub/{param2}")
	handle(mux, "GET /q/r/s/{t}/u")
	handle(mux, "PUT /item/{id}/update")
	handle(mux, "GET /日本/{国家}/都市")
	handle(mux, "POST /nueva/{ruta}/prueba")
	handle(mux, "GET /api/{v}/data")
	handle(mux, "DELETE /path/{one}/here/{two}")
	handle(mux, "GET /file/{type}/{name}")
	handle(mux, "PUT /root/{dir}/file/{name}")
	handle(mux, "GET /lorem/{ipsum}/dolor")
	handle(mux, "POST /short/{a}/long/{b}/path")
	handle(mux, "GET /region/{a}/{b}/{c}")
	handle(mux, "PATCH /edit/{section}/{item}")
	handle(mux, "GET /some/{path}/example")
	handle(mux, "POST /abc/{x}/{y}/def")
	handle(mux, "DELETE /delete/{item}/{info}")
	handle(mux, "GET /contact/{person}/details")
	handle(mux, "PUT /update/{entry}/value")
	handle(mux, "GET /upload/{file}/status")
	handle(mux, "POST /foo/bar/{id}/baz")
	handle(mux, "GET /level/{one}/{two}/{three}")
	handle(mux, "PATCH /settings/{module}/update")
	handle(mux, "DELETE /remove/{entry}/from/{list}")

	// TODO
	// request(mux, "GET", "...")
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
	fmt.Printf("Status: %v", recorder.Code)
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
