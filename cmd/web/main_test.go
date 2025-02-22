// Filename: cmd/web/main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(home)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("got %v, expected %v", status, http.StatusOK)
	}

	expected := "Hello from newsletter"
	got := rr.Body.String()
	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
