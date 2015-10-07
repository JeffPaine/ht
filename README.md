# ht

Endpoints for HTTP Tests, written in Go.

An homage to Kenneth Reitz's excellent [httpbin](https://github.com/kennethreitz/httpbin).

## Quick Start

```
$ git clone https://github.com/JeffPaine/ht.git
$ cd ht/
$ PORT=8000  # set port we want ht to listen on (default: 8000)
$ docker build -t ht .
$ docker run -it --rm -p $PORT:$PORT --name ht-running ht
```

Then, from another shell session, run

```
$ curl "localhost:$8000/user-agent"  # see Endpoints for supported endpoints
{
    "user-agent": "curl/7.35.0"
}
```

## Endpoints

Endpoint | Description
--- | ---
`/` | `ht` homepage
`/ip` | Requesting IP address
`/user-agent` | Requesting user-agent
`/headers` | Request headers
`/get` | GET request data
`/robots.txt` | robots.txt compatible data
`/deny` | robots.txt denied resource

## Why

Folks will hopefully find these endpoints useful for testing code that
retrieves data over the network. Just fire up your own instance of `ht` and you
can test your code against it without having lean on some random site or
endpoints.

It's written in Go as a fun experiment and to take advantage of Go's excellent
performance (with 250 concurrent connections, response times averaged 16ms!).

## Source

[github.com/JeffPaine/ht](https://github.com/JeffPaine/ht)
