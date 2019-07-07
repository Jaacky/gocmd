package main

import (
	"os"
	"strings"

	"github.com/Jaacky/gocmd/dropbox"
)

const (
	keyfileEnv = "KEYFILE"
)

func main() {
	args := os.Args

	if len(args) >= 2 && os.Args[1] == "sync" {
		keyfile := os.Getenv(keyfileEnv)

		split := strings.Split(keyfile, "/")
		fileName := split[len(split)-1]

		dropbox.UploadFile(keyfile, "/"+fileName)
	}
}
