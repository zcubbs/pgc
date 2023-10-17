#!/bin/bash
set -e

# Start PostgreSQL in the background
docker-entrypoint.sh postgres &

POSTGRES_PID=$!

# Wait for PostgreSQL to start
until psql -h localhost -U postgres -c '\l'; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 1
done

echo "PostgreSQL is ready - executing command"

# Start your application
/app/main

# Wait for PostgreSQL to finish
wait "$POSTGRES_PID"
