#!/bin/bash

VER=$(cat ./VERSION)

docker rmi -f app:$VER
docker build --tag app:$VER .
