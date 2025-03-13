package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yemiwebby/go-cloud-function/handler"
)

func TestHelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create HTTP request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.HelloWorld)

	handler.ServeHTTP(rr, req)

	expected := "Hello, World from Go Cloud Function!\n"
	if rr.Body.String() != expected {
		t.Errorf("Expected %q but got %q", expected, rr.Body.String())
	}
}
