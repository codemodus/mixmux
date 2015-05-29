package mixmux_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codemodus/mixmux"
	"github.com/dimfeld/httptreemux"
	"github.com/julienschmidt/httprouter"
)

const (
	Options = "OPTIONS"
	Get     = "GET"
	Post    = "POST"
	Put     = "PUT"
	Patch   = "PATCH"
	Delete  = "DELETE"
	Head    = "HEAD"
)

var (
	tMap = map[string]string{
		"/grouped/options": Options,
		"/grouped/get":     Get,
		"/grouped/post":    Post,
		"/grouped/put":     Put,
		"/grouped/patch":   Patch,
		"/grouped/delete":  Delete,
		"/grouped/head":    Head,
	}
)

func Example() {
	handler := http.HandlerFunc(methodHandler)

	mux := mixmux.NewRouter()
	mux.Get("/get", handler)
	mux.Post("/post", handler)

	muxGroup := mux.Group("/grouped")
	muxGroup.Get("/get0", handler)
	muxGroup.Get("/get1", handler)

	server := httptest.NewServer(mux)

	rBody0, err := getReqBody(server.URL+"/get", Get)
	if err != nil {
		fmt.Println(err)
	}

	rBody1, err := getReqBody(server.URL+"/post", Post)
	if err != nil {
		fmt.Println(err)
	}

	rBodyGrouped0, err := getReqBody(server.URL+"/grouped/get0", Get)
	if err != nil {
		fmt.Println(err)
	}

	rBodyGrouped1, err := getReqBody(server.URL+"/grouped/get1", Get)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get Body:", rBody0)
	fmt.Println("Post Body:", rBody1)
	fmt.Println("Grouped Bodies:", rBodyGrouped0, rBodyGrouped1)

	// Output:
	// Get Body: GET
	// Post Body: POST
	// Grouped Bodies: GET GET
}

func TestRouterMethods(t *testing.T) {
	h := http.HandlerFunc(methodHandler)
	m := mixmux.NewRouter()
	mg := m.Group("/grouped")
	mg.Options("/options", h)
	mg.Get("/get", h)
	mg.Post("/post", h)
	mg.Put("/put", h)
	mg.Patch("/patch", h)
	mg.Delete("/delete", h)
	mg.Head("/head", h)
	s := httptest.NewServer(m)

	for k, v := range tMap {
		rb, err := getReqBody(s.URL+k, v)
		if err != nil {
			t.Fatal(err)
		}
		want := v
		got := rb
		if got != want {
			t.Errorf("Body = %v, want %v", got, want)
		}
	}
}

func TestTreeMuxMethods(t *testing.T) {
	h := http.HandlerFunc(methodHandler)
	m := mixmux.NewTreeMux()
	mg := m.Group("/grouped")
	mg.Options("/options", h)
	mg.Get("/get", h)
	mg.Post("/post", h)
	mg.Put("/put", h)
	mg.Patch("/patch", h)
	mg.Delete("/delete", h)
	mg.Head("/head", h)
	s := httptest.NewServer(m)

	for k, v := range tMap {
		rb, err := getReqBody(s.URL+k, v)
		if err != nil {
			t.Fatal(err)
		}
		want := v
		got := rb
		if got != want {
			t.Errorf("Body = %v, want %v", got, want)
		}
	}
}

func getReqBody(url, method string) (string, error) {
	r, err := http.NewRequest(method, url, nil)
	cl := &http.Client{}
	resp, err := cl.Do(r)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if len(body) == 0 {
		body = []byte(resp.Header.Get(Head))
	}
	_ = resp.Body.Close()
	return string(body), nil
}

func methodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == Head {
		w.Header().Add(Head, Head)
	}
	_, _ = w.Write([]byte(r.Method))
	return
}

func BenchmarkHTTPServeMux0(b *testing.B) {
	m := http.NewServeMux()
	m.Handle("/test/test/1/test",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			return
		}),
	)
	s := httptest.NewServer(m)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		re0, err := http.Get(s.URL + "/test/test/1/test")
		if err != nil {
			b.Error(err)
		}
		_ = re0.Body.Close()
	}
}

func BenchmarkHTTPTreeMux2(b *testing.B) {
	m := httptreemux.New()
	m.Handle("GET", "/test/test/:id/:last",
		func(w http.ResponseWriter, r *http.Request, m map[string]string) {
			_, _ = m["id"]
			return
		},
	)
	s := httptest.NewServer(m)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		re0, err := http.Get(s.URL + "/test/test/1/test")
		if err != nil {
			b.Error(err)
		}
		_ = re0.Body.Close()
	}
}

func BenchmarkHTTPRouter2(b *testing.B) {
	m := httprouter.New()
	m.Handle("GET", "/test/test/:id/:last",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			return
		},
	)
	s := httptest.NewServer(m)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		re0, err := http.Get(s.URL + "/test/test/1/test")
		if err != nil {
			b.Error(err)
		}
		_ = re0.Body.Close()
	}
}

func BenchmarkMixmuxTreeMux2(b *testing.B) {
	m := mixmux.NewTreeMux()
	m.Get("/test/test/:id/:last",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			return
		}),
	)
	s := httptest.NewServer(m)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		re0, err := http.Get(s.URL + "/test/test/1/test")
		if err != nil {
			b.Error(err)
		}
		_ = re0.Body.Close()
	}
}

func BenchmarkMixmuxRouter2(b *testing.B) {
	m := mixmux.NewRouter()
	m.Get("/test/test/:id/:last",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			return
		}),
	)
	s := httptest.NewServer(m)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		re0, err := http.Get(s.URL + "/test/test/1/test")
		if err != nil {
			b.Error(err)
		}
		_ = re0.Body.Close()
	}
}

func BenchmarkMixmuxRouterGroup2(b *testing.B) {
	m := mixmux.NewRouter()
	mg := m.Group("/test")
	mg.Get("/test/:id/:last",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			return
		}),
	)
	s := httptest.NewServer(m)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		re0, err := http.Get(s.URL + "/test/test/1/test")
		if err != nil {
			b.Error(err)
		}
		_ = re0.Body.Close()
	}
}
