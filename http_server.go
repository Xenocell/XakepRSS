package main

import "net/http"

type server struct {
	*http.Server
}

func NewHttpServer(router http.Handler, port string) *server {
	httpServer := &http.Server{
		Handler: router,
		Addr:    port,
	}
	return &server{httpServer}
}

func (s *server) Start() {
	s.ListenAndServe()
}
