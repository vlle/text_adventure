#!/usr/bin/env bash

docker pull postgres:latest
docker run -d --rm --name testsio  -p 5500:5432  -e POSTGRES_USER=postgres  -e POSTGRES_PASSWORD=postgres  -e POSTGRES_DB=rec postgres -N 2000

