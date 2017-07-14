# kibana-asset-generator

Got tired of configuring stuff via the kibana GUI ...  so we are going to hackaboo data into the .kibana configuration index.  Why would we do this?
  - You need to instantiate many instances of the same searches, visualizations, dashboards but for slightly different index names.  Wildcard index is not optimal for you.
  - You hate building dashboards in the GUI.

Currently generates
  - Index Patterns
  - Searches
  - Dashboards

## Tested against
  - elasticsearch:2.4.5 / kibana:4.6.4
  - elasticsearch 5.1.1 / kibana:5.1.1
  - elasticsearch 5.5.0 / kibana:5.5.0

## Todos
https://github.com/hwshadow/kibana-asset-generator/projects/1
  - Implement simplistic Visualization package
  - Buff up Dashboard, Visualization, Search package as needed
  - API to load objects into target elasticsearch server
  - Abstract templates into a database?


## Testing
Made super easy with docker.

### Get docker
https://docs.docker.com/engine/installation/

### (Optionally) tweak the template information
```sh
$ ls ./testdata/dnd/kag_template
```

### Run the bootstrap script
```sh
$ ./dev/bootstrap.sh
```

### Browse to kibana
http://localhost:5601/ to enjoy


## Build
My preference is to use inside a docker container, but if so desired you can build locally.

### From container
#### Get docker
https://docs.docker.com/engine/installation/

#### Build inside a container
```sh
$ cd ./dev/
$ docker-compose run build-alpine-kag sh -c 'gb vendor restore && gb build'
$ ls ../bin/kag
```
By default we build against alpine.
Build behavior can be changed to target debian by using 'build-debian-kag'
Or any other distribution following the pattern in docker-compose.yml

### Locally
#### Get the gb build tool
https://getgb.io/docs/install/
```sh
$ go get github.com/constabulary/gb/...
```

#### Build local
```sh
$ gb vendor restore && gb build
$ ls ./bin/kag
```

## Usage
### Run dry with no connection to elastic
```sh
$ kag -conf="/etc/app.yaml" -template="/etc/" -ops="ids"
```
### Run dry with connection to elastic
```sh
$ kag -conf="/etc/app.yaml" -template="/etc/" -ops="ids" -index="testdata.dnd"
```
### Run dry with connection to elastic only generate searches
```sh
$ kag -conf="/etc/app.yaml" -template="/etc/" -ops="s" -index="testdata.dnd"
```
#### Run with connection to elastic and writes enabled
```sh
$ kag -conf="/etc/app.yaml" -template="./kag_template/" -ops="ids" -index="testdata.dnd" -dashTitle="dash" -prefix="testdata-dnd-" -timeField="dob" -write=true
```

## About
See package specific readmes: idxp, search, dashboard
