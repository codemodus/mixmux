package mixmux

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

var (
	WildcardRegex = `[a-zA-Z0-9=\-\/\.]+`
)

type MixMuxer interface {
	OPTIONS(path string, h http.Handler)
	GET(path string, h http.Handler)
	POST(path string, h http.Handler)
	PUT(path string, h http.Handler)
	PATCH(path string, h http.Handler)
	DELETE(path string, h http.Handler)
	HEAD(path string, h http.Handler)
}

type apemux struct {
	*mux.Router
}

func NewApeMux() *apemux {
	return &apemux{mux.NewRouter()}
}

func (am *apemux) OPTIONS(path string, h http.Handler) {
	am.Handle(path, h).Methods("OPTIONS")
}

func (am *apemux) GET(path string, h http.Handler) {
	am.Handle(path, h).Methods("GET")
}

func (am *apemux) POST(path string, h http.Handler) {
	am.Handle(path, h).Methods("POST")
}

func (am *apemux) PUT(path string, h http.Handler) {
	am.Handle(path, h).Methods("PUT")
}

func (am *apemux) PATCH(path string, h http.Handler) {
	am.Handle(path, h).Methods("PATCH")
}

func (am *apemux) DELETE(path string, h http.Handler) {
	am.Handle(path, h).Methods("DELETE")
}

func (am *apemux) HEAD(path string, h http.Handler) {
	am.Handle(path, h).Methods("HEAD")
}

type router struct {
	*httprouter.Router
}

func NewHttpRouter() *router {
	return &router{httprouter.New()}
}

func (r *router) OPTIONS(path string, h http.Handler) {
	r.Handle("OPTIONS", path, httpRouterWrapHandler(h))
}

func (r *router) GET(path string, h http.Handler) {
	r.Handle("GET", path, httpRouterWrapHandler(h))
}

func (r *router) POST(path string, h http.Handler) {
	r.Handle("POST", path, httpRouterWrapHandler(h))
}

func (r *router) PUT(path string, h http.Handler) {
	r.Handle("PUT", path, httpRouterWrapHandler(h))
}

func (r *router) PATCH(path string, h http.Handler) {
	r.Handle("PATCH", path, httpRouterWrapHandler(h))
}

func (r *router) DELETE(path string, h http.Handler) {
	r.Handle("DELETE", path, httpRouterWrapHandler(h))
}

func (r *router) HEAD(path string, h http.Handler) {
	r.Handle("HEAD", path, httpRouterWrapHandler(h))
}

func httpRouterWrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

type treemux struct {
	*httptreemux.TreeMux
}

func NewTreeMux() *treemux {
	return &treemux{httptreemux.New()}
}

func (tm *treemux) OPTIONS(path string, h http.Handler) {
	tm.Handle("OPTIONS", path, treeMuxWrapHandler(h))
}

func (tm *treemux) GET(path string, h http.Handler) {
	tm.Handle("GET", path, treeMuxWrapHandler(h))
}

func (tm *treemux) POST(path string, h http.Handler) {
	tm.Handle("POST", path, treeMuxWrapHandler(h))
}

func (tm *treemux) PUT(path string, h http.Handler) {
	tm.Handle("PUT", path, treeMuxWrapHandler(h))
}

func (tm *treemux) PATCH(path string, h http.Handler) {
	tm.Handle("PATCH", path, treeMuxWrapHandler(h))
}

func (tm *treemux) DELETE(path string, h http.Handler) {
	tm.Handle("DELETE", path, treeMuxWrapHandler(h))
}

func (tm *treemux) HEAD(path string, h http.Handler) {
	tm.Handle("HEAD", path, treeMuxWrapHandler(h))
}

func treeMuxWrapHandler(h http.Handler) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		h.ServeHTTP(w, r)
	}
}
