#! /bin/bash

docker-compose up -d

PG_URL="postgres://postgres:PGPWD123@localhost:5432/postgres?sslmode=disable"
MIGRATION_FOLDER="./ext/storage/migrations/postgres/"

migrate -source "file://$MIGRATION_FOLDER" -database $PG_URL up
