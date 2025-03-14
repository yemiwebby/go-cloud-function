package helloworld

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHTTP_WithName(t *testing.T) {
	// Create a JSON payload with a "name" field.
	payload := map[string]string{"name": "TestUser"}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	// Create a new POST request with the JSON payload.
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	w := httptest.NewRecorder()

	// Call the function.
	HelloHTTP(w, req)

	// Check that the response matches the expected output.
	expected := "Hello, TestUser!"
	if got := w.Body.String(); got != expected {
		t.Errorf("Expected response %q, got %q", expected, got)
	}
}

func TestHelloHTTP_WithoutName(t *testing.T) {
	// Create a request with an empty JSON object.
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(`{}`)))
	w := httptest.NewRecorder()

	// Call the function.
	HelloHTTP(w, req)

	// When no name is provided, expect "Hello, World!".
	expected := "Hello, World!"
	if got := w.Body.String(); got != expected {
		t.Errorf("Expected response %q, got %q", expected, got)
	}
}

func TestHelloHTTP_InvalidJSON(t *testing.T) {
	// Create a request with invalid JSON.
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte("invalid json")))
	w := httptest.NewRecorder()

	// Call the function.
	HelloHTTP(w, req)

	// With invalid JSON, the function should default to "Hello, World!".
	expected := "Hello, World!"
	if got := w.Body.String(); got != expected {
		t.Errorf("Expected response %q, got %q", expected, got)
	}
}
