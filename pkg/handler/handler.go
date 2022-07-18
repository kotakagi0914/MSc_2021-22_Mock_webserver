package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/sheva0914/MSc_2021-22_Mock_webserver/pkg/http_client"
)

const (
	templateHTMLFile    = "./web/template/login-page.html"
	secretFile          = "./.secret"
	validUsername       = "admin"
	validPassword       = "password"
	reCAPTCHAVerifyURL  = "https://www.google.com/recaptcha/api/siteverify"
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
	loginPageTemplateByte, err := os.ReadFile(templateHTMLFile)
	if err != nil {
		log.Println("Failed to load login page template", err)
		return nil, err
	}

	// Load site-key and secret-key from `.secret` file
	secretJsonByte, err := os.ReadFile(secretFile)
	if err != nil {
		log.Println("Failed to read reCAPTCHA secret file: ", err)
		return nil, err
	}

	// Unmarshal JSON byte into reCAPTCHASecretStruct type
	if err := json.Unmarshal(secretJsonByte, &recaptchaSecret); err != nil {
		log.Println("Failed to unmarshal reCAPTCHA secret: ", err)
		return nil, err
	}

	// Create template instance for login page
	loginPageTemp, err := template.New("loginpage").Parse(string(loginPageTemplateByte))
	if err != nil {
		log.Println("Failed to create new template instance: ", err)
		return nil, err
	}

	// Set site-key into login page HTML
	loginPageByte := new(bytes.Buffer)
	if err := loginPageTemp.Execute(loginPageByte, recaptchaSecret); err != nil {
		log.Println("Failed to inject values into template: ", err)
		return nil, err
	}
	loginPageHTML = loginPageByte.String()

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

	// Obtain username, password and user response token from POST body in request
	un := r.FormValue("username")
	pw := r.FormValue("password")
	urToken := r.FormValue("ur-token")

	// Prepare request body with secret-key and UR token
	reqBody := url.Values{
		"secret":   {recaptchaSecret.SecretKey},
		"response": {urToken},
		// "remoteip": {"1.1.1.1"},
	}

	// Obtain reCAPTCHA score by sending request to Google
	res, err := http_client.SendPostRequest(reCAPTCHAVerifyURL, reqBody)
	if err != nil {
		log.Println("Failed to get reCAPTCHA score: ", err)
		http.Error(w, "Failed to get reCAPTCHA score", http.StatusInternalServerError)
		return
	}
	log.Println(res)

	// Dummy score
	score := "0.5"

	// Check if the login credentials are valid
	if un != validUsername || pw != validPassword {
		log.Println("Invalid username or password")
		http.Redirect(w, r, "/failure", http.StatusFound)
		return
	}

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
- https://pkg.go.dev/html/template
- https://stackoverflow.com/questions/13765797/the-best-way-to-get-a-string-from-a-writer

# Line Count
- Total:      126
- Reused:     0
- Written:    107
- Referenced: 19
*/
