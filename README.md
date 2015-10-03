# ht

Endpoints for HTTP Tests, written in Go.

An homage to Kenneth Reitz's excellent [httpbin](https://github.com/kennethreitz/httpbin).

## Download

#### Command line

```
$ git clone https://github.com/JeffPaine/ht.git
```

#### Web browser
[ht download](https://github.com/JeffPaine/ht/archive/master.zip)

## Usage

```
$ cd ht/
# Default port is 8000, but you can use any you like / have permissions for.
$ PORT=8000
$ docker build -t ht .
$ docker run -it --rm -p $PORT:$PORT --name ht-running ht
# Try it out.
$ curl "localhost:$PORT/user-agent"
{
    "user-agent": "curl/7.35.0"
}
```

## Endpoints

Endpoint | Description
--- | ---
[`/`](/) | `ht` homepage
[`/ip`](/ip) | Requesting IP address
[`/user-agent`](/user-agent) | Requesting user-agent
[`/headers`](/headers) | Request headers
[`/get`](/get) | GET request data
[`/robots.txt`](/robots.txt) | robots.txt compatible data
[`/deny`](/deny) | robots.txt denied resource

## Why

Folks will hopefully find these endpoints useful for testing code that
retrieves data over the network. Just fire up your own instance of `ht` and you
can test your code against it without having lean on some random site or
endpoints.

It's written in Go as a fun experiment and to take advantage of Go's excellent
performance (with 250 concurrent connections, response times averaged 16ms!).

## Source

[github.com/JeffPaine/ht](https://github.com/JeffPaine/ht)
