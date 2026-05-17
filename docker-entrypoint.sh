#!/bin/sh
set -e

echo "Entry point: waiting for dependencies and running migrations if needed"

attempts=0
max_attempts=30
until [ $attempts -ge $max_attempts ]
do
  if /usr/local/bin/migrate status >/dev/null 2>&1; then
    echo "Database reachable"
    break
  fi
  attempts=$((attempts+1))
  echo "Waiting for database... attempt $attempts/$max_attempts"
  sleep 2
done

if [ $attempts -ge $max_attempts ]; then
  echo "Database did not become ready in time"
  exit 1
fi

echo "Running DB migrations..."
if ! /usr/local/bin/migrate up; then
  echo "Migrations failed"
  exit 1
fi

echo "Starting API"
exec /usr/local/bin/api
