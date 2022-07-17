package http

import (
	"log"
	"net/http"
)

func SendPostRequest(targetURL string) (string, error) {
	// Send HTTP POST request to target URL
	res, err := http.Post(targetURL, "", nil)
	if err != nil {
		log.Println("Failed to send POST request: ", err)
		return "", err
	}
	defer res.Body.Close()

	// Read response body
	var resBodyByte []byte
	if _, err := res.Body.Read(resBodyByte); err != nil {
		log.Println("Failed to read response body: ", err)
		return "", err
	}

	return string(resBodyByte), nil
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
