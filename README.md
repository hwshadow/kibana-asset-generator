# kibana-asset-generator

Got tired of configuring stuff via the kibana GUI ...  so we are going to hackaboo data into the .kibana configuration index.  Why would we do this?
  - You need to instantiate many instances of the same searches, visualizations, dashboards but for slightly different index names.  Wildcard index is not optimal for you.
  - You hate building dashboards in the GUI.

## About
### idxp package
Will pull an index's field mappings and translate that into the equivalent kibana  index-pattern.  Requires an index pattern, time field name, and also config for target to source mapping data and writing resulting index-pattern.

#### Field mappings from elasticsearch
http://localhost:9200/job*/_mapping/*/field/*?include_defaults=false

*payload below is truncated*
```json
{
  "job": {
    "mappings": {
      "record": {
        "_ttl": {
          "full_name": "_ttl",
          "mapping": {}
        },
        "Email": {
          "full_name": "Email",
          "mapping": {
            "Email": {
              "type": "keyword"
            }
          }
        },
        "script.keyword": {
          "full_name": "script.keyword",
          "mapping": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "ChippingOption": {
          "full_name": "ChippingOption",
          "mapping": {
            "ChippingOption": {
              "type": "keyword"
            }
          }
        },
        "TreeJobFolderURL": {
          "full_name": "TreeJobFolderURL",
          "mapping": {
            "TreeJobFolderURL": {
              "type": "text",
              "index": false
            }
          }
        },
        "TreeService": {
          "full_name": "TreeService",
          "mapping": {
            "TreeService": {
              "type": "text",
              "analyzer": "autocomplete",
              "search_analyzer": "standard"
            }
          }
        },
        "Arborist": {
          "full_name": "Arborist",
          "mapping": {
            "Arborist": {
              "type": "boolean"
            }
          }
        },
        "FirstName": {
          "full_name": "FirstName",
          "mapping": {
            "FirstName": {
              "type": "text",
              "analyzer": "autocomplete",
              "search_analyzer": "standard"
            }
          }
        },
        "_type": {
          "full_name": "_type",
          "mapping": {}
        },
        "Concerns": {
          "full_name": "Concerns",
          "mapping": {
            "Concerns": {
              "type": "text",
              "analyzer": "autocomplete",
              "search_analyzer": "standard"
            }
          }
        },
        "_field_names": {
          "full_name": "_field_names",
          "mapping": {}
        },
        "MainServices": {
          "full_name": "MainServices",
          "mapping": {
            "MainServices": {
              "type": "keyword"
            }
          }
        },
        "Emergency": {
          "full_name": "Emergency",
          "mapping": {
            "Emergency": {
              "type": "boolean"
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
	"_id": "job*",
	"_index": ".kibana",
	"_source": {
		"title": "job*",
		"fields": "[{\"name\":\"Arborist\",\"type\":\"boolean\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"ChippingOption\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"City\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"Cleanup\",\"type\":\"boolean\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"ClientFolderURL\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":true,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"ClientIPAddress\",\"type\":\"ip\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"CompAssoc\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"Concerns\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"DateRequested\",\"type\":\"date\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"Description\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"Email\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"Emergency\",\"type\":\"boolean\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"FirstName\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"GotPayment\",\"type\":\"boolean\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"LastName\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"LowestBid\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"MainServices\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"NeedsArborist\",\"type\":\"boolean\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"NeedsGrinding\",\"type\":\"boolean\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"NeedsSplitter\",\"type\":\"boolean\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"NeedsWoodGone\",\"type\":\"boolean\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"PhonePrimary\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"PhoneSecondary\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"Priority\",\"type\":\"boolean\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"ServiceType\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"SourceTracking\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"State\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"Status\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"Street\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"StumpGrinding\",\"type\":\"boolean\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"TaxCode\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"TaxRate\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"TreeJobFolderURL\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":true,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"TreeNumber\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"TreeService\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"UserAgent\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"WoodOption\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":false,\"doc_values\":true,\"searchable\":true,\"aggregatable\":true},{\"name\":\"ZipCode\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"_id\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"_index\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"_score\",\"type\":\"number\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"_source\",\"type\":\"_source\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false},{\"name\":\"_type\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":true,\"aggregatable\":true},{\"name\":\"script\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"searchable\":true,\"aggregatable\":false},{\"name\":\"script.keyword\",\"type\":\"string\",\"count\":0,\"scripted\":false,\"indexed\":false,\"analyzed\":false,\"doc_values\":false,\"searchable\":false,\"aggregatable\":false}]",
		"timeFieldName": "DateRequested"
	},
	"_type": "index-pattern"
}
```

### dashboard package
Will generate kibana dashboard kibanas.  Requires two inputs a **skeleton** and **configuration yaml**.

#### Skeleton
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

#### Configuration Yaml
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
00:
  id: sick_vs_nasty
  type: visualization
01:
  id: state_of_the world
  type: search
  columns:
   - sick
   - nasty
  sort:
   - size
02:
  id: age_ratios
  type: visualization
03:
  id: people
  type: search
  columns:
   - first_name
   - last_name
  sort:
   - age
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
[{
	"col": 1,
	"id": "sick_vs_nasty",
	"panelIndex": 1,
	"row": 1,
	"size_x": 5,
	"size_y": 4,
	"type": "visualization"
}, {
	"col": 8,
	"columns": ["sick", "nasty"],
	"id": "state_of_the world",
	"panelIndex": 8,
	"row": 1,
	"size_x": 5,
	"size_y": 4,
	"sort": ["size"],
	"type": "search"
}, {
	"col": 1,
	"id": "age_ratios",
	"panelIndex": 49,
	"row": 5,
	"size_x": 5,
	"size_y": 4,
	"type": "visualization"
}, {
	"col": 8,
	"columns": ["first_name", "last_name"],
	"id": "people",
	"panelIndex": 56,
	"row": 5,
	"size_x": 5,
	"size_y": 4,
	"sort": ["age"],
	"type": "search"
}, {
	"col": 6,
	"id": "count_nasty",
	"panelIndex": 6,
	"row": 1,
	"size_x": 2,
	"size_y": 4,
	"type": "visualization"
}, {
	"col": 6,
	"id": "count_size",
	"panelIndex": 54,
	"row": 5,
	"size_x": 2,
	"size_y": 2,
	"type": "visualization"
}, {
	"col": 6,
	"id": "count_snakebites",
	"panelIndex": 78,
	"row": 7,
	"size_x": 2,
	"size_y": 2,
	"type": "visualization"
}]
```
#### Generated Kibana Dashboard Object
```json
{
	"_id": "dashboard-noc2",
	"_index": ".kibana",
	"_source": {
		"title": "dashboard-noc2",
		"description": "",
		"version": 1,
		"uiStateJSON": "{\"P-1\":{\"vis\":{\"params\":{\"sort\":{\"columnIndex\":null,\"direction\":null}}}}}",
		"kibanaSavedObjectMeta": {
			"searchSourceJSON": "{\"filter\":[{\"query\":{\"query_string\":{\"query\":\"*\",\"analyze_wildcard\":true}}}]}"
		},
		"optionsJSON": "{\"darkTheme\":true}",
		"panelsJSON": "[{\"col\":8,\"columns\":[\"sick\",\"nasty\"],\"id\":\"state_of_the world\",\"panelIndex\":8,\"row\":1,\"size_x\":5,\"size_y\":4,\"sort\":[\"size\"],\"type\":\"search\"},{\"col\":1,\"id\":\"age_ratios\",\"panelIndex\":49,\"row\":5,\"size_x\":5,\"size_y\":4,\"type\":\"visualization\"},{\"col\":6,\"id\":\"count_size\",\"panelIndex\":54,\"row\":5,\"size_x\":2,\"size_y\":2,\"type\":\"visualization\"},{\"col\":8,\"columns\":[\"first_name\",\"last_name\"],\"id\":\"people\",\"panelIndex\":56,\"row\":5,\"size_x\":5,\"size_y\":4,\"sort\":[\"age\"],\"type\":\"search\"},{\"col\":6,\"id\":\"count_snakebites\",\"panelIndex\":78,\"row\":7,\"size_x\":2,\"size_y\":2,\"type\":\"visualization\"},{\"col\":1,\"id\":\"sick_vs_nasty\",\"panelIndex\":1,\"row\":1,\"size_x\":5,\"size_y\":4,\"type\":\"visualization\"},{\"col\":6,\"id\":\"count_nasty\",\"panelIndex\":6,\"row\":1,\"size_x\":2,\"size_y\":4,\"type\":\"visualization\"}]"
	},
	"_type": "dashboard"
}
```

#### As seen in Kibana
![exampleDash](http://i.imgur.com/ql115H7.png)

## Tested with
- Elasticsearch 5.1.1 / Kibana 5.1.1

## How-to
### Config
- edit /etc/app.yaml (house config for target elasticsearch server)
- edit /etc/dashboard.skeleton  (dashboard widget layout)
- edit /etc/dashboard.yaml (dashboard widget content)

### Install
```sh
$ go get gopkg.in/yaml.v2
$ go build
```
### Run
```sh
$ dash
$ dash -idx="job*" -timeField="DateRequested"
```

## Todos
- Move input out of code
- Implement simplistic Visualization package
- Implement simplistic Search package
- Buff up Dashboard, Visualization, Search package as needed
- Store dashboard layouts + yaml in couchbase
- Store visualization config in couchbase
- Store search config in couchbase
- App to load objects into target elasticsearch server


## License
MIT
