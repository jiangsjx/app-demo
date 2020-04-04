#!/bin/bash

VER=$(cat ./VERSION)
CWD=$(pwd)

docker rm -f app
docker run -d \
    --name=app \
    --network=host \
    -v $CWD:/etc/app/ \
    -v $CWD:/var/log/app/ \
    app:$VER
