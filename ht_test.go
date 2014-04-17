package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexReturns200StatusCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	index(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Received unexpected status code of %v", response.Code)
	}
}

func TestIpReturns200StatusCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/ip", nil)
	response := httptest.NewRecorder()

	ip(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Received unexpected status code of %v", response.Code)
	}
}

func TestUserAgentReturns200StatusCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/user-agent", nil)
	response := httptest.NewRecorder()

	userAgent(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Received unexpected status code of %v", response.Code)
	}
}

func TestHeadersReturns200StatusCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/headers", nil)
	response := httptest.NewRecorder()

	headers(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Received unexpected status code of %v", response.Code)
	}
}

func TestGetReturns200StatusCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/get", nil)
	response := httptest.NewRecorder()

	get(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Received unexpected status code of %v", response.Code)
	}
}
