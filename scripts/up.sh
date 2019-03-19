#!/bin/sh
set -eu

docker-compose -p petstore -f docker/docker-compose.yml up -d
