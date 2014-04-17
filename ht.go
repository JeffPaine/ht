package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// JSONResponse object
// http://nesv.blogspot.com/2012/09/super-easy-json-http-responses-in-go.html
type JSONResponse map[string]interface{}

func (r JSONResponse) String() string {
	b, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

// Return the IP address of the given *http.Request
func getIPAddress(r *http.Request) string {
	if val, ok := r.Header["X-Forwarded-For"]; ok {
		return val[0]
	}
	return r.RemoteAddr
}

// Return the index page
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// Return the requesting host's IP address
func ip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, JSONResponse{"origin": getIPAddress(r)})
}

// Return the requesting host's User Agent, if provided
func userAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, JSONResponse{"user-agent": r.UserAgent()})
}

// Return the requesting host's headers
func headers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, JSONResponse{"headers": r.Header})
}

// Return GET data
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := JSONResponse{
		"url":     r.URL.String(),
		"args":    r.URL.Query(),
		"origin":  getIPAddress(r),
		"headers": r.Header,
	}
	fmt.Fprint(w, resp)
}

func main() {
	// URLs
	http.HandleFunc("/", index)
	http.HandleFunc("/ip", ip)
	http.HandleFunc("/user-agent", userAgent)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/get", get)

	// Set our PORT to listen on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, nil)
}
