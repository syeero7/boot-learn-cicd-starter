package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	apiKey := "34re3fdwief23jdsjdfger/df"
	headers.Set("Authorization", fmt.Sprintf("ApiKey %s", apiKey))
	if key, err := GetAPIKey(headers); err != nil || key != apiKey {
		t.Fatalf("expected: %s, recived: %s", apiKey, key)
	}

	headers.Set("Authrization", "xdd")
	if _, err := GetAPIKey(headers); err == nil {
		t.Fatal("expected an error but recived none")
	}
}
