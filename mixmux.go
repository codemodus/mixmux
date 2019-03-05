// Package mixmux wraps HTTPRouter and HTTPTreeMux to provide consistent and
// idiomatic APIs, along with route grouping.  Multiplexer-based parameter
// handling is bypassed.
package mixmux

import "net/http"

var (
	methods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodHead,
		http.MethodTrace,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodConnect,
	}
)

// Options holds available options for a new Router.
type Options struct {
	RedirectTrailingSlash  bool
	RedirectFixedPath      bool
	HandleMethodNotAllowed bool
	NotFound               http.Handler
	MethodNotAllowed       http.Handler
}

// Mux ...
type Mux interface {
	http.Handler

	GroupMux(path string) Mux
	Any(path string, h http.Handler)
	Options(path string, h http.Handler)
	Get(path string, h http.Handler)
	Post(path string, h http.Handler)
	Put(path string, h http.Handler)
	Patch(path string, h http.Handler)
	Delete(path string, h http.Handler)
	Head(path string, h http.Handler)
	Trace(path string, h http.Handler)
	Connect(path string, h http.Handler)
	CORSMethods(path string, handlerWrappers ...func(http.Handler) http.Handler)
}
