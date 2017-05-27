package visualization

import (
	"dash/document"
	"fmt"
)

type (
	VisualizationDocs []VisualizationDoc
	VisualizationDoc  document.Doc
	VisualizationMap  map[string]Visualizations
	Visualizations    []Visualization
	Visualization     struct {
		Field string `json:"field,omitempty"`
	}
)

func (widgetMap VisualizationMap) ToDocs() (docs VisualizationDocs, err error) {

	// source := document.Source{
	// 	//		Title:       title,
	// 	//		Description: description,
	// 	UIStateJSON: `{"vis":{"params":{"sort":{"columnIndex":null,"direction":null}}}}`,
	// 	Version:     1,
	// }
	//
	// doc = VisualizationDoc{
	// 	Index: `.kibana`,
	// 	Type:  "visualization",
	// 	//		Id:     title,
	// 	Source: source,
	// }

	return
}

func tableTerms(field string, size, perPage int) (visState string) {
	visState = fmt.Sprintf(tableTermsF, field, field, size, perPage)
	return
}

var tableTermsF string = `{
	"title": "%s-table_terms_count",
	"type": "table",
	"params": {
		"perPage": %d,
		"showPartialRows": false,
		"showMeticsAtAllLevels": false,
		"sort": {
			"columnIndex": 0,
			"direction": "desc"
		},
		"showTotal": false,
		"totalFunc": "sum"
	},
	"aggs": [{
		"id": "count",
		"enabled": true,
		"type": "count",
		"schema": "metric",
		"params": {}
	}, {
		"id": "table",
		"enabled": true,
		"type": "terms",
		"schema": "bucket",
		"params": {
			"field": "%s",
			"size": %d,
			"order": "desc",
			"orderBy": "count"
		}
	}],
	"listeners": {}
}`
