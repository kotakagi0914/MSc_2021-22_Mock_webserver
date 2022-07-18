package http_client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func SendPostRequest(targetURL string, reqBody url.Values) (string, error) {
	// Send HTTP POST request to target URL
	res, err := http.PostForm(targetURL, reqBody)
	if err != nil {
		newErr := fmt.Errorf("[http_client.SendPostRequest] Failed to send POST request: %v", err)
		log.Println(newErr.Error())
		return "", newErr
	}
	defer res.Body.Close()

	// Read response body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		newErr := fmt.Errorf("[http_client.SendPostRequest] Failed to read response body: %v", err)
		log.Println(newErr.Error())
		return "", newErr
	}

	return string(resBody), nil
}

/*
# Reference
- https://pkg.go.dev/net/http

# Line Count
- Total:      30
- Reused:     0
- Written:    30
- Referenced: 0
*/
