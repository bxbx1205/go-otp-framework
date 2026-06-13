package otp

import (
	"testing"
)

func TestCreateAPIKey_Embedded(t *testing.T) {
	server := New()

	apiKey, err := server.CreateAPIKey("")
	if err != nil {
		t.Fatalf("Expected successful API key generation with empty userID, got error: %v", err)
	}
	if apiKey == "" {
		t.Fatalf("Expected non-empty API key")
	}

	_, err = server.CreateAPIKey("invalid-id")
	if err == nil {
		t.Fatalf("Expected error for invalid user id, got nil")
	}
	if err.Error() != "invalid user id" {
		t.Fatalf("Expected error 'invalid user id', got '%v'", err.Error())
	}
}
