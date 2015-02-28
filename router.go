package mixmux

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type router struct {
	*httprouter.Router
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
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		newPs := map[string]string{}
		for i := range ps {
			newPs[ps[i].Key] = ps[i].Value
		}
		context.Set(r, "params", Params{newPs})
		h.ServeHTTP(w, r)
	}
}
