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
