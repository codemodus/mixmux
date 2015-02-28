package mixmux

import (
	"github.com/gorilla/mux"
	"net/http"
)

var (
	WildcardRegex = `[a-zA-Z0-9=\-\/\.]+`
)

type apemux struct {
	*mux.Router
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
