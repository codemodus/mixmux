// Package mixmux wraps HTTPRouter and HTTPTreeMux to provide more consistent
// and more idiomatic APIs.
package mixmux

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/julienschmidt/httprouter"
)

// MixMuxer defines available methods for wrapped multiplexers.
type MixMuxer interface {
	Options(path string, h http.Handler)
	Get(path string, h http.Handler)
	Post(path string, h http.Handler)
	Put(path string, h http.Handler)
	Patch(path string, h http.Handler)
	Delete(path string, h http.Handler)
	Head(path string, h http.Handler)
}

// Router wraps HTTPRouter.
type Router struct {
	*httprouter.Router
	path string
}

// NewHTTPRouter returns a wrapped HTTPRouter.
func NewHTTPRouter() *Router {
	return &Router{httprouter.New(), ""}
}

// Group takes a path and returns a new Router around the original HTTPRouter.
func (r *Router) Group(path string) *Router {
	return &Router{r.Router, r.path + path}
}

// Options takes a path and http.Handler and adds them to the mux.
func (r *Router) Options(path string, h http.Handler) {
	r.Handle("OPTIONS", r.path+path, httpRouterWrapHandler(h))
}

// Get takes a path and http.Handler and adds them to the mux.
func (r *Router) Get(path string, h http.Handler) {
	r.Handle("GET", r.path+path, httpRouterWrapHandler(h))
}

// Post takes a path and http.Handler and adds them to the mux.
func (r *Router) Post(path string, h http.Handler) {
	r.Handle("POST", r.path+path, httpRouterWrapHandler(h))
}

// Put takes a path and http.Handler and adds them to the mux.
func (r *Router) Put(path string, h http.Handler) {
	r.Handle("PUT", r.path+path, httpRouterWrapHandler(h))
}

// Patch takes a path and http.Handler and adds them to the mux.
func (r *Router) Patch(path string, h http.Handler) {
	r.Handle("PATCH", r.path+path, httpRouterWrapHandler(h))
}

// Delete takes a path and http.Handler and adds them to the mux.
func (r *Router) Delete(path string, h http.Handler) {
	r.Handle("DELETE", r.path+path, httpRouterWrapHandler(h))
}

// Head takes a path and http.Handler and adds them to the mux.
func (r *Router) Head(path string, h http.Handler) {
	r.Handle("HEAD", r.path+path, httpRouterWrapHandler(h))
}

func httpRouterWrapHandler(next http.Handler) httprouter.Handle {
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
	return &TreeMux{httptreemux.New()}
}

// Group takes a path and returns a new TreeMux around the original HTTPTreeMux.
func (tm *TreeMux) Group(path string) *TreeMux {
	return &Router{tm.TreeMux, tm.path + path}
}

// Options takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Options(path string, h http.Handler) {
	tm.Handle("OPTIONS", tm.path+path, treeMuxWrapHandler(h))
}

// Get takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Get(path string, h http.Handler) {
	tm.Handle("GET", tm.path+path, treeMuxWrapHandler(h))
}

// Post takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Post(path string, h http.Handler) {
	tm.Handle("POST", tm.path+path, treeMuxWrapHandler(h))
}

// Put takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Put(path string, h http.Handler) {
	tm.Handle("PUT", tm.path+path, treeMuxWrapHandler(h))
}

// Patch takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Patch(path string, h http.Handler) {
	tm.Handle("PATCH", tm.path+path, treeMuxWrapHandler(h))
}

// Delete takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Delete(path string, h http.Handler) {
	tm.Handle("DELETE", tm.path+path, treeMuxWrapHandler(h))
}

// Head takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Head(path string, h http.Handler) {
	tm.Handle("HEAD", tm.path+path, treeMuxWrapHandler(h))
}

func treeMuxWrapHandler(next http.Handler) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		next.ServeHTTP(w, r)
	}
}
