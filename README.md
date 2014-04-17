# ht

A server with endpoints for HTTP Tests, written in Go.

An alpha quality homage to Kenneth Reitz's awesome [httpbin](https://github.com/kennethreitz/httpbin).

## Example

```bash
$ curl http://ht-go.herokuapp.com/user-agent
{
    "user-agent": "curl/7.24.0 (x86_64-apple-darwin12.0) libcurl/7.24.0 OpenSSL/0.9.8y zlib/1.2.5"
}
```

## Endpoints

* [`/`](http://ht-go.herokuapp.com/) Return the `ht` homepage
* [`/ip`](http://ht-go.herokuapp.com/ip) Returns the requesting IP address
* [`/user-agent`](http://ht-go.herokuapp.com/user-agent) Returns the requesting user-agent
* [`/headers`](http://ht-go.herokuapp.com/headers) Returns the request headers
* [`/get`](http://ht-go.herokuapp.com/get) Returns GET request data

## Why

Sometimes you just need to test your code against an endpoint on the open internet.
