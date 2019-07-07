package dropbox

import (
	"encoding/json"
	"fmt"
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

func createRequestListFolder(path string, recursive bool) *requestListFolder {
	return &requestListFolder{
		Path:      path,
		Recursive: recursive,
	}
}

// ListContents lists out the contents of the Dropbox folder
func ListContents() {
	dropboxAccessToken := os.Getenv(DropboxAccessTokenEnv)
	url := "https://api.dropboxapi.com/2/files/list_folder"

	requestBody, err := json.Marshal(createRequestListFolder("", true))
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

// UploadFileResponse is the response that Dropbox returns after uploading a file
// This is not an exhaustive list of all the fields that is in the response
type UploadFileResponse struct {
	ClientModified string `json:"client_modified"`
	ContentHash    string `json:"content_hash"`
	ID             string `json:"id"`
	IsDownloadable bool   `json:"is_downloadable"`
	Name           string `json:"name"`
	PathDisplay    string `json:"path_display"`
	PathLower      string `json:"path_lower"`
	Rev            string `json:"rev"`
	ServerModified string `json:"server_modified"`
	Size           uint64 `json:"size"`
}

// DropboxAPIArg is the request header passed to Dropbox containing info related to the action
type DropboxAPIArg struct {
	Path           string `json:"path"`
	Mode           string `json:"mode"`
	Mute           bool   `json:"mute"`
	StrictConflict bool   `json:"strict_conflict"`
}

// "Dropbox-Api-Arg", "{\"path\": \"/testfileupload.txt\",\"mode\": \"overwrite\",	\"mute\": false,\"strict_conflict\": false}"
func createDropboxAPIArg(filePath string) DropboxAPIArg {
	return DropboxAPIArg{
		Path:           filePath,
		Mode:           "overwrite",
		Mute:           false,
		StrictConflict: false,
	}
}

// UploadFile uploads the file at filePath to Dropbox
func UploadFile(srcFilePath string, dstFilePath string) {
	dropboxAccessToken := os.Getenv(DropboxAccessTokenEnv)
	url := "https://content.dropboxapi.com/2/files/upload"

	f, err := os.Open(srcFilePath)
	if err != nil {
		log.Fatalf("Failed to open file to upload: %v", err)
	}
	defer f.Close()

	req, err := http.NewRequest("POST", url, f)
	if err != nil {
		log.Fatalf("Error creating new HTTP request: %v", err)
	}

	jsonDropboxAPIArg, err := json.Marshal(createDropboxAPIArg(dstFilePath))
	if err != nil {
		log.Fatalf("Error marshalling DropboxAPIArg into json: %v", err)
	}

	bearer := "Bearer " + dropboxAccessToken
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Dropbox-Api-Arg", fmt.Sprintf("%s", jsonDropboxAPIArg))
	req.Header.Add("Content-Type", "application/octet-stream")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	switch resp.StatusCode {
	case 200:
		var responseBody UploadFileResponse
		err = json.Unmarshal(body, &responseBody)
		log.Printf("Successfully uploaded file: %v", responseBody)
	case 400:
		log.Printf("HTTP [%v] Error uploading file, response body: %v", resp.StatusCode, string(body))
	default:
		log.Printf("Unhandled HTTP error [%v] uploading file, response body: %v", resp.StatusCode, string(body))
	}
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
