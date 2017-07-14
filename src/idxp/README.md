## idxp package
Will pull an index's field mappings and translate that into the equivalent kibana  index-pattern.  Requires an index pattern, time field name, and also config for target es host to source mapping data from and write resulting index-pattern to.

### Field mappings from elasticsearch
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
### Generated Kibana Index-Pattern Object
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
