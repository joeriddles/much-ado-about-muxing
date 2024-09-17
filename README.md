# Much Ado About Muxing

`net/http` had a glow up... and its name is `ServeMux`.

Is it time to kick Chi, Echo, and Gin to the curb? It just might be.

Go 1.22 came out recently and added some much needed improvements to the standard library's HTTP request multiplexer (mux). We'll explore these new additions and help you figure out if it's time to return to sweet, stdlib goodness for your HTTP 'plexing needs.

_A presentation for the Spokane Go User Group: https://www.meetup.com/spokane-go-users-group/events/303409967/._

### Resources
- [`ServeMux` docs](https://pkg.go.dev/net/http#ServeMux)
- [net/http: enhanced ServeMux routing #61410](https://github.com/golang/go/issues/61410), released in 1.22
    - [net/http: add methods and path variables to ServeMux patterns #60227](https://github.com/golang/go/discussions/60227)
- [net/http: expose matched pattern in Request #66405](https://github.com/golang/go/issues/66405), released in 1.23.1
- [Notes on running Go in the browser with WebAssembly](https://eli.thegreenplace.net/2024/notes-on-running-go-in-the-browser-with-webassembly/)
