package visualization

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"elastic"
)

type (
	VisType   string
	VisSchema string
	VisScale  string
	VisMode   string

	AggDSL string
	Agg    struct {
		ID string `json:"id,omitempty"`
		//Enabled bool                   `json:"enabled,omitempty"`
		Type   VisType                `json:"type,omitempty"`
		Schema VisSchema              `json:"schema,omitempty"`
		Params map[string]interface{} `json:"params,omitempty"`
	}
	Aggs          []Agg
	Visualization struct {
		ID        string                 `json:"-"`
		Title     string                 `json:"title"`
		Type      VisType                `json:"type"`
		Params    map[string]interface{} `json:"params"`
		Aggs      Aggs                   `json:"aggs"`
		Listeners map[string]interface{} `json:"listeners"`

		Query      string   `json:"-"`
		Metrics    []AggDSL `json:"-"`
		Partitions []AggDSL `json:"-"`
	}
	Visualizations []Visualization
)

var (
	Metricc     VisType   = "metric"
	Count       VisType   = "count"
	Max         VisType   = "max"
	Avg         VisType   = "avg"
	Percentiles VisType   = "percentiles"
	Cardinality VisType   = "cardinality"
	Metric      VisSchema = "metric"

	Histogram     VisType   = "histogram"
	DateHistogram VisType   = "date_histogram"
	Segment       VisSchema = "segment"

	Filters VisType   = "filters"
	Terms   VisType   = "terms"
	Bucket  VisSchema = "bucket"
	Split   VisSchema = "split"
	Group   VisSchema = "group"

	Table VisType = "table"

	Linear  VisScale = "linear"
	Stacked VisMode  = "stacked"

	ToSchema = map[string]VisSchema{
		"x":     Segment,
		"slice": Group,
		"chart": Split,
	}
	DSLFormat = regexp.MustCompile(`(?P<type>[A-z]+)(?:<(?P<orientation>[^\>]*)>)?(?:\((?P<field>[^\)]*)\))?(?:\[(?P<list>[^\]]*)\])?(?P<extra>{.*})?`)
)

func (dsl *AggDSL) Parse(id int) (agg Agg, err error) {
	matches := DSLFormat.FindStringSubmatch(string(*dsl))
	typee := VisType(matches[1])
	ori := matches[2]
	field := matches[3]

	l := strings.Split(matches[4], ",")
	list := make([]interface{}, len(l))
	for i := 0; i < len(l); i++ {
		numeric, err := strconv.Atoi(l[i])
		if err != nil {
			list[i] = l[i]
		} else {
			list[i] = numeric
		}
	}

	extra := make(map[string]interface{}, 0)
	json.Unmarshal([]byte(matches[5]), &extra)

	agg = Agg{}
	agg.ID = strconv.Itoa(id)
	//agg.Enabled = true
	agg.Type = typee
	if ori != "" {
		agg.Schema = ToSchema[string(ori)]
	}
	agg.Params = map[string]interface{}{}

	switch typee {
	case Count:
		agg.Schema = Metric
	case Max:
		fallthrough
	case Avg:
		fallthrough
	case Cardinality:
		agg.Schema = Metric
		agg.Params = map[string]interface{}{
			"field": field,
		}
	case Percentiles:
		agg.Schema = Metric
		agg.Params = map[string]interface{}{
			"field":    field,
			"percents": list,
		}
	case Terms:
		agg.Params = map[string]interface{}{
			"field": field,
			"size":  5,
		}
	case Filters:
		filters := make([]map[string]interface{}, 0)
		for _, query := range list {
			filters = append(filters, map[string]interface{}{
				"input": map[string]interface{}{
					"query": map[string]interface{}{
						"query_string": map[string]interface{}{
							"query":            query,
							"analyze_wildcard": true,
						},
					},
				},
				"label": query,
			})
		}
		agg.Params = map[string]interface{}{
			"filters": filters,
		}
	case Histogram:
		agg.Params = map[string]interface{}{
			"extended_bounds": map[string]interface{}{},
			"field":           field,
			"interval":        list[0],
		}
	case DateHistogram:
		agg.Params = map[string]interface{}{
			"field":           field,
			"customInterval":  "2h",
			"interval":        "auto",
			"min_doc_count":   1,
			"extended_bounds": map[string]interface{}{},
		}
	}

	for key, value := range extra {
		agg.Params[key] = value

	}

	return
}

func (visualization *Visualization) Convert() {
	visualization.ID = visualization.Title
	for i, dsl := range visualization.Metrics {
		agg, err := dsl.Parse(i)
		if err != nil {
			return
		}
		visualization.Aggs = append(visualization.Aggs, agg)
	}
	visualization.Metrics = nil
	l := len(visualization.Aggs)
	for i, dsl := range visualization.Partitions {
		agg, err := dsl.Parse(l + i)
		if err != nil {
			return
		}
		visualization.Aggs = append(visualization.Aggs, agg)
	}
	visualization.Partitions = nil
	visualization.Listeners = map[string]interface{}{}

	switch visualization.Type {
	case Metricc:
		visualization.Params = map[string]interface{}{
			"handleNoResults": true,
			"fontSize":        30,
		}
	case Histogram:
		visualization.Params = map[string]interface{}{
			"shareYAxis":      true,
			"addTooltip":      true,
			"addLegend":       true,
			"legendPosition":  "right",
			"scale":           "linear",
			"mode":            "stacked",
			"times":           []interface{}{},
			"addTimeMarker":   false,
			"defaultYExtents": false,
			"setYExtents":     false,
			"yAxis":           map[string]interface{}{},
		}
	case Table:
		for i, agg := range visualization.Aggs {
			if agg.Schema == Group {
				agg.Schema = Bucket
			}
			visualization.Aggs[i] = agg
		}
		visualization.Params = map[string]interface{}{
			"perPage":               10,
			"showPartialRows":       false,
			"showMeticsAtAllLevels": false,
			"sort": map[string]interface{}{
				"columnIndex": nil,
				"direction":   nil,
			},
			"showTotal": false,
			"totalFunc": "sum",
		}
	}

	return
}

func Import(yml []byte) (visualizations Visualizations, err error) {
	err = yaml.Unmarshal(yml, &visualizations)
	if err != nil {
		return
	}

	for _, visualization := range visualizations {
		visualization.Convert()
		visualizations = append(visualizations, visualization)
	}

	return
}

func (visualization *Visualization) ToDoc(index, prefix string) (doc elastic.Doc) {

	bytez, err := json.Marshal(visualization)
	if err != nil {
		return
	}

	source := elastic.KibanaSource{
		Title:                 prefix + visualization.Title,
		VisState:              string(bytez),
		SavedSearchId:         prefix + visualization.Query,
		KibanaSavedObjectMeta: map[string]interface{}{"searchSourceJSON": "{\"filter\":[]}"},
		UIStateJSON:           "{}",
	}
	visualization.Query = ""

	doc = elastic.Doc{
		Index:  elastic.GlobalClient.KibanaIndex,
		Type:   "visualization",
		Id:     source.Title,
		Source: source,
	}

	return
}

func RenderDocs(index, prefix string, yml []byte) (docs elastic.Docs, err error) {
	docs = make([]elastic.Doc, 0)
	visualizations, err := Import(yml)
	if err != nil {
		return
	}

	for _, visualization := range visualizations {
		doc := visualization.ToDoc(index, prefix)
		docs = append(docs, doc)
	}

	bytez, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(bytez))

	return

}
