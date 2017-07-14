## dashboard package
Will generate a kibana dashboard.  Requires two inputs: a **dashboard skeleton** and **dashboard yaml**.

### Dashboard Skeleton
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

#### Empty board
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
#### Style 1 (raw)
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
#### Style 2 (walled)
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
#### Style 3 (elegant)
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

### Dashboard Yaml
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
### Generated Panels JSON
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
### Generated Kibana Dashboard Object
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
