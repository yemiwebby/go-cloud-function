package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create HTTP request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloWorld)

	handler.ServeHTTP(rr, req)

	expected := "Hello, World from Go Cloud Function!\n"
	if rr.Body.String() != expected {
		t.Errorf("Expected %q but got %q", expected, rr.Body.String())
	}
}
