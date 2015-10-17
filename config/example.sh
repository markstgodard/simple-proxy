#!/bin/bash
JSON=$(cat <<EOF
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
EOF
)

echo $JSON
export ROUTES=$JSON
