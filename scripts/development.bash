#!/usr/bin/env bash
PG_CONTAINER="pg-dev"

docker-compose up --build
./postgres/migrate.dev.bash


# if docker-compose up -d --build ${PG_CONTAINER}; then
#   ./postgres/migrate.dev.bash
#   docker-compose up --build dev
# fi

