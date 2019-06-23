package dropbox

import (
	"bytes"
	"log"
	"net/http"
)

func createRequest(accessToken string, method string, url string, requestBody []byte) *http.Request {
	// Create a Bearer string by appending string access token
	bearer := "Bearer " + accessToken

	// Create a new request using http
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	return req
}
