#!/bin/sh
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR
echo "Up-ing elastic docker container"
docker-compose up -d elastic-kag
echo "\nUp-ing kibana docker container"
docker-compose up -d kibana-kag
echo "\nSleep while containers get ready (6 seconds)"
sleep 6
echo "\nBuilding binary (default alpine)"
#docker-compose run build-alpine-kag sh -c 'gb build'
docker-compose run build-alpine-kag sh -c 'gb vendor restore && gb build'
echo "\nBootstrap dnd test senario"
docker-compose run build-alpine-kag sh -c '/project/testdata/dnd/bootstrap.sh'
