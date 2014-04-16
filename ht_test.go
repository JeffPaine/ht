package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	)

func TestIndexReturns200StatusCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	index(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Received unexpected status code of %v", response.Code)
	}
}

func TestIp(t *testing.T) {
	request, _ := http.NewRequest("GET", "/ip", nil)
	response := httptest.NewRecorder()

	ip(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Received unexpected status code of %v", response.Code)
	}
}
