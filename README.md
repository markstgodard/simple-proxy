# Simple proxy
A simple context path based reverse proxy for micro services

## Installation
    go get github.com/markstgodard/simple-proxy

## Config
Set an environment variable 'ROUTES' containing path and hosts.

For example:
```json
[
    {
       "Path": "/products",
       "Host": "http://localhost:8080"
    },
    {
       "Path": "/cart",
       "Host": "http://localhost:8081"
    },
    {
       "Path": "/orders",
       "Host": "http://localhost:8082"
    },
    {
       "Path": "/store",
       "Host": "http://localhost:8000"
    }
]
```

## Running
    simple-proxy
