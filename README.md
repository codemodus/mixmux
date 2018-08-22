# mixmux

    go get "github.com/codemodus/mixmux"

Package mixmux wraps HTTPRouter and HTTPTreeMux to provide consistent and
idiomatic APIs, along with route grouping.  Multiplexer-based parameter 
handling is bypassed.

## Usage

```go
type Mux
type Options
type Router
    func NewRouter(opts *Options) *Router
    func (m *Router) CORSMethods(path string, handlerWrappers ...func(http.Handler) http.Handler)
    func (m *Router) Connect(path string, h http.Handler)
    func (m *Router) Delete(path string, h http.Handler)
    func (m *Router) Get(path string, h http.Handler)
    func (m *Router) Group(path string) *Router
    func (m *Router) GroupMux(path string) Mux
    func (m *Router) Head(path string, h http.Handler)
    func (m *Router) Options(path string, h http.Handler)
    func (m *Router) Patch(path string, h http.Handler)
    func (m *Router) Post(path string, h http.Handler)
    func (m *Router) Put(path string, h http.Handler)
    func (m *Router) ServeHTTP(w http.ResponseWriter, r *http.Request)
    func (m *Router) Trace(path string, h http.Handler)
type TreeMux
    func NewTreeMux(opts *Options) *TreeMux
    func (m *TreeMux) CORSMethods(path string, handlerWrappers ...func(http.Handler) http.Handler)
    func (m *TreeMux) Connect(path string, h http.Handler)
    func (m *TreeMux) Delete(path string, h http.Handler)
    func (m *TreeMux) Get(path string, h http.Handler)
    func (m *TreeMux) Group(path string) *TreeMux
    func (m *TreeMux) GroupMux(path string) Mux
    func (m *TreeMux) Head(path string, h http.Handler)
    func (m *TreeMux) Options(path string, h http.Handler)
    func (m *TreeMux) Patch(path string, h http.Handler)
    func (m *TreeMux) Post(path string, h http.Handler)
    func (m *TreeMux) Put(path string, h http.Handler)
    func (m *TreeMux) ServeHTTP(w http.ResponseWriter, r *http.Request)
    func (m *TreeMux) Trace(path string, h http.Handler)
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
