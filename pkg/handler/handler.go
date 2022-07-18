package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/sheva0914/MSc_2021-22_Mock_webserver/pkg/recaptcha"
)

const (
	// File paths to template HTMLs
	loginPageTemplateFilePath  = "./web/template/login-page.html"
	resultPageTemplateFilePath = "./web/template/result-page.html"

	// Credentials
	secretFile    = "./.secret"
	validUsername = "admin"
	validPassword = "password"

	// Query strings for result page
	loginSuccessQueryStr     = "loginSuccess"
	loginErrorQueryStr       = "loginError"
	reCAPTCHASuccessQueryStr = "reCAPTCHASuccess"
	reCAPTCHAScoreQueryStr   = "reCAPTCHAScore"
)

// Struct for reCAPTCHA secret containing site-key and secret-key
type reCAPTCHASecretStruct struct {
	SiteKey   string `json:"site-key"`
	SecretKey string `json:"secret-key"`
}

// Struct for reCAPTCHA result containing success flag, score, and error description if applicable
type reCAPTCHAResultStruct struct {
	LoginSuccess     bool
	LoginError       string
	ReCAPTCHASuccess bool
	ReCAPTCHAScore   float64
}

var (
	recaptchaSecret    reCAPTCHASecretStruct
	loginPageHTML      string
	resultPageTemplate *template.Template
)

func Init() (*http.ServeMux, error) {
	// Initialise HTTP server handler
	h := http.NewServeMux()
	h.HandleFunc("/", MainPageHandler)
	h.HandleFunc("/login", LoginHandler)
	h.HandleFunc("/login-result", LoginResultHandler)
	h.HandleFunc("/success", SuccessPageHandler)
	h.HandleFunc("/failure", FailurePageHandler)

	// Load login page template
	loginPageTemplateByte, err := os.ReadFile(loginPageTemplateFilePath)
	if err != nil {
		log.Println("Failed to load login page template: ", err)
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
		log.Println("Failed to create new template instance for login page: ", err)
		return nil, err
	}

	// Set site-key into login page HTML
	loginPageByte := new(bytes.Buffer)
	if err := loginPageTemp.Execute(loginPageByte, recaptchaSecret); err != nil {
		log.Println("Failed to inject values into login page template: ", err)
		return nil, err
	}
	loginPageHTML = loginPageByte.String()

	// Load login result page template
	resultPageTemplateByte, err := os.ReadFile(resultPageTemplateFilePath)
	if err != nil {
		log.Println("Failed to load result page template", err)
		return nil, err
	}

	// Create template instance for result page
	resultPageTemplate, err = template.New("resultpage").Parse(string(resultPageTemplateByte))
	if err != nil {
		log.Println("Failed to create new template instance for result page: ", err)
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

	// Obtain username, password, user response token and client IP address from request
	un := r.FormValue("username")
	pw := r.FormValue("password")
	urToken := r.FormValue("ur-token")
	remoteIP := r.RemoteAddr

	// Verify the user with reCAPTCHA
	isSuccess, score, err := recaptcha.Verify(recaptchaSecret.SecretKey, urToken, remoteIP)
	if err != nil {
		log.Println("Failed to get reCAPTCHA verification result: ", err)
		http.Redirect(w, r, "/failure", http.StatusFound)
		return
	}

	// Check reCAPTCHA verification result
	if !isSuccess {
		log.Println("reCAPTCHA verification failed")
		http.Redirect(w, r, "/failure", http.StatusFound)
		return
	}

	// Check if the login credentials are valid
	if un != validUsername || pw != validPassword {
		loginErrStr = "Invalid username or password"
		log.Println(loginErrStr)
		http.Redirect(
			w,
			r,
			makeQueryString(isLoginSuccess, isReCAPTCHASuccess, loginErrStr, reCAPTCHAScore),
			http.StatusFound,
		)
		// http.Redirect(w, r, "/failure", http.StatusFound)
		return
	}
	isLoginSuccess = true

	http.Redirect(w, r, fmt.Sprintf("/success?%s=%f", reCAPTCHAScoreQueryStr, reCAPTCHAScore), http.StatusFound)
}

func LoginResultHandler(w http.ResponseWriter, r *http.Request) {
	loginSuccessStr := r.URL.Query().Get(loginSuccessQueryStr)
	var isLoginSuccess bool
	if loginSuccessStr == "true" {
		isLoginSuccess = true
	}

	loginErrorStr := r.URL.Query().Get(loginErrorQueryStr)

	reCAPTCHASuccessStr := r.URL.Query().Get(reCAPTCHASuccessQueryStr)
	var isReCAPTCHASuccess bool
	if reCAPTCHASuccessStr == "true" {
		isReCAPTCHASuccess = true
	}

	reCAPTCHAScoreStr := r.URL.Query().Get(reCAPTCHAScoreQueryStr)
	log.Println("score: ", reCAPTCHAScoreStr)
	reCAPTCHAScore, err := strconv.ParseFloat(reCAPTCHAScoreStr, 64)
	if err != nil {
		log.Println("Invalid score type: ", err)
		http.Error(w, "Invalid score type", http.StatusBadRequest)
		return
	}

	reCAPTCHAResult := reCAPTCHAResultStruct{
		LoginSuccess:     isLoginSuccess,
		LoginError:       loginErrorStr,
		ReCAPTCHASuccess: isReCAPTCHASuccess,
		ReCAPTCHAScore:   reCAPTCHAScore,
	}

	// Set each params into result page HTML
	resultPageByte := new(bytes.Buffer)
	if err := resultPageTemplate.Execute(resultPageByte, reCAPTCHAResult); err != nil {
		log.Println("Failed to inject values into result page template: ", err)
		http.Error(w, "Failed to generate result page", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, resultPageByte.String())
}

func SuccessPageHandler(w http.ResponseWriter, r *http.Request) {
	score := r.URL.Query().Get(reCAPTCHAScoreQueryStr)
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
- Total:      139
- Reused:     0
- Written:    120
- Referenced: 19
*/
