package mixmux

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

var (
	// WildcardRegex is a convenience var for working with Gorilla Mux.
	WildcardRegex = `[a-zA-Z0-9=\-\/\.]+`
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

// ApeMux wraps Gorilla Mux.
type ApeMux struct {
	*mux.Router
}

// NewApeMux returns a wrapped Gorilla Mux
func NewApeMux() *ApeMux {
	return &ApeMux{mux.NewRouter()}
}

// Options takes a path and http.Handler and adds them to the mux.
func (am *ApeMux) Options(path string, h http.Handler) {
	am.Handle(path, h).Methods("OPTIONS")
}

// Get takes a path and http.Handler and adds them to the mux.
func (am *ApeMux) Get(path string, h http.Handler) {
	am.Handle(path, h).Methods("GET")
}

// Post takes a path and http.Handler and adds them to the mux.
func (am *ApeMux) Post(path string, h http.Handler) {
	am.Handle(path, h).Methods("POST")
}

// Put takes a path and http.Handler and adds them to the mux.
func (am *ApeMux) Put(path string, h http.Handler) {
	am.Handle(path, h).Methods("PUT")
}

// Patch takes a path and http.Handler and adds them to the mux.
func (am *ApeMux) Patch(path string, h http.Handler) {
	am.Handle(path, h).Methods("PATCH")
}

// Delete takes a path and http.Handler and adds them to the mux.
func (am *ApeMux) Delete(path string, h http.Handler) {
	am.Handle(path, h).Methods("DELETE")
}

// Head takes a path and http.Handler and adds them to the mux.
func (am *ApeMux) Head(path string, h http.Handler) {
	am.Handle(path, h).Methods("HEAD")
}

// Router wraps HTTPRouter.
type Router struct {
	*httprouter.Router
}

// NewHTTPRouter returns a wrapped HTTPRouter.
func NewHTTPRouter() *Router {
	return &Router{httprouter.New()}
}

// Options takes a path and http.Handler and adds them to the mux.
func (r *Router) Options(path string, h http.Handler) {
	r.Handle("OPTIONS", path, httpRouterWrapHandler(h))
}

// Get takes a path and http.Handler and adds them to the mux.
func (r *Router) Get(path string, h http.Handler) {
	r.Handle("GET", path, httpRouterWrapHandler(h))
}

// Post takes a path and http.Handler and adds them to the mux.
func (r *Router) Post(path string, h http.Handler) {
	r.Handle("POST", path, httpRouterWrapHandler(h))
}

// Put takes a path and http.Handler and adds them to the mux.
func (r *Router) Put(path string, h http.Handler) {
	r.Handle("PUT", path, httpRouterWrapHandler(h))
}

// Patch takes a path and http.Handler and adds them to the mux.
func (r *Router) Patch(path string, h http.Handler) {
	r.Handle("PATCH", path, httpRouterWrapHandler(h))
}

// Delete takes a path and http.Handler and adds them to the mux.
func (r *Router) Delete(path string, h http.Handler) {
	r.Handle("DELETE", path, httpRouterWrapHandler(h))
}

// Head takes a path and http.Handler and adds them to the mux.
func (r *Router) Head(path string, h http.Handler) {
	r.Handle("HEAD", path, httpRouterWrapHandler(h))
}

func httpRouterWrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

// TreeMux wraps HTTPTreeMux.
type TreeMux struct {
	*httptreemux.TreeMux
}

// NewTreeMux returns a wrapped HTTPTreeMux.
func NewTreeMux() *TreeMux {
	return &TreeMux{httptreemux.New()}
}

// Options takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Options(path string, h http.Handler) {
	tm.Handle("OPTIONS", path, treeMuxWrapHandler(h))
}

// Get takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Get(path string, h http.Handler) {
	tm.Handle("GET", path, treeMuxWrapHandler(h))
}

// Post takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Post(path string, h http.Handler) {
	tm.Handle("POST", path, treeMuxWrapHandler(h))
}

// Put takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Put(path string, h http.Handler) {
	tm.Handle("PUT", path, treeMuxWrapHandler(h))
}

// Patch takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Patch(path string, h http.Handler) {
	tm.Handle("PATCH", path, treeMuxWrapHandler(h))
}

// Delete takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Delete(path string, h http.Handler) {
	tm.Handle("DELETE", path, treeMuxWrapHandler(h))
}

// Head takes a path and http.Handler and adds them to the mux.
func (tm *TreeMux) Head(path string, h http.Handler) {
	tm.Handle("HEAD", path, treeMuxWrapHandler(h))
}

func treeMuxWrapHandler(h http.Handler) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		h.ServeHTTP(w, r)
	}
}
