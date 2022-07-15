package main

import (
	"log"

	"github.com/sheva0914/MSc_2021-22_Mock_webserver/pkg/server"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatalln("Failed to initialise server: ", err)
	}

	if err := s.Run(); err != nil {
		log.Fatalln("Failed to run server: ", err)
	}
}

/*
# Line Count
- Total:      14
- Reused:     0
- Written:    14
- Referenced: 0
*/
