#!/bin/bash

echo "Starting postgres ..."

docker network inspect dev_network >/dev/null 2>&1 || \
    docker network create dev_network

docker run -d --rm --name dev-database --network dev_network \
 -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres:13-alpine

echo
echo "Listening port: 5432"
echo "Default database: postgres"
echo "Default username: postgres"
echo "Default password: postgres"
echo

while [ ! "$(docker inspect -f '{{.State.Running}}' dev-database)" ]; do echo "Waiting for postgres container to be up" && sleep 1; done;

postgres_host=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' dev-database)
echo "Started postgres at ${postgres_host} ..."
