#!/bin/sh
set -e

if [ -n "$DB_HOST" ]; then
  echo "Waiting for MySQL at ${DB_HOST}:${DB_PORT:-3306} ..."
  until nc -z "$DB_HOST" "${DB_PORT:-3306}"; do
    sleep 1
  done
  echo "MySQL is up."
fi

exec /usr/local/bin/server
