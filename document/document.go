package document

type Docs []Doc
type Doc struct {
	Id     string `json:"_id"`
	Index  string `json:"_index"`
	Source Source `json:"_source"`
	Type   string `json:"_type"`
}

type Source struct {
	// Common
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     int    `json:"version"`
	UIStateJSON string `json:"uiStateJSON"`
	// Optional
	KibanaSavedObjectMeta KSOM `json:"kibanaSavedObjectMeta,omitempty"`
	// Dashboard
	OptionsJSON string `json:"optionsJSON,omitempty"`
	PanelsJSON  string `json:"panelsJSON,omitempty"`
	TimeRestore bool   `json:"timeRestore,omitempty"`
	// Visualizations
	VisState      string `json:"visState,omitempty"`
	SavedSearchId string `json:"savedSearchId,omitempty"`
}

type KSOM struct {
	SearchSourceJSON string `json:"searchSourceJSON"`
}
