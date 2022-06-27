package server

import (
	"net/http"

	"github.com/sheva0914/MSc_2021-22_Mock_webserver/pkg/handler"
)

const port = ":8000"

type Server struct {
	s *http.Server
}

func New() *Server {
	h := http.NewServeMux()
	h.HandleFunc("/", handler.MainPageHandler)
	h.HandleFunc("/login", handler.LoginHandler)
	h.HandleFunc("/success", handler.SuccessPageHandler)
	h.HandleFunc("/failure", handler.FailurePageHandler)

	return &Server{
		s: &http.Server{
			Addr:    port,
			Handler: h,
		},
	}
}

func (s *Server) Run() error {
	return s.s.ListenAndServe()
}

/*
# Reference
- https://pkg.go.dev/net/http

# Line Count
- Total:      32
- Reused:     0
- Written:    28
- Referenced: 5
*/
