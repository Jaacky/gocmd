package dropbox

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	// DropboxAccessTokenEnv is the ENV var name for the Dropbox access token
	DropboxAccessTokenEnv = "DROPBOX_ACCESS_TOKEN"
)

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

// ListContents lists out the contents of the Dropbox folder
func ListContents() {
	dropboxAccessToken := os.Getenv(DropboxAccessTokenEnv)
	url := "https://api.dropboxapi.com/2/files/list_folder"

	requestBody, err := json.Marshal(createNewRequestListFolder("", true))
	if err != nil {
		log.Fatalln(err)
	}

	req := createRequest(dropboxAccessToken, "POST", url, requestBody)

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

	req := createRequest(dropboxAccessToken, "POST", url, requestBody)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))
}
