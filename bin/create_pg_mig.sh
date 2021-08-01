#! /bin/bash

if ! type migrate &> /dev/null; then
    # migrate doesn't exist, install it
    echo "Installing github.com/golang-migrate/migrate"
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
fi


MIGRATION_FOLDER="./ext/storage/migrations/postgres"

mkdir -p $MIGRATION_FOLDER
migrate create -ext sql -dir=$MIGRATION_FOLDER $1