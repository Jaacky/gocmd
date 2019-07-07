package dropbox

import (
	"testing"
)

func TestCreateRequestListFolder(t *testing.T) {
	expectedPath := "/test/path/file.txt"
	expectedRecursive := true

	request := createRequestListFolder(expectedPath, expectedRecursive)

	if request.Path != expectedPath {
		t.Errorf("Path doesn't match. Expected: %v, got: %v\n", expectedPath, request.Path)
	}

	if request.Recursive != expectedRecursive {
		t.Errorf("Recursive doesn't match. Expected: %v, got: %v\n", expectedRecursive, request.Recursive)
	}
}
func TestCreateDropboxAPIArg(t *testing.T) {
	expectedPath := "/test/path/file.txt"
	expectedMode := "overwrite"
	expectedMute := false
	expectedStrictConflict := false

	dropboxAPIArg := createDropboxAPIArg(expectedPath)

	if dropboxAPIArg.Path != expectedPath {
		t.Errorf("Path doesn't match. Expected: %v, got: %v\n", expectedPath, dropboxAPIArg.Path)
	}

	if dropboxAPIArg.Mode != expectedMode {
		t.Errorf("Mode doesn't match. Expected: %v, got: %v\n", expectedMode, dropboxAPIArg.Mode)
	}

	if dropboxAPIArg.Mute != expectedMute {
		t.Errorf("Mute doesn't match. Expected: %v, got: %v\n", expectedMute, dropboxAPIArg.Mute)
	}

	if dropboxAPIArg.StrictConflict != expectedStrictConflict {
		t.Errorf("StrictConflict doesn't match. Expected: %v, got: %v\n", expectedStrictConflict, dropboxAPIArg.StrictConflict)
	}
}
