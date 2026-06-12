package framework

import (
	"testing"
)

func TestCreateAPIKey_Embedded(t *testing.T) {
	// Create a server without any MongoDB or Redis connections
	server := New()

	// 1. If userID is empty: automatically create/use a default embedded user, generate API key successfully
	apiKey, err := server.CreateAPIKey("")
	if err != nil {
		t.Fatalf("Expected successful API key generation with empty userID, got error: %v", err)
	}
	if apiKey == "" {
		t.Fatalf("Expected non-empty API key")
	}

	// 2. If userID is not a valid ObjectID: Return a descriptive error "invalid user id"
	_, err = server.CreateAPIKey("invalid-id")
	if err == nil {
		t.Fatalf("Expected error for invalid user id, got nil")
	}
	if err.Error() != "invalid user id" {
		t.Fatalf("Expected error 'invalid user id', got '%v'", err.Error())
	}
}
