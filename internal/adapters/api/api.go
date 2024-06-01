package api

import (
	"fmt"
	"net/http"
)

type api struct {
	addr   string
	server *http.ServeMux
}

func (a *api) Post(path string, handlerFunc http.HandlerFunc) {
	pattern := fmt.Sprintf("POST %s", path)
	a.server.HandleFunc(pattern, handlerFunc)
}

func (a *api) Get(path string, handlerFunc http.HandlerFunc) {
	pattern := fmt.Sprintf("GET %s", path)
	a.server.HandleFunc(pattern, handlerFunc)
}

func (a *api) Run() error {
	s := http.Server{
		Addr:    a.addr,
		Handler: a.server,
	}

	return s.ListenAndServe()
}

func NewApiServer(addr string) *api {

	return &api{addr: addr, server: http.NewServeMux()}

}
