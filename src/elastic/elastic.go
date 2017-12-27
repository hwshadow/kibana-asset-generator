package elastic

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

const HOST_ENDPOINT = `{{.Proto}}://{{.Host}}:{{.Port}}`
const FIELD_MAPPINGS_ENDPOINT = `/{{.Index}}/_mapping/*/field/*?include_defaults=false`

var HOST_TEMPLATE = template.Must(template.New("HOST_ENDPOINT").Parse(HOST_ENDPOINT))
var FIELD_MAPPINGS_TEMPLATE = template.Must(template.New("FIELD_MAPPINGS_ENDPOINT").Parse(FIELD_MAPPINGS_ENDPOINT))

var GlobalClient *Client

type Target struct {
	Index string
	Type  string
	Id    string
}

type Docs []Doc
type Doc struct {
	Id     string       `json:"_id"`
	Index  string       `json:"_index"`
	Source KibanaSource `json:"_source"`
	Type   string       `json:"_type"`
}

type KibanaSource struct {
	// Common
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	//Version     int    `json:"version,omitempty"`
	UIStateJSON string `json:"uiStateJSON,omitempty"`
	// Optional / Search
	KibanaSavedObjectMeta interface{} `json:"kibanaSavedObjectMeta,omitempty"`
	// Dashboard
	OptionsJSON string `json:"optionsJSON,omitempty"`
	PanelsJSON  string `json:"panelsJSON,omitempty"`
	TimeRestore bool   `json:"timeRestore,omitempty"`
	// Visualizations
	VisState      string `json:"visState,omitempty"`
	SavedSearchId string `json:"savedSearchId,omitempty"`
	// Search
	Columns []string `json:"columns,omitempty"`
	Sort    []string `json:"sort,omitempty"`
	// Index-Pattern
	Fields        string `json:"fields,omitempty"`
	TimeFieldName string `json:"timeFieldName,omitempty"`
}

type Client struct {
	Proto       string
	Host        string
	Port        string
	KibanaIndex string `yaml:"kibana_index"`
}

func GetServerAddress() (server string, err error) {
	var serv bytes.Buffer
	err = HOST_TEMPLATE.Execute(&serv, GlobalClient)
	if err != nil {
		return
	}
	server = serv.String()
	return
}

func InitClient(config []byte) (err error) {
	GlobalClient = &Client{}
	err = yaml.Unmarshal(config, GlobalClient)
	return
}

func GetFieldMappings(tar Target) (body []byte, err error) {
	var url bytes.Buffer
	err = FIELD_MAPPINGS_TEMPLATE.Execute(&url, tar)
	if err != nil {
		return
	}
	address, err := GetServerAddress()
	if err != nil {
		return
	}
	resp, err := http.Get(address + url.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
func (doc Doc) BulkIndexStanza() (cmd string) {
	innerDocument := map[string]interface{}{"_index": doc.Index, "_type": doc.Type, "_id": doc.Id}
	bytes, _ := json.Marshal(map[string]interface{}{"index": innerDocument})
	return string(bytes)
}

func (doc Doc) BulkContent() (content string) {
	bytes, _ := json.Marshal(doc.Source)
	return string(bytes)
}

func (doc Doc) Save() (err error) {
	docs := Docs{doc}
	docs.Save()
	return
}

func (docs Docs) Save() (err error) {
	bulk := make([]byte, 0)
	for _, doc := range docs {
		bulk = append(bulk, doc.BulkIndexStanza()...)
		bulk = append(bulk, []byte("\n")...)
		bulk = append(bulk, doc.BulkContent()...)
		bulk = append(bulk, []byte("\n")...)

	}

	address, err := GetServerAddress()
	if err != nil {
		return
	}

	_, err = http.Post(address+"/_bulk", "text/json", bytes.NewBuffer(bulk))
	return
}
