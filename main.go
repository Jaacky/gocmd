package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	DropboxAccessTokenEnv = "DROPBOX_ACCESS_TOKEN"
)

func main() {
	listContents()
}

type requestListFolder struct {
	Path      string `json:"path"`
	Recursive bool   `json:"recursive"`
}

func createNewRequestListFolder(path string, recursive bool) *requestListFolder {
	return &requestListFolder{
		Path:      path,
		Recursive: recursive,
	}
}

func listContents() {
	dropboxAccessToken := os.Getenv(DropboxAccessTokenEnv)
	url := "https://api.dropboxapi.com/2/files/list_folder"

	requestBody, err := json.Marshal(createNewRequestListFolder("", true))

	if err != nil {
		log.Fatalln(err)
	}

	// Create a Bearer string by appending string access token
	bearer := "Bearer " + dropboxAccessToken

	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))
}

func createFolder() {
	dropboxAccessToken := os.Getenv(DropboxAccessTokenEnv)
	url := "https://api.dropboxapi.com/2/files/create_folder_v2"

	requestBody, err := json.Marshal(map[string]string{
		"path": "/Test/Inside",
	})

	if err != nil {
		log.Fatalln(err)
	}

	// Create a Bearer string by appending string access token
	bearer := "Bearer " + dropboxAccessToken

	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))
}
