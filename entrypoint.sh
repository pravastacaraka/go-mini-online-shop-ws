#!/bin/bash

set -e

# Wait for the PostgreSQL database to be ready
until pg_isready -h ${POSTGRES_HOST} -p ${POSTGRES_PORT}; do
  sleep 1
done

# Set the SSL mode based on the APP_ENV value
SSL_MODE="disable"
if [ "$APP_ENV" = "production" ]; then
  SSL_MODE="require"
fi

# Run migrations
migrate -path ${APP_HOME}/db/migrations -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${SSL_MODE} up

# Start the application
exec "$@"
