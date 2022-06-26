package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

func main() {
	http.HandleFunc("/hello", hello)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

/*
# Reference
- https://gobyexample.com/http-servers

# Line Count
- Total:      17
- Reused:     0
- Written:    3
- Referenced: 14
*/
