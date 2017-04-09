package mixmux

import (
	"net/http"
	"strings"

	"github.com/dimfeld/httptreemux"
)

// TreeMux wraps HTTPTreeMux.
type TreeMux struct {
	t    *httptreemux.TreeMux
	path string
	reg  map[string][]string
}

// NewTreeMux returns a wrapped HTTPTreeMux.
func NewTreeMux(opts *Options) *TreeMux {
	t := &TreeMux{
		t:    httptreemux.New(),
		path: "",
		reg:  make(map[string][]string),
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

func (m *TreeMux) addToReg(method, path string) {
	_, ok := m.reg[path]
	if !ok {
		m.reg[path] = []string{}
	}

	m.reg[path] = append(m.reg[path], method)
}

func (m *TreeMux) handle(method, path string, handler http.Handler) {
	m.t.Handle(method, path, treeMuxWrapper(handler))
	m.addToReg(method, path)
}

func (m *TreeMux) lookup(method, path string) (bool, bool) {
	ckSlsh := false
	if m.t.RedirectTrailingSlash && path[len(path)-1] != '/' {
		ckSlsh = true
	}

	meths, ok := m.reg[path]
	if !ok {
		return false, ckSlsh
	}

	for _, m := range meths {
		if m == method {
			return true, false
		}
	}

	return false, ckSlsh
}

// Group takes a path and returns a new TreeMux wrapping the original TreeMux.
func (m *TreeMux) Group(path string) *TreeMux {
	return &TreeMux{m.t, m.path + path, m.reg}
}

// Options takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Options(path string, h http.Handler) {
	m.handle(http.MethodOptions, m.path+path, h)
}

// Get takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Get(path string, h http.Handler) {
	m.handle(http.MethodGet, m.path+path, h)
}

// Post takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Post(path string, h http.Handler) {
	m.handle(http.MethodPost, m.path+path, h)
}

// Put takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Put(path string, h http.Handler) {
	m.handle(http.MethodPut, m.path+path, h)
}

// Patch takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Patch(path string, h http.Handler) {
	m.handle(http.MethodPatch, m.path+path, h)
}

// Delete takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Delete(path string, h http.Handler) {
	m.handle(http.MethodDelete, m.path+path, h)
}

// Head takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Head(path string, h http.Handler) {
	m.handle(http.MethodHead, m.path+path, h)
}

// Trace takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Trace(path string, h http.Handler) {
	m.handle(http.MethodTrace, m.path+path, h)
}

// Connect takes a path and http.Handler and adds them to the mux.
func (m *TreeMux) Connect(path string, h http.Handler) {
	m.handle(http.MethodConnect, m.path+path, h)
}

// ServeHTTP satisfies the http.Handler interface.
func (m *TreeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.t.ServeHTTP(w, r)
}

// CORSMethods ... TODO:
func (m *TreeMux) CORSMethods(path string, handlerWrappers ...func(http.Handler) http.Handler) {
	x, s := m.lookup(http.MethodOptions, path)
	if s {
		x, _ = m.lookup(http.MethodOptions, path+"/")
	}
	if !x {
		return
	}

	ms := []string{http.MethodOptions}

	for _, v := range methods {
		if v == http.MethodOptions {
			continue
		}

		x, s = m.lookup(v, path)
		if s {
			x, _ = m.lookup(v, path+"/")
		}
		if !x {
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

	m.Options(path, fn)
}

func treeMuxWrapper(next http.Handler) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		next.ServeHTTP(w, r)
	}
}
