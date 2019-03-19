#!/bin/sh
set -eu

docker build --force-rm -t "petstore-api:latest" -f "docker/Dockerfile.api" .
docker build --force-rm -t "petstore-db:latest" -f "docker/Dockerfile.db" .
