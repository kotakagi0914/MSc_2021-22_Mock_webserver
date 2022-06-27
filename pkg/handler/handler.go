package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	templateHTMLFile    = "./web/login-page.html"
	validUsername       = "admin"
	validPassword       = "password"
	recaptchaScoreQuery = "recaptchaScore"
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

	if un != validUsername || pw != validPassword {
		log.Println("Invalid username or password")
		http.Redirect(w, r, "/failure", http.StatusFound)
		return
	}

	// Verify reCAPTCHA score by sending request to Google
	// Dummy valu for now
	score := "0.5"

	http.Redirect(w, r, "/success?"+recaptchaScoreQuery+"="+score, http.StatusFound)
}

func SuccessPageHandler(w http.ResponseWriter, r *http.Request) {
	score := r.URL.Query().Get(recaptchaScoreQuery)
	floatScore, err := strconv.ParseFloat(score, 64)
	if err != nil {
		log.Println("Invalid score type: ", err)
		http.Error(w, "Invalid score type", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Login Success: %f", floatScore)
}

func FailurePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Login Failure")
	http.Error(w, "Login Failure", http.StatusBadRequest)
}

/*
# Reference
- https://pkg.go.dev/net/http
- https://gobyexample.com/http-servers
- https://gobyexample.com/reading-files

# Line Count
- Total:      70
- Reused:     0
- Written:    64
- Referenced: 6
*/
