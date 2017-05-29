package idxp

import (
	"dash/elastic"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type (
	Fields []Field
	Field  struct {
		Name         string `json:"name"`
		Type         string `json:"type"`
		Count        int    `json:"count"`
		Scripted     bool   `json:"scripted"`
		Indexed      bool   `json:"indexed"`
		Analyzed     bool   `json:"analyzed"`
		DocValues    bool   `json:"doc_values"`
		Searchable   bool   `json:"searchable"`
		Aggregatable bool   `json:"aggregatable"`
	}

	//      index     mapping     type      fieldname
	Mapping     map[string]map[string]map[string]map[string]MappingWrap
	MappingWrap struct {
		FullName string                 `json:"full_name"`
		Mapping  map[string]MappingInfo `json:"mapping"`
	}
	MappingInfo map[string]interface{}
)

func (s Fields) Len() int {
	return len(s)
}
func (s Fields) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Fields) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (mapping Mapping) ToFields() (fields Fields, err error) {
	fields = make(Fields, 0)
	alreadyEncountered := make(map[string]bool)

	for _, indexContent := range mapping {
		//indexName
		for _, typeContent := range indexContent["mappings"] {
			//typeName
			for fieldName, mappingWrap := range typeContent {
				if _, yes := alreadyEncountered[fieldName]; yes {
					continue
				}
				alreadyEncountered[fieldName] = true

				field := Field{Name: fieldName, Count: 0}
				foundType := false
				foundIndex := false
				if fieldName == "_all" || fieldName == "_uid" || fieldName == "_ttl" || fieldName == "_field_names" || fieldName == "_parent" || fieldName == "_timestamp" || fieldName == "_version" || fieldName == "_routing" {
					continue
				}
				for fieldAttribute, value := range mappingWrap.Mapping[fieldName] {
					switch {
					case "type" == fieldAttribute:
						foundType = true
						field.Type = value.(string)
						switch field.Type {
						case "text":
							field.Type = "string"
							field.Analyzed = true
						case "keyword":
							field.Type = "string"
							field.DocValues = true
							field.Aggregatable = true
						case "boolean":
							fallthrough
						case "date":
							fallthrough
						case "ip":
							field.DocValues = true
							field.Searchable = true
							field.Aggregatable = true
						case "float":
							fallthrough
						case "short":
							fallthrough
						case "long":
							fallthrough
						case "byte":
							fallthrough
						case "int":
							field.Type = "number"
							field.DocValues = true
							field.Searchable = true
							field.Aggregatable = true
						}
					case "index" == fieldAttribute:
						foundIndex = true
						castedValue, ok := value.(bool)
						if ok {
							field.Indexed = !(castedValue == false)
						} else {
							castedValue, ok := value.(string)
							if ok {
								field.Indexed = !(castedValue == "not_analyzed" || castedValue == "no")
							}
						}
						field.Searchable = field.Indexed
					case strings.Contains(fieldAttribute, "analyzer"):
						field.Analyzed = true
					case "doc_values" == fieldAttribute:
						field.DocValues = true
					}

				}

				if !foundType {
					field.Type = "string"
				} else if !foundIndex {
					field.Indexed = true
					field.Searchable = true
				}

				//TODO: Resolve these two ugly hackaboos
				if strings.HasPrefix(fieldName, "_type") {
					field.Searchable = true
					field.Aggregatable = true
				}
				if fieldName == "_source" {
					field.Type = "_source"
				}

				fields = append(fields, field)

			}
		}
	}

	//TODO: Resolve this ugly duck
	fields = append(fields, Field{Name: "_score", Type: "number", Count: 0})

	sort.Sort(fields)
	return
}

func (fields Fields) ToDoc(index string, timeFieldName string) (doc elastic.Doc, err error) {
	bytez, err := json.Marshal(fields)
	if err != nil {
		return
	}

	fieldsString := string(bytez)

	source := elastic.KibanaSource{
		Title:  index,
		Fields: fieldsString,
	}

	if timeFieldName != "" {
		source.TimeFieldName = timeFieldName
	}

	doc = elastic.Doc{
		Index:  `.kibana`,
		Type:   "index-pattern",
		Id:     index,
		Source: source,
	}

	return
}

func RenderDoc(index, timeFieldName string) (doc elastic.Doc, err error) {
	data, err := elastic.GlobalClient.GetFieldMappings(elastic.Target{Index: index})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var mapping Mapping
	err = json.Unmarshal(data, &mapping)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fields, err := mapping.ToFields()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	doc, err = fields.ToDoc(index, timeFieldName)
	if err != nil {
		return
	}

	bytez, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(bytez))
	return
}

// grab keys
/// type is whatever elastic said
/// count always 0
/// scripted always false
/// !strings.Contains(key, "index") or it does and that key is "false", "not_analyzed", "no" then searchable: false
/// strings.Contains(key, "analyzer") then analyzer: true
/// strings.Contains(key, "doc_values") then doc_values: true
///!strings.Contains(key, "index") or it does and that key is "false", "not_analyzed", "no" then searchable: false
///  strings.Contains(key, "type") and that key is "text" || that key is "string" and strings.Contains(key, "index") and equal to "false", "not_analyzed", "no" then aggreable: false

// MAPPING INFO
//http://localhost:9200/job/_mapping/*/field/*?include_defaults=false&pretty=true
