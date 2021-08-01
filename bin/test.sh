#! /bin/bash

PG_URL="postgres://postgres:PGPWD123@localhost:5432/postgres?sslmode=disable"
TEST_POSTGRES_URL=$PG_URL go test ./...