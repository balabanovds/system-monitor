#!/bin/sh


docker run --rm=true -it -v "$(pwd)":/app -w /app golang:1.14 go "$@"
