package server

import (
	"net/http"
	"time"
)

var Instance = &Server{Running: false}

type Server struct {
	Http    *http.Server
	Running bool
}

func (s *Server) Start(port string) {
	Routes = make(map[string]Route)

	s.Http = &http.Server{
		Handler:      &Router{},
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	s.Running = true

	go s.Http.ListenAndServe()
}
