const EXAMPLES = [
    {
        "title": "Hello World",
        "code": "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"hello world\")\n}\n"
    },
    {
        "title": "Errors",
        "code": "package main\n\nimport \"log\"\n\nfunc main() {\n    log.Fatalf(\"dead\")\n}\n"
    },
    {
        "title": "Mux 01",
        "code": "package main\n\nimport (\n    \"bytes\"\n    \"fmt\"\n    \"log\"\n    \"net/http\"\n    \"net/http/httptest\"\n)\n\nfunc main() {\n    mux := http.NewServeMux()\n    mux.HandleFunc(\"GET /\", func(w http.ResponseWriter, r *http.Request) {\n        w.Write([]byte(\"Hello MUX\"))\n    })\n\n    recorder := httptest.NewRecorder()\n    request, err := http.NewRequest(\"GET\", \"/\", &bytes.Buffer{})\n    mux.ServeHTTP(recorder, request)\n\n    if err != nil {\n        log.Fatal(err)\n    }\n    fmt.Printf(\"Code: %v\\n\", recorder.Code)\n    fmt.Printf(\"Body: %v\\n\", recorder.Body)\n    fmt.Printf(\"Pattern: %v\\n\", request.Pattern)\n}\n"
    }
]