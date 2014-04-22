package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestJsonResponseString(t *testing.T) {
	j := jsonResponse{"a": "b"}
	expected := `{
    "a": "b"
}`
	if j.String() != expected {
		t.Errorf("Expected: %v, Received: %v", expected, j.String())
	}
}

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

func TestHandlersReturnExpectedStatusCodes(t *testing.T) {
	type handler struct {
		f      func(http.ResponseWriter, *http.Request)
		method string
		path   string
		status int
	}

	handlers := []handler{
		handler{index, "GET", "/", 200},
		handler{ip, "GET", "/ip", 200},
		handler{userAgent, "GET", "/user-agent", 200},
		handler{headers, "GET", "/headers", 200},
		handler{get, "GET", "/get", 200},
	}

	for _, h := range handlers {
		request, _ := http.NewRequest(h.method, h.path, nil)
		response := httptest.NewRecorder()

		h.f(response, request)

		if response.Code != h.status {
			t.Errorf("Expected status code %v, received %v", h.status, response.Code)
		}
	}
}
