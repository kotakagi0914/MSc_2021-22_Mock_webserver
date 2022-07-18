package recaptcha

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/sheva0914/MSc_2021-22_Mock_webserver/pkg/http_client"
)

const (
	reCAPTCHAVerifyURL = "https://www.google.com/recaptcha/api/siteverify"
)

// Struct for reCAPTCHA score returned from Google
type reCAPTCHAScoreStruct struct {
	Success bool    `json:"success"`
	Score   float64 `json:"score"`
}

func Verify(secretKey, urToken, remoteIP string) (isSuccess bool, score float64, err error) {
	// Prepare request body with secret-key and UR token
	reqBody := url.Values{
		"secret":   {secretKey},
		"response": {urToken},
		"remoteip": {remoteIP},
	}

	// Obtain reCAPTCHA score by sending request to Google
	res, err := http_client.SendPostRequest(reCAPTCHAVerifyURL, reqBody)
	if err != nil {
		log.Println("Failed to get reCAPTCHA score: ", err)
		return
	}
	log.Println("reCAPTCHA result: ", res)

	// Verify score result by parsing response JSON
	var scoreResult reCAPTCHAScoreStruct
	if err = json.Unmarshal([]byte(res), &scoreResult); err != nil {
		log.Println("Failed to unmarshal reCAPTCHA score: ", err)
		return
	}

	isSuccess = scoreResult.Success
	score = scoreResult.Score

	return
}

/*
# Reference
- https://pkg.go.dev/net/http
- https://zetcode.com/golang/getpostrequest/

# Line Count
- Total:      48
- Reused:     0
- Written:    42
- Referenced: 6
*/
