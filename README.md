# ht

Endpoints for HTTP Tests, written in Go.

An homage to Kenneth Reitz's excellent [httpbin](https://github.com/kennethreitz/httpbin).

## Example

```bash
$ curl http://ht-go.herokuapp.com/user-agent
{
    "user-agent": "curl/7.24.0 (x86_64-apple-darwin12.0) libcurl/7.24.0 OpenSSL/0.9.8y zlib/1.2.5"
}
```

## Endpoints

Endpoint | Description
--- | ---
[`/`](http://ht-go.herokuapp.com/) | `ht` homepage
[`/ip`](http://ht-go.herokuapp.com/ip) | Requesting IP address
[`/user-agent`](http://ht-go.herokuapp.com/user-agent) | Requesting user-agent
[`/headers`](http://ht-go.herokuapp.com/headers) | Request headers
[`/get`](http://ht-go.herokuapp.com/get) | GET request data

## Why

Folks will hopefully find these endpoints useful for testing code that retrieves data from the internet. It's written in Go as a fun experiment and to take advantage of Go's excellent performance (with 250 concurrent connections, response times averaged 16ms!). So, feel free to take advantage of the endpoints located at [http://ht-go.herokuapp.com/](http://ht-go.herokuapp.com/).

## Source

[github.com/JeffPaine/ht](https://github.com/JeffPaine/ht)
