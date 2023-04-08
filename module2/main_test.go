package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	os.Setenv("VERSION", "1.0.0")

	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-My-Header", "test-value")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFunc)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Header().Get("X-My-Header") != "test-value" {
		t.Errorf("handler didn't return expected header value: got %v want %v", rr.Header().Get("X-My-Header"), "test-value")
	}

	if rr.Header().Get("VERSION") != "1.0.0" {
		t.Errorf("handler didn't return expected VERSION header value: got %v want %v", rr.Header().Get("VERSION"), "1.0.0")
	}
}
