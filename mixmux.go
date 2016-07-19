// Package mixmux wraps HTTPRouter and HTTPTreeMux to provide consistent and
// idiomatic APIs, along with route grouping.  Multiplexer-based parameter
// handling is bypassed.
package mixmux

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/julienschmidt/httprouter"
)

type Options struct {
	RedirectTrailingSlash  bool
	RedirectFixedPath      bool
	HandleMethodNotAllowed bool
	HandleOptions          bool
	NotFound               http.Handler
	MethodNotAllowed       http.Handler
}

// Router wraps HTTPRouter.
type Router struct {
	hr   *httprouter.Router
	path string
}

// NewRouter returns a wrapped HTTPRouter.
func NewRouter(opts *Options) *Router {
	r := &Router{
		path: "",
	}

	if opts == nil {
		r.hr = httprouter.New()

		return r
	}

	r.hr = &httprouter.Router{
		RedirectTrailingSlash:  opts.RedirectTrailingSlash,
		RedirectFixedPath:      opts.RedirectFixedPath,
		HandleMethodNotAllowed: opts.HandleMethodNotAllowed,
		HandleOPTIONS:          opts.HandleOptions,
		NotFound:               opts.NotFound,
		MethodNotAllowed:       opts.MethodNotAllowed,
	}

	return r
}

// Group takes a path and returns a new Router wrapping the original Router.
func (r *Router) Group(path string) *Router {
	return &Router{r.hr, r.path + path}
}

// Options takes a path and http.Handler and adds them to the mux.
func (r *Router) Options(path string, h http.Handler) {
	r.hr.Handler("OPTIONS", r.path+path, h)
}

// Get takes a path and http.Handler and adds them to the mux.
func (r *Router) Get(path string, h http.Handler) {
	r.hr.Handler("GET", r.path+path, h)
}

// Post takes a path and http.Handler and adds them to the mux.
func (r *Router) Post(path string, h http.Handler) {
	r.hr.Handler("POST", r.path+path, h)
}

// Put takes a path and http.Handler and adds them to the mux.
func (r *Router) Put(path string, h http.Handler) {
	r.hr.Handler("PUT", r.path+path, h)
}

// Patch takes a path and http.Handler and adds them to the mux.
func (r *Router) Patch(path string, h http.Handler) {
	r.hr.Handler("PATCH", r.path+path, h)
}

// Delete takes a path and http.Handler and adds them to the mux.
func (r *Router) Delete(path string, h http.Handler) {
	r.hr.Handler("DELETE", r.path+path, h)
}

// Head takes a path and http.Handler and adds them to the mux.
func (r *Router) Head(path string, h http.Handler) {
	r.hr.Handler("HEAD", r.path+path, h)
}

func (r *Router) Handle(method string, path string, h http.Handler) {
	r.hr.Handler(method, path, h)
}

func (mr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mr.hr.ServeHTTP(w, r)
}

// TreeMux wraps HTTPTreeMux.
type TreeMux struct {
	tm   *httptreemux.TreeMux
	path string
}

// NewTreeMux returns a wrapped HTTPTreeMux.
func NewTreeMux() *TreeMux {
	return &TreeMux{
		tm:   httptreemux.New(),
		path: "",
	}
}

// Group takes a path and returns a new TreeMux wrapping the original TreeMux.
func (tm *TreeMux) Group(path string) *TreeMux {
	return &TreeMux{tm.tm, tm.path + path}
}

// Options takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Options(path string, h http.Handler) {
	tm.tm.Handle("OPTIONS", tm.path+path, treeMuxWrapper(h))
}

// Get takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Get(path string, h http.Handler) {
	tm.tm.Handle("GET", tm.path+path, treeMuxWrapper(h))
}

// Post takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Post(path string, h http.Handler) {
	tm.tm.Handle("POST", tm.path+path, treeMuxWrapper(h))
}

// Put takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Put(path string, h http.Handler) {
	tm.tm.Handle("PUT", tm.path+path, treeMuxWrapper(h))
}

// Patch takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Patch(path string, h http.Handler) {
	tm.tm.Handle("PATCH", tm.path+path, treeMuxWrapper(h))
}

// Delete takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Delete(path string, h http.Handler) {
	tm.tm.Handle("DELETE", tm.path+path, treeMuxWrapper(h))
}

// Head takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Head(path string, h http.Handler) {
	tm.tm.Handle("HEAD", tm.path+path, treeMuxWrapper(h))
}

func (tm *TreeMux) Handle(method string, path string, h http.Handler) {
	tm.tm.Handle(method, path, treeMuxWrapper(h))
}

func (tm *TreeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm.tm.ServeHTTP(w, r)
}

func treeMuxWrapper(next http.Handler) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		next.ServeHTTP(w, r)
	}
}
