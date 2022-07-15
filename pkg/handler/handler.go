package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	templateHTMLFile    = "./web/template/login-page.html"
	secretFile          = "./.secret"
	validUsername       = "admin"
	validPassword       = "password"
	recaptchaScoreQuery = "recaptchaScore"
)

// Struct for reCAPTCHA secret containing site-key and secret-key
type reCAPTCHASecretStruct struct {
	SiteKey   string `json:"site-key"`
	SecretKey string `json:"secret-key"`
}

var (
	recaptchaSecret reCAPTCHASecretStruct
	loginPageHTML   string
)

func Init() (*http.ServeMux, error) {
	// Initialise HTTP server handler
	h := http.NewServeMux()
	h.HandleFunc("/", MainPageHandler)
	h.HandleFunc("/login", LoginHandler)
	h.HandleFunc("/success", SuccessPageHandler)
	h.HandleFunc("/failure", FailurePageHandler)

	// Load login page template
	loginPageByte, err := os.ReadFile(templateHTMLFile)
	if err != nil {
		log.Println("Failed to load login page template", err)
		return nil, err
	}
	loginPageHTML = string(loginPageByte)

	// Load site-key and
	secretJsonByte, err := os.ReadFile(secretFile)
	if err != nil {
		log.Println("Failed to read reCAPTCHA secret file: ", err)
		return nil, err
	}

	// Unmarshal JSON byte into Golang struct
	if err := json.Unmarshal(secretJsonByte, &recaptchaSecret); err != nil {
		log.Println("Failed to unmarshal reCAPTCHA secret: ", err)
		return nil, err
	}

	return h, nil
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, loginPageHTML)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Stop processing if request is NOT POST
	if r.Method != "POST" {
		log.Println("Invalid HTTP method")
		http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		return
	}

	// Obtain username and password from POST body in request
	un := r.FormValue("username")
	pw := r.FormValue("password")

	// Check if the login credentials are valid
	if un != validUsername || pw != validPassword {
		log.Println("Invalid username or password")
		http.Redirect(w, r, "/failure", http.StatusFound)
		return
	}

	// Verify reCAPTCHA score by sending request to Google
	// Dummy value for now
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
- https://gobyexample.com/json

# Line Count
- Total:      108
- Reused:     0
- Written:    96
- Referenced: 12
*/
