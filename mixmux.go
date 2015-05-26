// Package mixmux wraps HTTPRouter and HTTPTreeMux to provide more consistent
// and idiomatic APIs, along with route grouping.
package mixmux

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/julienschmidt/httprouter"
)

// Router wraps HTTPRouter.
type Router struct {
	*httprouter.Router
	path string
}

// NewRouter returns a wrapped HTTPRouter.
func NewRouter() *Router {
	return &Router{httprouter.New(), ""}
}

// Group takes a path and returns a new Router wrapping the original Router.
func (r *Router) Group(path string) *Router {
	return &Router{r.Router, r.path + path}
}

// Options takes a path and http.Handler and adds them to the mux.
func (r *Router) Options(path string, h http.Handler) {
	r.Handle("OPTIONS", r.path+path, routerWrapper(h))
}

// Get takes a path and http.Handler and adds them to the mux.
func (r *Router) Get(path string, h http.Handler) {
	r.Handle("GET", r.path+path, routerWrapper(h))
}

// Post takes a path and http.Handler and adds them to the mux.
func (r *Router) Post(path string, h http.Handler) {
	r.Handle("POST", r.path+path, routerWrapper(h))
}

// Put takes a path and http.Handler and adds them to the mux.
func (r *Router) Put(path string, h http.Handler) {
	r.Handle("PUT", r.path+path, routerWrapper(h))
}

// Patch takes a path and http.Handler and adds them to the mux.
func (r *Router) Patch(path string, h http.Handler) {
	r.Handle("PATCH", r.path+path, routerWrapper(h))
}

// Delete takes a path and http.Handler and adds them to the mux.
func (r *Router) Delete(path string, h http.Handler) {
	r.Handle("DELETE", r.path+path, routerWrapper(h))
}

// Head takes a path and http.Handler and adds them to the mux.
func (r *Router) Head(path string, h http.Handler) {
	r.Handle("HEAD", r.path+path, routerWrapper(h))
}

func routerWrapper(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		next.ServeHTTP(w, r)
	}
}

// TreeMux wraps HTTPTreeMux.
type TreeMux struct {
	*httptreemux.TreeMux
	path string
}

// NewTreeMux returns a wrapped HTTPTreeMux.
func NewTreeMux() *TreeMux {
	return &TreeMux{httptreemux.New(), ""}
}

// Group takes a path and returns a new TreeMux wrapping the original TreeMux.
func (tm *TreeMux) Group(path string) *TreeMux {
	return &TreeMux{tm.TreeMux, tm.path + path}
}

// Options takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Options(path string, h http.Handler) {
	tm.Handle("OPTIONS", tm.path+path, treeMuxWrapper(h))
}

// Get takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Get(path string, h http.Handler) {
	tm.Handle("GET", tm.path+path, treeMuxWrapper(h))
}

// Post takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Post(path string, h http.Handler) {
	tm.Handle("POST", tm.path+path, treeMuxWrapper(h))
}

// Put takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Put(path string, h http.Handler) {
	tm.Handle("PUT", tm.path+path, treeMuxWrapper(h))
}

// Patch takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Patch(path string, h http.Handler) {
	tm.Handle("PATCH", tm.path+path, treeMuxWrapper(h))
}

// Delete takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Delete(path string, h http.Handler) {
	tm.Handle("DELETE", tm.path+path, treeMuxWrapper(h))
}

// Head takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Head(path string, h http.Handler) {
	tm.Handle("HEAD", tm.path+path, treeMuxWrapper(h))
}

func treeMuxWrapper(next http.Handler) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		next.ServeHTTP(w, r)
	}
}
