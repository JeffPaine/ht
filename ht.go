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

func (r Response) String() (s string) {
	b, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

// Return the index page
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>HT</h1><p><strong>H</strong>TTP <strong>T</strong>esting</p>")
}

// Return the requesting host's IP address
func ip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, Response{"origin": r.RemoteAddr})
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

func main() {
	// URLs
	http.HandleFunc("/", index)
	http.HandleFunc("/ip", ip)
	http.HandleFunc("/user-agent", userAgent)
	http.HandleFunc("/headers", headers)

	// Set our PORT to listen on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, nil)
}
