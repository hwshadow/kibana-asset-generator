# kibana-asset-generator

Got tired of configuring stuff via the kibana GUI ...  so we are going to hackaboo data into the .kibana configuration index.  Why would we do this?
  - You need to instantiate many instances of the same searches, visualizations, dashboards but for slightly different index names.  Wildcard index is not optimal for you.
  - You hate building dashboards in the GUI.


# About
### dashboard package
Will generate kibana dashboard documents.  Requires two inputs a **skeleton** and **configuration yaml**.

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
#### Panels JSON
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
#### Kibana Dashboard Object
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

### How-to
#### Config
- Edit dashSkeleton variable in main.go
- Edit dashYamlz variable in main.go

#### Install
```sh
$ go get gopkg.in/yaml.v2
$ go build
```
#### Run
```sh
$ ./dash
```

### Todos
- Move input out of code
- Implement simplistic Visualization package
- Implement simplistic Search package
- Implement Index-Pattern package (generate dynamically based on index)
- Buff up Dashboard, Visualization, Search package as needed
- Store dashboard layouts + yaml in couchbase
- Store visualization config in couchbase
- Store search config in couchbase
- App to load objects into target elasticsearch server


License
----
MIT
