package handler

import (
	"fmt"
	"net/http"
)

// HelloWorld is the Cloud Function handler.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World from Go Cloud Function!")
}
