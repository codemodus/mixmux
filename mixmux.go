// Package mixmux wraps HTTPRouter and HTTPTreeMux to provide consistent and
// idiomatic APIs, along with route grouping.  Multiplexer-based parameter
// handling is bypassed.
package mixmux

import (
	"net/http"
	"strings"

	"github.com/dimfeld/httptreemux"
	"github.com/julienschmidt/httprouter"
)

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
		opts = &Options{}
	}

	r.hr = &httprouter.Router{
		RedirectTrailingSlash:  opts.RedirectTrailingSlash,
		RedirectFixedPath:      opts.RedirectFixedPath,
		HandleMethodNotAllowed: opts.HandleMethodNotAllowed,
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
	r.hr.Handler(http.MethodOptions, r.path+path, h)
}

// Get takes a path and http.Handler and adds them to the mux.
func (r *Router) Get(path string, h http.Handler) {
	r.hr.Handler(http.MethodGet, r.path+path, h)
}

// Post takes a path and http.Handler and adds them to the mux.
func (r *Router) Post(path string, h http.Handler) {
	r.hr.Handler(http.MethodPost, r.path+path, h)
}

// Put takes a path and http.Handler and adds them to the mux.
func (r *Router) Put(path string, h http.Handler) {
	r.hr.Handler(http.MethodPut, r.path+path, h)
}

// Patch takes a path and http.Handler and adds them to the mux.
func (r *Router) Patch(path string, h http.Handler) {
	r.hr.Handler(http.MethodPatch, r.path+path, h)
}

// Delete takes a path and http.Handler and adds them to the mux.
func (r *Router) Delete(path string, h http.Handler) {
	r.hr.Handler(http.MethodDelete, r.path+path, h)
}

// Head takes a path and http.Handler and adds them to the mux.
func (r *Router) Head(path string, h http.Handler) {
	r.hr.Handler(http.MethodHead, r.path+path, h)
}

// Trace takes a path and http.Handler and adds them to the mux.
func (r *Router) Trace(path string, h http.Handler) {
	r.hr.Handler(http.MethodTrace, r.path+path, h)
}

// Connect takes a path and http.Handler and adds them to the mux.
func (r *Router) Connect(path string, h http.Handler) {
	r.hr.Handler(http.MethodConnect, r.path+path, h)
}

// OptionsAuto ... TODO:
func (r *Router) OptionsAuto(path string, handlerWrapper func(http.Handler) http.Handler) {
	h, _, s := r.hr.Lookup(http.MethodOptions, path)
	if s {
		h, _, _ = r.hr.Lookup(http.MethodOptions, path+"/")
	}
	if h != nil {
		return
	}

	ms := []string{http.MethodOptions}

	for _, v := range methods {
		if v == http.MethodOptions {
			continue
		}

		h, _, s = r.hr.Lookup(v, path)
		if s {
			h, _, _ = r.hr.Lookup(v, path+"/")
		}
		if h == nil {
			continue
		}

		ms = append(ms, v)
	}

	opts := strings.Join(ms, ", ")

	var fn http.Handler
	fn = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", opts)
	})

	if handlerWrapper != nil {
		fn = handlerWrapper(fn)
	}

	r.hr.Handler(http.MethodOptions, path, fn)
}

// ServeHTTP satisfies the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.hr.ServeHTTP(w, req)
}

// TreeMux wraps HTTPTreeMux.
type TreeMux struct {
	tm   *httptreemux.TreeMux
	path string
}

// NewTreeMux returns a wrapped HTTPTreeMux.
func NewTreeMux(opts *Options) *TreeMux {
	t := &TreeMux{
		tm:   httptreemux.New(),
		path: "",
	}

	if opts == nil {
		opts = &Options{}
	}

	if opts.NotFound != nil {
		t.tm.NotFoundHandler = opts.NotFound.ServeHTTP
	}

	if opts.MethodNotAllowed != nil {
		t.tm.MethodNotAllowedHandler = func(w http.ResponseWriter, r *http.Request, m map[string]httptreemux.HandlerFunc) {
			opts.MethodNotAllowed.ServeHTTP(w, r)
		}
	}

	t.tm.RedirectTrailingSlash = opts.RedirectTrailingSlash
	t.tm.RedirectCleanPath = opts.RedirectFixedPath
	t.tm.RedirectTrailingSlash = true

	return t
}

// Group takes a path and returns a new TreeMux wrapping the original TreeMux.
func (tm *TreeMux) Group(path string) *TreeMux {
	return &TreeMux{tm.tm, tm.path + path}
}

// Options takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Options(path string, h http.Handler) {
	tm.tm.Handle(http.MethodOptions, tm.path+path, treeMuxWrapper(h))
}

// Get takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Get(path string, h http.Handler) {
	tm.tm.Handle(http.MethodGet, tm.path+path, treeMuxWrapper(h))
}

// Post takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Post(path string, h http.Handler) {
	tm.tm.Handle(http.MethodPost, tm.path+path, treeMuxWrapper(h))
}

// Put takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Put(path string, h http.Handler) {
	tm.tm.Handle(http.MethodPut, tm.path+path, treeMuxWrapper(h))
}

// Patch takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Patch(path string, h http.Handler) {
	tm.tm.Handle(http.MethodPatch, tm.path+path, treeMuxWrapper(h))
}

// Delete takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Delete(path string, h http.Handler) {
	tm.tm.Handle(http.MethodDelete, tm.path+path, treeMuxWrapper(h))
}

// Head takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Head(path string, h http.Handler) {
	tm.tm.Handle(http.MethodHead, tm.path+path, treeMuxWrapper(h))
}

// Trace takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Trace(path string, h http.Handler) {
	tm.tm.Handle(http.MethodTrace, tm.path+path, treeMuxWrapper(h))
}

// Connect takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Connect(path string, h http.Handler) {
	tm.tm.Handle(http.MethodConnect, tm.path+path, treeMuxWrapper(h))
}

// ServeHTTP satisfies the http.Handler interface.
func (tm *TreeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm.tm.ServeHTTP(w, r)
}

func treeMuxWrapper(next http.Handler) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		next.ServeHTTP(w, r)
	}
}
