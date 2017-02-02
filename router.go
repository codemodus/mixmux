package mixmux

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Router wraps HTTPRouter.
type Router struct {
	r    *httprouter.Router
	path string
}

// NewRouter returns a wrapped HTTPRouter.
func NewRouter(opts *Options) *Router {
	// TODO:
	r := &Router{
		path: "",
	}

	if opts == nil {
		opts = &Options{}
	}

	r.r = &httprouter.Router{
		RedirectTrailingSlash:  opts.RedirectTrailingSlash,
		RedirectFixedPath:      opts.RedirectFixedPath,
		HandleMethodNotAllowed: opts.HandleMethodNotAllowed,
		NotFound:               opts.NotFound,
		MethodNotAllowed:       opts.MethodNotAllowed,
	}

	return r
}

// Group takes a path and returns a new Router wrapping the original Router.
func (m *Router) Group(path string) *Router {
	return &Router{m.r, m.path + path}
}

// Options takes a path and http.Handler and adds them to the mux.
func (m *Router) Options(path string, h http.Handler) {
	m.r.Handler(http.MethodOptions, m.path+path, h)
}

// Get takes a path and http.Handler and adds them to the mux.
func (m *Router) Get(path string, h http.Handler) {
	m.r.Handler(http.MethodGet, m.path+path, h)
}

// Post takes a path and http.Handler and adds them to the mux.
func (m *Router) Post(path string, h http.Handler) {
	m.r.Handler(http.MethodPost, m.path+path, h)
}

// Put takes a path and http.Handler and adds them to the mux.
func (m *Router) Put(path string, h http.Handler) {
	m.r.Handler(http.MethodPut, m.path+path, h)
}

// Patch takes a path and http.Handler and adds them to the mux.
func (m *Router) Patch(path string, h http.Handler) {
	m.r.Handler(http.MethodPatch, m.path+path, h)
}

// Delete takes a path and http.Handler and adds them to the mux.
func (m *Router) Delete(path string, h http.Handler) {
	m.r.Handler(http.MethodDelete, m.path+path, h)
}

// Head takes a path and http.Handler and adds them to the mux.
func (m *Router) Head(path string, h http.Handler) {
	m.r.Handler(http.MethodHead, m.path+path, h)
}

// Trace takes a path and http.Handler and adds them to the mux.
func (m *Router) Trace(path string, h http.Handler) {
	m.r.Handler(http.MethodTrace, m.path+path, h)
}

// Connect takes a path and http.Handler and adds them to the mux.
func (m *Router) Connect(path string, h http.Handler) {
	m.r.Handler(http.MethodConnect, m.path+path, h)
}

// OptionsHeaders ... TODO:
func (m *Router) OptionsHeaders(path string, handlerWrappers ...func(http.Handler) http.Handler) {
	h, _, s := m.r.Lookup(http.MethodOptions, path)
	if s {
		h, _, _ = m.r.Lookup(http.MethodOptions, path+"/")
	}
	if h != nil {
		return
	}

	ms := []string{http.MethodOptions}

	for _, v := range methods {
		if v == http.MethodOptions {
			continue
		}

		h, _, s = m.r.Lookup(v, path)
		if s {
			h, _, _ = m.r.Lookup(v, path+"/")
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

	for _, wrap := range handlerWrappers {
		if wrap != nil {
			fn = wrap(fn)
		}
	}

	m.r.Handler(http.MethodOptions, path, fn)
}

// ServeHTTP satisfies the http.Handler interface.
func (m *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.r.ServeHTTP(w, r)
}
