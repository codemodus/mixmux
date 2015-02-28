package mixmux

import (
	"github.com/dimfeld/httptreemux"
	"github.com/gorilla/context"
	"net/http"
)

type treemux struct {
	*httptreemux.TreeMux
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
	return func(w http.ResponseWriter, r *http.Request, ps map[string]string) {
		context.Set(r, "params", Params{ps})
		h.ServeHTTP(w, r)
	}
}
