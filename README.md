# mixmux

    go get "github.com/codemodus/mixmux"

Package mixmux wraps HTTPRouter and HTTPTreeMux to provide consistent and
idiomatic APIs, along with route grouping.  Multiplexer-based parameter 
handling is bypassed.

## Usage

```go
type Router
    func NewRouter() *Router
    func (r *Router) Delete(path string, h http.Handler)
    func (r *Router) Get(path string, h http.Handler)
    func (r *Router) Group(path string) *Router
    func (r *Router) Head(path string, h http.Handler)
    func (r *Router) Options(path string, h http.Handler)
    func (r *Router) Patch(path string, h http.Handler)
    func (r *Router) Post(path string, h http.Handler)
    func (r *Router) Put(path string, h http.Handler)
type TreeMux
    func NewTreeMux() *TreeMux
    func (tm *TreeMux) Delete(path string, h http.Handler)
    func (tm *TreeMux) Get(path string, h http.Handler)
    func (tm *TreeMux) Group(path string) *TreeMux
    func (tm *TreeMux) Head(path string, h http.Handler)
    func (tm *TreeMux) Options(path string, h http.Handler)
    func (tm *TreeMux) Patch(path string, h http.Handler)
    func (tm *TreeMux) Post(path string, h http.Handler)
    func (tm *TreeMux) Put(path string, h http.Handler)
```

### Setup

```go
import (
    "net/http"
    "net/http/httptest"

    "github.com/codemodus/mixmux"
)

func main() {
    handler := http.HandlerFunc(methodHandler)

    mux := mixmux.NewRouter()
    mux.Get("/get", handler)
    mux.Post("/post", handler)

    muxGroup := mux.Group("/grouped")
    muxGroup.Get("/get0", handler) // path = "/grouped/get0"
    muxGroup.Get("/get1", handler) // path = "/grouped/get1"

    // ...
}
```

## More Info

### Why is multiplexer-based parameter handling bypassed?

Multiplexer-based parameter handling is bypassed in favor of a single solution 
that is multiplexer-agnostic.  Please review codemodus/parth for a simple and 
effective package covering this need.

## Documentation

View the [GoDoc](http://godoc.org/github.com/codemodus/mixmux)

## Benchmarks

These results demonstrate that mixmux does not increase resource usage.  The 
digit suffix indicates how many named parameters are used.  http.ServeMux is 
included for reference.

    benchmark                       iter      time/iter   bytes alloc         allocs
    ---------                       ----      ---------   -----------         ------
    BenchmarkHTTPServeMux0        200000    55.55 μs/op     3547 B/op   54 allocs/op
    BenchmarkHTTPTreeMux2         200000    55.57 μs/op     3879 B/op   56 allocs/op
    BenchmarkHTTPRouter2          200000    54.16 μs/op     3600 B/op   55 allocs/op
    BenchmarkMixmuxTreeMux2       200000    55.26 μs/op     3869 B/op   56 allocs/op
    BenchmarkMixmuxRouter2        200000    54.75 μs/op     3593 B/op   55 allocs/op
    BenchmarkMixmuxRouterGroup2   200000    54.48 μs/op     3592 B/op   55 allocs/op
