#!/bin/bash
JSON=$(cat <<EOF
{
   "routes": [
      {
         "path": "/products",
         "host": "http://localhost:8080"
      },
      {
         "path": "/cart",
         "host": "http://localhost:8081"
      },
      {
         "path": "/orders",
         "host": "http://localhost:8082"
      },
   ]
}
EOF
)

echo $JSON
export VCAP_SERVICES=$JSON
export VCAP_APPLICATION='{"host":"localhost", "port":3000, "uris": "localhost", "name": "store proxy" }'
