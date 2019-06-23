package dropbox

import (
	"bytes"
	"log"
	"net/http"
	"os"
)

const (
	// DropboxAccessTokenEnv is the ENV var name for the Dropbox access token
	DropboxAccessTokenEnv = "DROPBOX_ACCESS_TOKEN"
)

func createRequest(method string, url string, requestBody []byte) *http.Request {
	dropboxAccessToken := os.Getenv(DropboxAccessTokenEnv)

	// Create a Bearer string by appending string access token
	bearer := "Bearer " + dropboxAccessToken

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
