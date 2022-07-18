package main

import (
	"log"

	"github.com/sheva0914/MSc_2021-22_Mock_webserver/pkg/server"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatalln("[main] Failed to initialise server: ", err)
	}

	if err := s.Run(); err != nil {
		log.Fatalln("[main] Failed to run server: ", err)
	}
}

/*
# Line Count
- Total:      18
- Reused:     0
- Written:    18
- Referenced: 0
*/
