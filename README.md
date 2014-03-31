# ht

A simple endpoint for HTTP Tests.

An incomplete, alpha quality Go port of, and homage to, Kenneth Reitz's [httpbin](https://github.com/kennethreitz/httpbin).

## Endpoints

* `/` Return the homepage
* `/ip` Returns the requesting IP address
* `/user-agent` Returns the user agent of the requestor
* `/headers` Returns the request headers

## Why

Sometimes you just need to test your code against an endpoint on the open internet. `ht` allows you to test your code plus gives you some handy endpoints for other tasks. The only limit is your creativeness.
