package server

import (
	"fmt"
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
		newErr := fmt.Errorf("[server.New()] Failed to initialise handler pkg: %v", err)
		log.Println(newErr.Error())
		return nil, newErr
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
- Total:      35
- Reused:     0
- Written:    30
- Referenced: 5
*/
