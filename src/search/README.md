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
