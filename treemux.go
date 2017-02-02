package mixmux

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
)

// TreeMux wraps HTTPTreeMux.
type TreeMux struct {
	t    *httptreemux.TreeMux
	path string
}

// NewTreeMux returns a wrapped HTTPTreeMux.
func NewTreeMux(opts *Options) *TreeMux {
	t := &TreeMux{
		t:    httptreemux.New(),
		path: "",
	}

	if opts == nil {
		opts = &Options{}
	}

	if opts.NotFound != nil {
		t.t.NotFoundHandler = opts.NotFound.ServeHTTP
	}

	if opts.MethodNotAllowed != nil {
		t.t.MethodNotAllowedHandler = func(w http.ResponseWriter, r *http.Request, m map[string]httptreemux.HandlerFunc) {
			opts.MethodNotAllowed.ServeHTTP(w, r)
		}
	}

	t.t.RedirectTrailingSlash = opts.RedirectTrailingSlash
	t.t.RedirectCleanPath = opts.RedirectFixedPath
	t.t.RedirectTrailingSlash = true

	return t
}

// Group takes a path and returns a new TreeMux wrapping the original TreeMux.
func (m *TreeMux) Group(path string) *TreeMux {
	return &TreeMux{m.t, m.path + path}
}

// Options takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Options(path string, h http.Handler) {
	m.t.Handle(http.MethodOptions, m.path+path, treeMuxWrapper(h))
}

// Get takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Get(path string, h http.Handler) {
	m.t.Handle(http.MethodGet, m.path+path, treeMuxWrapper(h))
}

// Post takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Post(path string, h http.Handler) {
	m.t.Handle(http.MethodPost, m.path+path, treeMuxWrapper(h))
}

// Put takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Put(path string, h http.Handler) {
	m.t.Handle(http.MethodPut, m.path+path, treeMuxWrapper(h))
}

// Patch takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Patch(path string, h http.Handler) {
	m.t.Handle(http.MethodPatch, m.path+path, treeMuxWrapper(h))
}

// Delete takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Delete(path string, h http.Handler) {
	m.t.Handle(http.MethodDelete, m.path+path, treeMuxWrapper(h))
}

// Head takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Head(path string, h http.Handler) {
	m.t.Handle(http.MethodHead, m.path+path, treeMuxWrapper(h))
}

// Trace takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Trace(path string, h http.Handler) {
	m.t.Handle(http.MethodTrace, m.path+path, treeMuxWrapper(h))
}

// Connect takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Connect(path string, h http.Handler) {
	m.t.Handle(http.MethodConnect, m.path+path, treeMuxWrapper(h))
}

// ServeHTTP satisfies the http.Handler interface.
func (m *TreeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.t.ServeHTTP(w, r)
}

func treeMuxWrapper(next http.Handler) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		next.ServeHTTP(w, r)
	}
}
