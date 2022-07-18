package recaptcha

import (
	"encoding/json"
	"fmt"
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

func Verify(secretKey, urToken, remoteIP string) (bool, float64, error) {
	// Prepare request body with secret-key and UR token
	reqBody := url.Values{
		"secret":   {secretKey},
		"response": {urToken},
		"remoteip": {remoteIP},
	}

	// Obtain reCAPTCHA score by sending request to Google
	res, err := http_client.SendPostRequest(reCAPTCHAVerifyURL, reqBody)
	if err != nil {
		newErr := fmt.Errorf("[recaptcha.Verifty()] Failed to get reCAPTCHA score: %v", err)
		log.Println(newErr.Error())
		return false, 0.0, newErr
	}
	log.Println("reCAPTCHA result: ", res)

	// Verify score result by parsing response JSON
	var scoreResult reCAPTCHAScoreStruct
	if err = json.Unmarshal([]byte(res), &scoreResult); err != nil {
		newErr := fmt.Errorf("[recaptcha.Verify()] Failed to unmarshal reCAPTCHA score: %v", err)
		log.Println(newErr.Error())
		return false, 0.0, newErr
	}

	return scoreResult.Success, scoreResult.Score, nil
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
