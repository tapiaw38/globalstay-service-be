#!/bin/sh

until nc -z globalstay-mongo-db 27017; do
  echo "Esperando a MongoDB..."
  sleep 2
done

echo "========== Starting Go application =========="
exec go run ./cmd/api/ --host 0.0.0.0