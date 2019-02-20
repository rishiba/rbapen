#! /bin/bash
# urlencoded
set -x

# curl -X POST \
  # http://localhost:9999/usersf \
  # -H 'Content-Type: application/x-www-form-urlencoded' \
  # -H 'Postman-Token: 9da5d67b-950a-46a3-8203-4099e758e20a' \
  # -H 'cache-control: no-cache' \
  # -d 'name=rishi&email=rishi%40gmail.com%0A&undefined='


curl -X POST http://localhost:9999/usersf -F name=name111111 -F email=rishi@gmail.com
