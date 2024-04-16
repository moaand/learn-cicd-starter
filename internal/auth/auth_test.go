package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	header := http.Header{}

	header.Add("Authorization", "ApiKey hej")

	_, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf(`%v`, err)
	}
}
