package server

import (
	"log"
	"net/http"

	"github.com/sheva0914/MSc_2021-22_Mock_webserver/pkg/handler"
)

const port = ":8000"

type Server struct {
	s *http.Server
}

func New() (*Server, error) {
	h, err := handler.Init()
	if err != nil {
		log.Println("Failed to initialise handler pkg: ", err)
		return nil, err
	}

	return &Server{
		s: &http.Server{
			Addr:    port,
			Handler: h,
		},
	}, nil
}

func (s *Server) Run() error {
	return s.s.ListenAndServe()
}

/*
# Reference
- https://pkg.go.dev/net/http

# Line Count
- Total:      28
- Reused:     0
- Written:    23
- Referenced: 5
*/
