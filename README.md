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
### idxp package
Will pull an index's field mappings and translate that into the equivalent kibana  index-pattern.  Requires an index pattern, time field name, and also config for target es host to source mapping data from and write resulting index-pattern to.

#### Field mappings from elasticsearch
http://localhost:9200/job*/_mapping/*/field/*?include_defaults=false

*payload below is truncated*
```json
{
  "testdata.dnd": {
    "mappings": {
      "testdata": {
        "backstory": {
          "full_name": "backstory",
          "mapping": {
            "backstory": {
              "type": "text"
            }
          }
        },
        "_ttl": {
          "full_name": "_ttl",
          "mapping": {}
        },
        "_index": {
          "full_name": "_index",
          "mapping": {}
        },
        "last_name.keyword": {
          "full_name": "last_name.keyword",
          "mapping": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "city": {
          "full_name": "city",
          "mapping": {
            "city": {
              "type": "text",
              "fields": {
                "keyword": {
                  "type": "keyword",
                  "ignore_above": 256
                }
              }
            }
          }
        },
        "_all": {
          "full_name": "_all",
          "mapping": {}
        },
        "_parent": {
          "full_name": "_parent",
          "mapping": {}
        },
        "sex.keyword": {
          "full_name": "sex.keyword",
          "mapping": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "_timestamp": {
          "full_name": "_timestamp",
          "mapping": {}
        },
        "_version": {
          "full_name": "_version",
          "mapping": {}
        },
        "class": {
          "full_name": "class",
          "mapping": {
            "class": {
              "type": "keyword"
            }
          }
        },
        "first_name": {
          "full_name": "first_name",
          "mapping": {
            "first_name": {
              "type": "text",
              "fields": {
                "keyword": {
                  "type": "keyword",
                  "ignore_above": 256
                }
              }
            }
          }
        },
        "_routing": {
          "full_name": "_routing",
          "mapping": {}
        },
        "race": {
          "full_name": "race",
          "mapping": {
            "race": {
              "type": "keyword"
            }
          }
        },
        "level": {
          "full_name": "level",
          "mapping": {
            "level": {
              "type": "long"
            }
          }
        },
        "sex": {
          "full_name": "sex",
          "mapping": {
            "sex": {
              "type": "text",
              "fields": {
                "keyword": {
                  "type": "keyword",
                  "ignore_above": 256
                }
              }
            }
          }
        },
        "_type": {
          "full_name": "_type",
          "mapping": {}
        },
        "first_name.keyword": {
          "full_name": "first_name.keyword",
          "mapping": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "last_name": {
          "full_name": "last_name",
          "mapping": {
            "last_name": {
              "type": "text",
              "fields": {
                "keyword": {
                  "type": "keyword",
                  "ignore_above": 256
                }
              }
            }
          }
        },
        "city.keyword": {
          "full_name": "city.keyword",
          "mapping": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "_field_names": {
          "full_name": "_field_names",
          "mapping": {}
        },
        "dob": {
          "full_name": "dob",
          "mapping": {
            "dob": {
              "type": "date"
            }
          }
        },
        "_source": {
          "full_name": "_source",
          "mapping": {}
        },
        "_id": {
          "full_name": "_id",
          "mapping": {}
        },
        "_uid": {
          "full_name": "_uid",
          "mapping": {}
        },
        "age": {
          "full_name": "age",
          "mapping": {
            "age": {
              "type": "long"
            }
          }
        },
        "coin": {
          "full_name": "coin",
          "mapping": {
            "coin": {
              "type": "float"
            }
          }
        }
      }
    }
  }
}
```
#### Generated Kibana Index-Pattern Object
```json
{
  "_index": ".kibana",
  "_type": "index-pattern",
  "_id": "testdata.dnd",
  "_score": 1,
  "_source": {
    "title": "testdata.dnd",
    "fields": "[{\"name\":\"_id\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"_index\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"_score\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"_source\",\"type\":\"_source\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"_type\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":true,\"aggregatable\":true},{\"name\":\"age\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"backstory\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"city\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"city.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"class\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"coin\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"dob\",\"type\":\"date\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"first_name\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"first_name.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"last_name\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"last_name.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"level\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"race\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"sex\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"sex.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false}]",
    "timeFieldName": "dob"
  }
}
```

### searches package
Will generate kibana searches.  Requires a single input: **search yaml**.

#### Search Yaml
```json
---
- title: characters-all
  columns:
  - first_name
  - last_name
  - race
  - class
  - age
  - level
  - coin
  - city
  - backstory
  - dob
  - _id
  sort:
  - race
  - asc
  query: "*"
- title: female-thirty-or-under-not-tiefling
  columns:
  - first_name
  - last_name
  - race
  - class
  - age
  sort:
  - age
  - desc
  query: age:<=30
  filters:
  - key: sex
    value: female
  - key: race
    value: tiefling
    negate: true
```

#### Generated Kibana Search Objects
```json
[
  {
    "_index": ".kibana",
    "_type": "search",
    "_id": "testdata-dnd-characters-all",
    "_score": 1,
    "_source": {
      "title": "testdata-dnd-characters-all",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"testdata.dnd\",\"query\":{\"query_string\":{\"analyze_wildcard\":true,\"query\":\"*\"}},\"filter\":[]}"
      },
      "columns": [
        "first_name",
        "last_name",
        "race",
        "class",
        "age",
        "level",
        "coin",
        "city",
        "backstory",
        "dob",
        "_id"
      ],
      "sort": [
        "race",
        "asc"
      ]
    }
  },
  {
    "_index": ".kibana",
    "_type": "search",
    "_id": "testdata-dnd-female-thirty-or-under-not-tiefling",
    "_score": 1,
    "_source": {
      "title": "testdata-dnd-female-thirty-or-under-not-tiefling",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"index\":\"testdata.dnd\",\"query\":{\"query_string\":{\"analyze_wildcard\":true,\"query\":\"age:\\u003c=30\"}},\"filter\":[{\"meta\":{\"negate\":false,\"index\":\"testdata.dnd\",\"key\":\"sex\",\"value\":\"female\",\"disabled\":false,\"alias\":null},\"query\":{\"match\":{\"sex\":{\"query\":\"female\",\"type\":\"phrase\"}}},\"$state\":{\"store\":\"appState\"}},{\"meta\":{\"negate\":true,\"index\":\"testdata.dnd\",\"key\":\"race\",\"value\":\"tiefling\",\"disabled\":false,\"alias\":null},\"query\":{\"match\":{\"race\":{\"query\":\"tiefling\",\"type\":\"phrase\"}}},\"$state\":{\"store\":\"appState\"}}]}"
      },
      "columns": [
        "first_name",
        "last_name",
        "race",
        "class",
        "age"
      ],
      "sort": [
        "age",
        "desc"
      ]
    }
  }
]
```

### dashboard package
Will generate a kibana dashboard.  Requires two inputs: a **dashboard skeleton** and **dashboard yaml**.

#### Dashboard Skeleton
- is a visual representation of a kibana dashboard
- each row is delimited by a newline, infinite rows are allowed
- each column is delimited by a period ".", there are always exactly 12 columns in each row
- a cell is valid if it contains a two-digit numeric, "__", "||", "==", ">>", "<<", or "^^"
- a cell holding "__", "||", "==", ">>", "<<", or "^^" do not hold any significance other than styling
- a cell holding at two-digit numeric will serve as instructions for building a widget
   - first occurrence defines the top-left coordinate of a widget
   - second occurrence defines the bottom-right coordinate of a widget
   - if only a single occurrence the widget will consume a 1x1 square
   - the numeric directly correlates to a dictionary key in the configuration yaml

##### Empty board
```
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
```
##### Style 1 (raw)
```
00.__.__.__.__.20.__.01.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.00.__.20.__.__.__.__.01
02.__.__.__.__.21.__.03.__.__.__.__
__.__.__.__.__.__.21.__.__.__.__.__
__.__.__.__.__.22.__.__.__.__.__.__
__.__.__.__.02.__.22.__.__.__.__.03
```
##### Style 2 (walled)
```
00.||.||.||.||.20.||.01.||.||.||.||
||.__.__.__.||.||.||.||.__.__.__.||
||.__.__.__.||.||.||.||.__.__.__.||
||.||.||.||.00.||.20.||.||.||.||.01
02.||.||.||.||.21.||.03.||.||.||.||
||.__.__.__.||.||.21.||.__.__.__.||
||.__.__.__.||.22.||.||.__.__.__.||
||.||.||.||.02.||.22.||.||.||.||.03
```
##### Style 3 (elegant)
```
00.==.==.==.<<.20.<<.01.==.==.==.<<
||.__.__.__.||.||.||.||.__.__.__.||
||.__.__.__.||.||.||.||.__.__.__.||
<<.==.==.==.00.<<.20.==.==.==.==.01
02.==.==.==.<<.21.<<.03.==.==.==.<<
||.__.__.__.||.<<.21.||.__.__.__.||
||.__.__.__.||.22.<<.||.__.__.__.||
<<.==.==.==.02.<<.22.==.==.==.==.03
```

#### Dashboard Yaml
The yaml config indicates what each widget is to become/linked to.  Valid entries are:
- *id*: name of a saved visualization or search
- *type*: "visualization" or "search"
- *columns*:
     - for search only
     - array of fields that will make columns
     - top-down runs columns left-right
- *sort*:
     - for search only
     - array of fields that the search table will be sorted on
     - sorts top-down

```yaml
01:
  id: characters-all
  type: search
  columns:
  - first_name
  - last_name
  - race
  - class
  - age
  - level
  - coin
  - city
  - backstory
  - dob
  - _id
  sort:
  - race
  - asc
03:
  id: female-thirty-or-under-not-tiefling
  type: search
  columns:
  - first_name
  - last_name
  - race
  - class
  - age
  sort:
  - age
  - desc
00:
  id: sick_vs_nasty
  type: visualization
02:
  id: age_ratios
  type: visualization
20:
  id: count_nasty
  type: visualization
21:
  id: count_size
  type: visualization
22:
  id: count_snakebites
  type: visualization
```
#### Generated Panels JSON
Notice that the skeleton and yaml combine to create this portion
```json
[
  {
    "col": 1,
    "columns": [
      "first_name",
      "last_name",
      "race",
      "class",
      "age",
      "level",
      "coin",
      "city",
      "backstory",
      "dob",
      "_id"
    ],
    "id": "testdata-dnd-characters-all",
    "panelIndex": 1,
    "row": 1,
    "size_x": 12,
    "size_y": 5,
    "sort": [
      "race",
      "asc"
    ],
    "type": "search"
  },
  {
    "col": 1,
    "columns": [
      "first_name",
      "last_name",
      "race",
      "class",
      "age"
    ],
    "id": "testdata-dnd-female-thirty-or-under-not-tiefling",
    "panelIndex": 61,
    "row": 6,
    "size_x": 12,
    "size_y": 3,
    "sort": [
      "age",
      "desc"
    ],
    "type": "search"
  },
  {
    "col": 1,
    "id": "testdata-dnd-sick_vs_nasty",
    "panelIndex": 97,
    "row": 9,
    "size_x": 5,
    "size_y": 4,
    "type": "visualization"
  },
  {
    "col": 6,
    "id": "testdata-dnd-count_nasty",
    "panelIndex": 102,
    "row": 9,
    "size_x": 2,
    "size_y": 4,
    "type": "visualization"
  },
  {
    "col": 1,
    "id": "testdata-dnd-age_ratios",
    "panelIndex": 145,
    "row": 13,
    "size_x": 5,
    "size_y": 4,
    "type": "visualization"
  },
  {
    "col": 6,
    "id": "testdata-dnd-count_size",
    "panelIndex": 150,
    "row": 13,
    "size_x": 2,
    "size_y": 2,
    "type": "visualization"
  },
  {
    "col": 6,
    "id": "testdata-dnd-count_snakebites",
    "panelIndex": 174,
    "row": 15,
    "size_x": 2,
    "size_y": 2,
    "type": "visualization"
  }
]
```
#### Generated Kibana Dashboard Object
```json
{
  "_id": "testdata-dnd-dash",
  "_index": ".kibana",
  "_source": {
    "title": "testdata-dnd-dash",
    "version": 1,
    "uiStateJSON": "{\"P-1\":{\"vis\":{\"params\":{\"sort\":{\"columnIndex\":null,\"direction\":null}}}}}",
    "kibanaSavedObjectMeta": {
      "searchSourceJSON": "{\"filter\":[{\"query\":{\"query_string\":{\"query\":\"*\",\"analyze_wildcard\":true}}}]}"
    },
    "optionsJSON": "{\"darkTheme\":true}",
    "panelsJSON": "[{\"col\":1,\"columns\":[\"first_name\",\"last_name\",\"race\",\"class\",\"age\",\"level\",\"coin\",\"city\",\"backstory\",\"dob\",\"_id\"],\"id\":\"testdata-dnd-characters-all\",\"panelIndex\":1,\"row\":1,\"size_x\":12,\"size_y\":5,\"sort\":[\"race\",\"asc\"],\"type\":\"search\"},{\"col\":1,\"columns\":[\"first_name\",\"last_name\",\"race\",\"class\",\"age\"],\"id\":\"testdata-dnd-female-thirty-or-under-not-tiefling\",\"panelIndex\":61,\"row\":6,\"size_x\":12,\"size_y\":3,\"sort\":[\"age\",\"desc\"],\"type\":\"search\"},{\"col\":1,\"id\":\"testdata-dnd-sick_vs_nasty\",\"panelIndex\":97,\"row\":9,\"size_x\":5,\"size_y\":4,\"type\":\"visualization\"},{\"col\":6,\"id\":\"testdata-dnd-count_nasty\",\"panelIndex\":102,\"row\":9,\"size_x\":2,\"size_y\":4,\"type\":\"visualization\"},{\"col\":1,\"id\":\"testdata-dnd-age_ratios\",\"panelIndex\":145,\"row\":13,\"size_x\":5,\"size_y\":4,\"type\":\"visualization\"},{\"col\":6,\"id\":\"testdata-dnd-count_size\",\"panelIndex\":150,\"row\":13,\"size_x\":2,\"size_y\":2,\"type\":\"visualization\"},{\"col\":6,\"id\":\"testdata-dnd-count_snakebites\",\"panelIndex\":174,\"row\":15,\"size_x\":2,\"size_y\":2,\"type\":\"visualization\"}]"
  },
  "_type": "dashboard"
}
```

#### As seen in Kibana
![exampleDash](http://i.imgur.com/ql115H7.png)


## License
MIT
