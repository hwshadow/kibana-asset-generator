### visualization package
Will generate kibana visualizations.  Requires a single input: **visualization yaml**.

#### About the DSL
DSL follows the general format
```
agg<scheme>(fieldname)[cvs]{params}
```
It is used in the metrics and partitions yaml sections

A simple yaml configuration looks like this
```
- title: character-race-dob
  query: characters-all
  type: histogram
  metrics:
  - 'count'
  partitions:
  - 'date_histogram<x>(dob)'
  - 'terms<slice>(race){"size": 1000}'
```
This expresses a desired histogram visualization named "character-race-dob" fed by the "characters-all" search (defined in the search yaml) with Y axis being a basic count, X axis spread across character date of birth, and histogram bars sliced by race (up to 1000).

##### type - REQUIRED

Controls the what visualization will be generated.  It supports the following:
  * histogram - vertical bar chart
  * table - data table
  * metric - good old metric

##### agg - REQUIRED

Controls the aggregation type used.  It supports the following:
  * histogram - aggs a given field at an interval of the first element specific in the csv (integer fields only)
  * date_histogram - aggs a given field at auto calculated by the time window (time fields only)
  * terms - aggs top terms for a given field (by default 5)
  * filters - aggs based on lucene queries in the csv
  * (count|max|average|cardinality|percentiles) - will direct the creation of a metric agg

##### scheme - PARTITIONS ONLY

Controls how the aggregated data is arranged inside a visualization. It supports the following:
  * x - will define aggregate data as the x axis (can only be used on visualization with x axis: histograms)
  * slice - will split aggregate data inside the a chart
  * chart - will split aggregate data across multiple charts

##### csv - AGGS of filters, percentiles, or histogram

Provides any input an aggregation could require.  The input is separated at commas.  For example this could be the percentiles for a percentiles metric agg or it could be several lunece queries for a filter agg.

##### params - OPTIONAL

Provides a way to set/override any params for a given agg.  Say you don't like the default size in terms you could set
```
{"size": 10}
```
This section is interpreted as JSON and must be valid or absent.

##### what can be used where
  * 'count{params}' - metric
  * 'max(field){params}' - metric
  * 'avg(field){params}' - metric
  * 'cardinality(field){params}' - metric
  * 'percentiles(field)[csv]{params}' - metric
  * 'terms<scheme>(field)' - partitions
  * 'filters<scheme>[csv]{params}' - partitions
  * 'date_histogram<scheme>{params}' - partitions  
  * 'histogram<scheme>[csv]{params}' - partitions

#### Search Yaml
```json
---
- title: female-count
  query: characters-female
  type: metric
  metrics:
  - 'count'
- title: male-count
  query: characters-male
  type: metric
  metrics:
  - 'count'
- title: unique-race-count
  query: characters-all
  type: metric
  metrics:
  - 'cardinality(race)'
- title: age-percentiles
  query: characters-all
  type: metric
  metrics:
  - 'percentiles(age)[0,50,90,95,99,100]'
- title: age-avg
  query: characters-all
  type: metric
  metrics:
  - 'avg(age)'
- title: age-avg-race
  query: characters-all
  type: metric
  metrics:
  - 'avg(age)'
  partitions:
  - 'terms<slice>(race){"size": 1000}'
- title: character-coin
  query: characters-all
  type: histogram
  metrics:
  - 'count'
  partitions:
  - 'histogram<x>(coin)[2000]'
- title: character-race
  query: characters-all
  type: histogram
  metrics:
  - 'count'
  partitions:
  - 'terms<x>(race){"size": 1000}'
- title: character-race-dob
  query: characters-all
  type: histogram
  metrics:
  - 'count'
  partitions:
  - 'date_histogram<x>(dob)'
  - 'terms<slice>(race){"size": 1000}'
- title: character-weapons
  query: characters-all
  type: table
  metrics:
  - 'count'
  partitions:
  - 'terms<slice>(weapons){"size": 1000}'
- title: character-sweet-spot
  query: characters-all
  type: histogram
  metrics:
  - 'count'
  partitions:
  - 'filters<x>[weapons:"bow" AND class:"ranger",weapons:"dagger" AND class:"rogue",weapons:"staff" AND (class:"warlock" OR class:"sorcerer"),weapons:"sword" AND class:"paladin"]'

```

#### Generated Kibana Search Objects
```json
[
  {
    "_id": "testdata-dnd-female-count",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-female-count",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"female-count\",\"type\":\"metric\",\"params\":null,\"aggs\":null,\"listeners\":null}",
      "savedSearchId": "testdata-dnd-characters-female"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-male-count",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-male-count",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"male-count\",\"type\":\"metric\",\"params\":null,\"aggs\":null,\"listeners\":null}",
      "savedSearchId": "testdata-dnd-characters-male"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-unique-race-count",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-unique-race-count",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"unique-race-count\",\"type\":\"metric\",\"params\":null,\"aggs\":null,\"listeners\":null}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-age-percentiles",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-age-percentiles",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"age-percentiles\",\"type\":\"metric\",\"params\":null,\"aggs\":null,\"listeners\":null}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-age-avg",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-age-avg",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"age-avg\",\"type\":\"metric\",\"params\":null,\"aggs\":null,\"listeners\":null}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-age-avg-race",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-age-avg-race",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"age-avg-race\",\"type\":\"metric\",\"params\":null,\"aggs\":null,\"listeners\":null}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-character-coin",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-character-coin",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"character-coin\",\"type\":\"histogram\",\"params\":null,\"aggs\":null,\"listeners\":null}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-character-race",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-character-race",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"character-race\",\"type\":\"histogram\",\"params\":null,\"aggs\":null,\"listeners\":null}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-character-race-dob",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-character-race-dob",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"character-race-dob\",\"type\":\"histogram\",\"params\":null,\"aggs\":null,\"listeners\":null}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-character-weapons",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-character-weapons",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"character-weapons\",\"type\":\"table\",\"params\":null,\"aggs\":null,\"listeners\":null}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-female-count",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-female-count",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"female-count\",\"type\":\"metric\",\"params\":{\"fontSize\":30,\"handleNoResults\":true},\"aggs\":[{\"id\":\"0\",\"type\":\"count\",\"schema\":\"metric\"}],\"listeners\":{}}",
      "savedSearchId": "testdata-dnd-characters-female"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-male-count",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-male-count",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"male-count\",\"type\":\"metric\",\"params\":{\"fontSize\":30,\"handleNoResults\":true},\"aggs\":[{\"id\":\"0\",\"type\":\"count\",\"schema\":\"metric\"}],\"listeners\":{}}",
      "savedSearchId": "testdata-dnd-characters-male"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-unique-race-count",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-unique-race-count",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"unique-race-count\",\"type\":\"metric\",\"params\":{\"fontSize\":30,\"handleNoResults\":true},\"aggs\":[{\"id\":\"0\",\"type\":\"cardinality\",\"schema\":\"metric\",\"params\":{\"field\":\"race\"}}],\"listeners\":{}}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-age-percentiles",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-age-percentiles",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"age-percentiles\",\"type\":\"metric\",\"params\":{\"fontSize\":30,\"handleNoResults\":true},\"aggs\":[{\"id\":\"0\",\"type\":\"percentiles\",\"schema\":\"metric\",\"params\":{\"field\":\"age\",\"percents\":[0,50,90,95,99,100]}}],\"listeners\":{}}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-age-avg",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-age-avg",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"age-avg\",\"type\":\"metric\",\"params\":{\"fontSize\":30,\"handleNoResults\":true},\"aggs\":[{\"id\":\"0\",\"type\":\"avg\",\"schema\":\"metric\",\"params\":{\"field\":\"age\"}}],\"listeners\":{}}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-age-avg-race",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-age-avg-race",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"age-avg-race\",\"type\":\"metric\",\"params\":{\"fontSize\":30,\"handleNoResults\":true},\"aggs\":[{\"id\":\"0\",\"type\":\"avg\",\"schema\":\"metric\",\"params\":{\"field\":\"age\"}},{\"id\":\"1\",\"type\":\"terms\",\"schema\":\"group\",\"params\":{\"field\":\"race\",\"size\":1000}}],\"listeners\":{}}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-character-coin",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-character-coin",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"character-coin\",\"type\":\"histogram\",\"params\":{\"addLegend\":true,\"addTimeMarker\":false,\"addTooltip\":true,\"defaultYExtents\":false,\"legendPosition\":\"right\",\"mode\":\"stacked\",\"scale\":\"linear\",\"setYExtents\":false,\"shareYAxis\":true,\"times\":[],\"yAxis\":{}},\"aggs\":[{\"id\":\"0\",\"type\":\"count\",\"schema\":\"metric\"},{\"id\":\"1\",\"type\":\"histogram\",\"schema\":\"segment\",\"params\":{\"extended_bounds\":{},\"field\":\"coin\",\"interval\":2000}}],\"listeners\":{}}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-character-race",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-character-race",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"character-race\",\"type\":\"histogram\",\"params\":{\"addLegend\":true,\"addTimeMarker\":false,\"addTooltip\":true,\"defaultYExtents\":false,\"legendPosition\":\"right\",\"mode\":\"stacked\",\"scale\":\"linear\",\"setYExtents\":false,\"shareYAxis\":true,\"times\":[],\"yAxis\":{}},\"aggs\":[{\"id\":\"0\",\"type\":\"count\",\"schema\":\"metric\"},{\"id\":\"1\",\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"race\",\"size\":1000}}],\"listeners\":{}}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-character-race-dob",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-character-race-dob",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"character-race-dob\",\"type\":\"histogram\",\"params\":{\"addLegend\":true,\"addTimeMarker\":false,\"addTooltip\":true,\"defaultYExtents\":false,\"legendPosition\":\"right\",\"mode\":\"stacked\",\"scale\":\"linear\",\"setYExtents\":false,\"shareYAxis\":true,\"times\":[],\"yAxis\":{}},\"aggs\":[{\"id\":\"0\",\"type\":\"count\",\"schema\":\"metric\"},{\"id\":\"1\",\"type\":\"date_histogram\",\"schema\":\"segment\",\"params\":{\"customInterval\":\"2h\",\"extended_bounds\":{},\"field\":\"dob\",\"interval\":\"auto\",\"min_doc_count\":1}},{\"id\":\"2\",\"type\":\"terms\",\"schema\":\"group\",\"params\":{\"field\":\"race\",\"size\":1000}}],\"listeners\":{}}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-character-weapons",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-character-weapons",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"character-weapons\",\"type\":\"table\",\"params\":{\"perPage\":10,\"showMeticsAtAllLevels\":false,\"showPartialRows\":false,\"showTotal\":false,\"sort\":{\"columnIndex\":null,\"direction\":null},\"totalFunc\":\"sum\"},\"aggs\":[{\"id\":\"0\",\"type\":\"count\",\"schema\":\"metric\"},{\"id\":\"1\",\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"weapons\",\"size\":1000}}],\"listeners\":{}}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  },
  {
    "_id": "testdata-dnd-character-sweet-spot",
    "_index": ".kibana",
    "_source": {
      "title": "testdata-dnd-character-sweet-spot",
      "uiStateJSON": "{}",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"filter\":[]}"
      },
      "visState": "{\"title\":\"character-sweet-spot\",\"type\":\"histogram\",\"params\":{\"addLegend\":true,\"addTimeMarker\":false,\"addTooltip\":true,\"defaultYExtents\":false,\"legendPosition\":\"right\",\"mode\":\"stacked\",\"scale\":\"linear\",\"setYExtents\":false,\"shareYAxis\":true,\"times\":[],\"yAxis\":{}},\"aggs\":[{\"id\":\"0\",\"type\":\"count\",\"schema\":\"metric\"},{\"id\":\"1\",\"type\":\"filters\",\"schema\":\"segment\",\"params\":{\"filters\":[{\"input\":{\"query\":{\"query_string\":{\"analyze_wildcard\":true,\"query\":\"weapons:\\\"bow\\\" AND class:\\\"ranger\\\"\"}}},\"label\":\"weapons:\\\"bow\\\" AND class:\\\"ranger\\\"\"},{\"input\":{\"query\":{\"query_string\":{\"analyze_wildcard\":true,\"query\":\"weapons:\\\"dagger\\\" AND class:\\\"rogue\\\"\"}}},\"label\":\"weapons:\\\"dagger\\\" AND class:\\\"rogue\\\"\"},{\"input\":{\"query\":{\"query_string\":{\"analyze_wildcard\":true,\"query\":\"weapons:\\\"staff\\\" AND (class:\\\"warlock\\\" OR class:\\\"sorcerer\\\")\"}}},\"label\":\"weapons:\\\"staff\\\" AND (class:\\\"warlock\\\" OR class:\\\"sorcerer\\\")\"},{\"input\":{\"query\":{\"query_string\":{\"analyze_wildcard\":true,\"query\":\"weapons:\\\"sword\\\" AND class:\\\"paladin\\\"\"}}},\"label\":\"weapons:\\\"sword\\\" AND class:\\\"paladin\\\"\"}]}}],\"listeners\":{}}",
      "savedSearchId": "testdata-dnd-characters-all"
    },
    "_type": "visualization"
  }
]
```
