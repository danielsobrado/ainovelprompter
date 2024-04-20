#!/bin/bash
set -e

# Run the SQL schema script
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -f ./db/model.sql
