package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	templateHTMLFile = "./web/template.html"
	validUsername    = "admin"
	validPassword    = "password"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	// Load template HTML
	loginPage, err := os.ReadFile(templateHTMLFile)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to read template HTML", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(loginPage))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Stop processing if request is NOT POST
	if r.Method != "POST" {
		log.Println("Invalid HTTP method")
		http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		return
	}

	// Obtain username and password
	un := r.FormValue("username")
	pw := r.FormValue("password")

	if un != validUsername && pw != validPassword {
		log.Println("Invalid username or password")
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	fmt.Fprint(w, "Login successful\n")
}

func SuccessPageHandler(w http.ResponseWriter, r *http.Request) {

}

func FailurePageHandler(w http.ResponseWriter, r *http.Request) {

}

/*
# Reference
- https://pkg.go.dev/net/http
- https://gobyexample.com/reading-files

# Line Count
- Total:      28
- Reused:     0
- Written:    3
- Referenced: 25
*/
