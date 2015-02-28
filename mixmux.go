package mixmux

import (
	"github.com/dimfeld/httptreemux"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	httpRouter = 1
	treeMux    = 2
	apeMux     = 3
)

var (
	routerType int
)

type MixMux interface {
	OPTIONS(path string, h http.Handler)
	GET(path string, h http.Handler)
	POST(path string, h http.Handler)
	PUT(path string, h http.Handler)
	PATCH(path string, h http.Handler)
	DELETE(path string, h http.Handler)
	HEAD(path string, h http.Handler)
}

func NewHttpRouter() *router {
	routerType = httpRouter
	return &router{httprouter.New()}
}

func NewTreeMux() *treemux {
	routerType = treeMux
	return &treemux{httptreemux.New()}
}

func NewApeMux() *apemux {
	routerType = apeMux
	return &apemux{mux.NewRouter()}
}

type Params [1]map[string]string

func (ps Params) ByName(name string) string {
	if v, ok := ps[0][name]; ok {
		return v
	}

	return ""
}

func GetParams(r *http.Request) Params {
	switch routerType {
	case httpRouter, treeMux:
		return context.Get(r, "params").(Params)
	case apeMux:
		return Params{mux.Vars(r)}
	}

	return Params{}
}
