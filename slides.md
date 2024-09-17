---
title: "Much Ado About Muxing"
marp: true
html: true
theme: gaia

---
<style>
section {
    --color-background: #ffdeff11 !important;
    --color-foreground: white;
}
section code {
  --color-background: #bbb; /* text color, weirdly */
  --color-dimmed: none;
}
section table {
  font-size: 0.65rem;

  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
section table thead th {
  --color-background: black;
}
.emphasis {
    font-weight: bolder;
    color: #11779C;
}
h1 {
    font-size: 1.8rem;
    font-weight: 600;
    margin-bottom: 1rem;
}
h2 {
    font-size: 2rem;
}
ul {
    margin-top: 0px;
    margin-left: 8px;
}
</style>

<!-- _class: lead -->
<h1>Much Ado About <span class="emphasis">Muxing</span></h1>

Joe Riddle

---

# Outline

- [What's muxing?](#whats-muxing)
- [Muxing in Go 1.21](#muxing-in-go-121)
- [Muxing in Go 1.21 — Features](#muxing-in-go-121-features)
- [Changes in Go 1.22 — New Features](#changes-in-go-122new-features)
- [Examples](#examples)
- [`ServeMux` vs. the other guys](#servemux-vs-the-other-guys)

---

# What's muxing? 

> `ServeMux` is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL. 

— https://pkg.go.dev/net/http#ServeMux

---

# Muxing in Go 1.21

- `ServeMux` is Go's built-in HTTP request multiplexer
- Sanitizes URL, headers, etc.
- Not very robust pattern matching

---

# Muxing in Go 1.21 — Features

* Match fixed paths, `/contact`
* Match wildcard, `/api/`
* Match hostname, `example.com/api/`
* What's missing?

<!--
Note the trailing slash in `/api/`. This means it matches `/api/*`.
-->

---

# Changes in Go 1.22 — New Features

* Match HTTP method, `GET /contact`
* Match segments, `GET /api/user/{id}`
* Match wildcard, `GET /api/{wc...}`
* Match end of URL, `GET /{$}`

---

<!-- _class: lead -->
# Examples

https://joeriddles.github.io/much-ado-about-muxing/

---

<!-- _class: lead -->
# `ServeMux` vs. the other guys

---

![](img/you%20vs%20the%20muxer.jpg)

---

<!--
We're focused on muxing/routing features.
Some of these packages do much more.
-->

| Features       | [ServeMux](https://pkg.go.dev/net/http#ServeMux) | [gorilla/mux](https://github.com/gorilla/mux) | [chi](https://github.com/go-chi/chi)                      | [echo](https://echo.labstack.com/) | [gin](https://github.com/gin-gonic/gin) |
| -------------- | ------------------------------------------------ | --------------------------------------------- | --------------------------------------------------------- | ---------------------------------- | --------------------------------------- |
| Fixed          | ✅                                                | ✅                                             | ✅                                                         | ✅                                  | ✅                                       |
| HTTP Method    | 1.22+ `"GET /"`                                  | `r.Methods("GET")`                            | `r.Get`                                                   | `e.GET`                            | `r.GET`                                 |
| Segment        | 1.22+ `"/{id}"`                                  | `"/{id}"`                                     | `"/{id}"`                                                 | `"/:id"`                           | `"/:id"`                                |
| Wildcard       | 1.22+ `"/static/"`                               | `r.PathPrefix("/static")`                     | `"/static/*"`                                             | `/static/*`                        | `/static/*`                             |
| Domain         | ✅                                                | `r.Host("example.com")`                       | [go-chi/hostrouter](https://github.com/go-chi/hostrouter) | ❌                                  | ❌                                       |
| Regex          | ❌                                                | `"/{id:[0-9]+}"`                              | `"/{id:[0-9]+}"`                                          | ❌                                  | ❌                                       |
| Subrouting     | ❌                                                | `r.Subrouter()`                               | `r.Route(...)`                                            | `e.Group(...)`                     | `r.Group(...)`                          |
| Named URLs     | ❌                                                | `r.Name("index")`                             | ❌                                                         | `route.Name="index"`               | ❌                                       |
| Other features |                                                  | Schemes, headers, queries                     | Compatible with `net/http`                                | Framework                          | Framework                               |
| Precedence     | Specific                                         | In order                                      | Static, param, * ?                                        | Static, param, *                   | N/A, unambigous                         |
| Data structure | Radix tree                                       | List                                          | Radix tree                                                | Radix tree                         | Radix tree                              |
---

<!-- _class: lead -->
# Questions 🎤
