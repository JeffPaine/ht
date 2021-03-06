package main

import (
	"encoding/json"
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// jsonResponse objects allow you to easily create arbitrary JSON objects.
// http://nesv.blogspot.com/2012/09/super-easy-json-http-responses-in-go.html
type jsonResponse map[string]interface{}

// String allows you to auto-marshal jsonResponse objects into properly formatted JSON
// when called.
func (r jsonResponse) String() string {
	b, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

// getIPAddress returns the IP address of the given *http.Request.
func getIPAddress(r *http.Request) string {
	if val, ok := r.Header["X-Forwarded-For"]; ok {
		return val[0]
	}
	return r.RemoteAddr
}

// flatten takes a map[string][]string and flattens it into a nice map[string]string.
// This saves us from having JSON values that are unnecessarily nested.
// If there are multiple entries in the original []string, they will be joined
// into a single string like "a,b" as per HTTP headers spec for multiple entries.
// Before:
//    "args": {
//            "a": [
//                "b"
//            ],
//            "c": [
//                "d", "e"
//            ]
//        }
// After:
//    "args": {
//            "a": "b",
//            "c": "d,e"
//        }
func flatten(m map[string][]string) map[string]string {
	newHeader := make(map[string]string)
	for k, v := range m {
		newHeader[k] = strings.Join(v, ",")
	}
	return newHeader
}

// index parses the README.md file and returns its contents in HTML form.
func index(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadFile("README.md")
	if err != nil {
		fmt.Fprint(w, "Error")
	}
	output := blackfriday.MarkdownCommon(input)
	fmt.Fprint(w, string(output))
}

// ip returns the requesting host's IP address.
func ip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, jsonResponse{"origin": getIPAddress(r)})
}

// userAgent returns the requesting host's User Agent, if provided.
func userAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, jsonResponse{"user-agent": r.UserAgent()})
}

// headers returns the requesting host's request headers.
func headers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, jsonResponse{"headers": flatten(r.Header)})
}

// get returns GET request data.
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := jsonResponse{
		"url":     r.URL.String(),
		"args":    flatten(r.URL.Query()),
		"origin":  getIPAddress(r),
		"headers": flatten(r.Header),
	}
	fmt.Fprint(w, resp)
}

// robots returns robots.txt compatible data.
func robots(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "User-agent: *\nDisallow: /deny")
}

// deny returns a robots.txt denied resource.
func deny(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This URL is denied in robots.txt")
}

func main() {
	// URLs
	http.HandleFunc("/", index)
	http.HandleFunc("/ip", ip)
	http.HandleFunc("/user-agent", userAgent)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/get", get)
	http.HandleFunc("/robots.txt", robots)
	http.HandleFunc("/deny", deny)

	// Set our PORT to listen on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, nil)
}
