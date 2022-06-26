package main

import (
	"log"

	"github.com/sheva0914/MSc_2021-22_Mock_webserver/pkg/server"
)

func main() {
	s := server.New()
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

/*
# Line Count
- Total:      14
- Reused:     0
- Written:    14
- Referenced: 0
*/
