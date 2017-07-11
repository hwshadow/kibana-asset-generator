package search

import (
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"elastic"
)

type Searches []Search
type Search struct {
	Title       string   `json:"title"`
	Columns     []string `json:"columns"`
	Sort        []string `json:"sort"`
	QueryString string   `yaml:"query",json:"query"`
	Filters     Filters  `json:"filters"`
}

type Filters []Filter
type Filter struct {
	Negate   bool   `json:"negate,omitempty"`
	Index    string `json:"index,omitempty"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
	Alias    string `json:"alias,omitempty"`
}

type KSOMFilters []KSOMFilter
type KSOMFilter struct {
	Meta  Filter            `json:"meta"`
	Query interface{}       `json:"query"`
	State map[string]string `json:"$state"`
}

type KSOM struct {
	Index  string      `json:"index"`
	Query  interface{} `json:"query"`
	Filter KSOMFilters `json:"filter"`
}

func Import(yml []byte) (searches Searches, err error) {
	err = yaml.Unmarshal(yml, &searches)
	return
}

func (search Search) ToDoc(index, prefix string) (doc elastic.Doc) {
	query_string := map[string]map[string]interface{}{
		"query_string": map[string]interface{}{
			"query":            search.QueryString,
			"analyze_wildcard": true,
		}}

	ksomfilters := make([]KSOMFilter, 0)
	for _, filterMeta := range search.Filters {
		filter_query := map[string]map[string]map[string]interface{}{
			"match": map[string]map[string]interface{}{
				filterMeta.Key: map[string]interface{}{
					"query": filterMeta.Value,
					"type":  "phrase",
				},
			},
		}

		ksomfilter := KSOMFilter{
			Meta:  filterMeta,
			Query: filter_query,
			State: map[string]string{"store": "appState"},
		}

		ksomfilters = append(ksomfilters, ksomfilter)
	}

	ksom := KSOM{
		Index:  index,
		Query:  query_string,
		Filter: ksomfilters,
	}

	bytez, err := json.Marshal(ksom)
	if err != nil {
		return
	}

	source := elastic.KibanaSource{
		Title:   prefix + search.Title,
		Columns: search.Columns,
		Sort:    search.Sort,
		KibanaSavedObjectMeta: map[string]string{"searchSourceJSON": string(bytez)},
	}

	doc = elastic.Doc{
		Index:  `.kibana`,
		Type:   "search",
		Id:     source.Title,
		Source: source,
	}

	return
}

func RenderDocs(index, prefix string, yml []byte) (docs elastic.Docs, err error) {
	docs = make([]elastic.Doc, 0)
	searches, err := Import(yml)
	if err != nil {
		return
	}

	for _, search := range searches {
		doc := search.ToDoc(index, prefix)
		docs = append(docs, doc)
	}

	bytez, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(bytez))

	return

}
