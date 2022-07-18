package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func SendPostRequest(targetURL string, reqBody url.Values) (string, error) {
	// Send HTTP POST request to target URL
	res, err := http.PostForm(targetURL, reqBody)
	if err != nil {
		log.Println("Failed to send POST request: ", err)
		return "", err
	}
	defer res.Body.Close()

	// Read response body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Failed to read response body: ", err)
		return "", err
	}

	return string(resBody), nil
}

/*
# Reference
- https://pkg.go.dev/net/http

# Line Count
- Total:      25
- Reused:     0
- Written:    25
- Referenced: 0
*/
