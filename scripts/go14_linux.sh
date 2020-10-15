#!/bin/sh

NET_NAME=party

docker network create $NET_NAME
docker run --net=$NET_NAME --rm=true -it -v "$(pwd)":/app -w /app golang:1.14 go "$@"
docker network rm $NET_NAME
