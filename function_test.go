package helloworld

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloHTTP tests the HelloHTTP function with various inputs.
func TestHelloHTTP(t *testing.T) {
	tests := []struct {
		name       string
		body       interface{} // JSON body to send (nil for no body)
		wantStatus int         // Expected HTTP status code
		wantBody   string      // Expected response body
	}{
		{
			name:       "no body",
			body:       nil,
			wantStatus: http.StatusOK,
			wantBody:   "Hello, World!",
		},
		{
			name:       "empty name",
			body:       struct{ Name string }{Name: ""},
			wantStatus: http.StatusOK,
			wantBody:   "Hello, World!",
		},
		{
			name:       "valid name",
			body:       struct{ Name string }{Name: "Alice"},
			wantStatus: http.StatusOK,
			wantBody:   "Hello, Alice!",
		},
		{
			name:       "name with HTML",
			body:       struct{ Name string }{Name: "<script>alert('xss')</script>"},
			wantStatus: http.StatusOK,
			wantBody:   "Hello, &lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;!",
		},
		{
			name:       "invalid JSON",
			body:       "not json",
			wantStatus: http.StatusOK,
			wantBody:   "Hello, World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request
			var req *http.Request
			if tt.body == nil {
				req = httptest.NewRequest(http.MethodPost, "/", nil)
			} else {
				var bodyBytes []byte
				if str, ok := tt.body.(string); ok {
					bodyBytes = []byte(str) // For invalid JSON case
				} else {
					var err error
					bodyBytes, err = json.Marshal(tt.body)
					if err != nil {
						t.Fatalf("Failed to marshal JSON: %v", err)
					}
				}
				req = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bodyBytes))
				req.Header.Set("Content-Type", "application/json")
			}

			// Create a response recorder
			w := httptest.NewRecorder()

			// Call the handler
			HelloHTTP(w, req)

			// Check status code
			if got := w.Code; got != tt.wantStatus {
				t.Errorf("HelloHTTP() status = %v, want %v", got, tt.wantStatus)
			}

			// Check response body
			if got := w.Body.String(); got != tt.wantBody {
				t.Errorf("HelloHTTP() body = %q, want %q", got, tt.wantBody)
			}
		})
	}
}