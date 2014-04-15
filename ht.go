package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// JSON response object
// http://nesv.blogspot.com/2012/09/super-easy-json-http-responses-in-go.html
type Response map[string]interface{}

func (r Response) String() string {
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
	} else {
		return r.RemoteAddr
	}
}

// Return the index page
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>HT</h1><p><strong>H</strong>TTP <strong>T</strong>esting</p>")
}

// Return the requesting host's IP address
func ip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, Response{"origin": getIPAddress(r)})
}

// Return the requesting host's User Agent, if provided
func userAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, Response{"user-agent": r.UserAgent()})
}

// Return the requesting host's headers
func headers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, Response{"headers": r.Header})
}

// Return GET data
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := r.URL.String()
	args := r.URL.Query()
	origin := getIPAddress(r)
	headers := r.Header
	resp := Response{"url": url, "args": args, "origin": origin, "headers": headers}
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
