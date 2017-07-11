#!/bin/sh
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR
docker-compose up -d elastic-kag
docker-compose up -d kibana-kag
docker-compose run build-alpine-kag sh -c 'gb build'
docker-compose run build-alpine-kag sh -c '/project/testdata/dnd/bootstrap.sh'
