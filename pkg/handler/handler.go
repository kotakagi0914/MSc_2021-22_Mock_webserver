package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const templateHTMLPath = "./web/"
const templateHTMLFile = "template.html"

func MainPageHandler(w http.ResponseWriter, req *http.Request) {
	// Load template HTML
	loginPage, err := os.ReadFile(templateHTMLPath + templateHTMLFile)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to read template HTML", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(loginPage))
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	// Stop processing if request is NOT POST
}

func SuccessPageHandler(w http.ResponseWriter, req *http.Request) {

}

func FailurePageHandler(w http.ResponseWriter, req *http.Request) {

}

/*
# Reference
- https://gobyexample.com/reading-files

# Line Count
- Total:      28
- Reused:     0
- Written:    3
- Referenced: 25
*/
