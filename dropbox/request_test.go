package dropbox

import (
	"encoding/json"
	"testing"
)

func TestCreateRequest(t *testing.T) {
	testToken := "testToken"
	testMethod := "testMethod"
	testURL := "testURL"
	testRequestBody := createRequestListFolder("what/is/going/on", true)
	jsonRequestBody, err := json.Marshal(testRequestBody)
	if err != nil {
		t.Errorf("Error creating JSON request body: %s", err)
	}
	req := createRequest(testToken, testMethod, testURL, jsonRequestBody)

	expectedAuthorization := "Bearer " + testToken
	expectedContentType := "application/json"

	authorization := req.Header.Get("Authorization")
	contentType := req.Header.Get("Content-Type")

	d := json.NewDecoder(req.Body)
	requestBody := requestListFolder{}
	err = d.Decode(&requestBody)
	if err != nil {
		t.Errorf("Decoding request body failed: %s", err)
	}

	if authorization != expectedAuthorization {
		t.Errorf("Authorization doesn't match, expected: %v, got: %v", expectedAuthorization, authorization)
	}
	if contentType != expectedContentType {
		t.Errorf("Content-Type doesn't match, expected: %v, got: %v", expectedContentType, contentType)
	}
	if *testRequestBody != requestBody {
		t.Errorf("Request body doesn't match, expected: %v, got: %v", *testRequestBody, requestBody)
	}

	testFalseRequestBody := createRequestListFolder("what/is/going/on/here", true)
	// SANITY CHECK: Wrapping comparison in not to check if a random request body will equal expected
	if !(*testFalseRequestBody != requestBody) {
		t.Errorf("Unequal request body match, should not equal: %v, but got: %v", *testFalseRequestBody, requestBody)
	}
}
