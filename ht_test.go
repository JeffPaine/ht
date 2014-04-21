package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetIPAddress(t *testing.T) {
	request, _ := http.NewRequest("GET", "/get", nil)

	remoteAddr := "1.2.3.4"
	request.RemoteAddr = remoteAddr
	ip := getIPAddress(request)
	if ip != remoteAddr {
		t.Errorf("Received incorrect IP address. Expected: %v, received: %v.", remoteAddr, ip)
	}

	headerIP := "5.6.7.8"
	request.Header.Set("X-Forwarded-For", headerIP)
	ip = getIPAddress(request)
	if ip != headerIP {
		t.Errorf("Failed to parse X-Forwarded-For properly. Expected: %v, received: %v", headerIP, ip)
	}
}

func TestFlatten(t *testing.T) {
	input := map[string][]string{
		"a": []string{"b"},
		"c": []string{"d"},
	}
	expected := map[string]string{
		"a": "b",
		"c": "d",
	}
	output := flatten(input)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v, received %v", expected, output)
	}
}

func TestIndexReturns200StatusCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	index(response, request)

	if response.Code != 200 {
		t.Errorf("Received unexpected status code of %v", response.Code)
	}
}

func TestIpReturns200StatusCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/ip", nil)
	response := httptest.NewRecorder()

	ip(response, request)

	if response.Code != 200 {
		t.Errorf("Received unexpected status code of %v", response.Code)
	}
}

func TestUserAgentReturns200StatusCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/user-agent", nil)
	response := httptest.NewRecorder()

	userAgent(response, request)

	if response.Code != 200 {
		t.Errorf("Received unexpected status code of %v", response.Code)
	}
}

func TestHeadersReturns200StatusCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/headers", nil)
	response := httptest.NewRecorder()

	headers(response, request)

	if response.Code != 200 {
		t.Errorf("Received unexpected status code of %v", response.Code)
	}
}

func TestGetReturns200StatusCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/get", nil)
	response := httptest.NewRecorder()

	get(response, request)

	if response.Code != 200 {
		t.Errorf("Received unexpected status code of %v", response.Code)
	}
}
