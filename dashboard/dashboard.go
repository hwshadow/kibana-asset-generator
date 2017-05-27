package dashboard

import (
	"bufio"
	"dash/document"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

var layoutPattern *regexp.Regexp = regexp.MustCompile(`^(?:[0-9]|_|\^|\<){2}(?:\.(?:[0-9]|_|\^|\<){2}){11}$`)
var numberPattern *regexp.Regexp = regexp.MustCompile(`^[0-9]{2}$`)

type (
	DashboardDocs []DashboardDoc
	DashboardDoc  document.Doc
	WidgetMap     map[string]Widget
	Widgets       []Widget
	Widget        struct {
		Col        int      `json:"col"`
		Columns    []string `json:"columns,omitempty"`
		ID         string   `json:"id"`
		PanelIndex int      `json:"panelIndex"`
		Row        int      `json:"row"`
		SizeX      int      `json:"size_x"`
		SizeY      int      `json:"size_y"`
		Sort       []string `json:"sort,omitempty"`
		Type       string   `json:"type"`
	}
)

func Skeleton(layout string) (widgetMap WidgetMap, err error) {
	widgetMap = make(WidgetMap, 0)
	row := 1
	scanner := bufio.NewScanner(strings.NewReader(layout))
	for scanner.Scan() {
		line := scanner.Text()
		if !layoutPattern.MatchString(line) {
			err = fmt.Errorf("could not parse line %d '%s'", row, line)
			return
		}

		parts := strings.Split(line, ".")
		for column, part := range parts {
			column += 1
			if numberPattern.MatchString(part) {
				if _, ok := widgetMap[part]; !ok {
					widgetMap[part] = Widget{PanelIndex: (12*(row-1) + column), Col: column, Row: row, ID: part, SizeX: 1, SizeY: 1}
				} else {
					widg := widgetMap[part]
					if widg.SizeX != 1 || widg.SizeY != 1 {
						err = fmt.Errorf("cannot have more than two points for selector '%s'", part)
						return
					}
					widg.SizeX = column - widg.Col + 1
					widg.SizeY = row - widg.Row + 1
					widgetMap[part] = widg
				}
			}
		}
		row += 1
	}

	return
}

func (widgetMap WidgetMap) Supplement(yml []byte) (err error) {
	supplements := make(WidgetMap, 0)
	err = yaml.Unmarshal(yml, &supplements)
	if err != nil {
		return
	}

	for key, supplement := range supplements {
		if widg, ok := widgetMap[key]; ok {
			widg.ID = supplement.ID
			widg.Type = supplement.Type
			widg.Columns = supplement.Columns
			widg.Sort = supplement.Sort
			widgetMap[key] = widg
		}
	}
	return
}

func (widgetMap WidgetMap) Validate() (err error) {
	for key, widget := range widgetMap {
		if widget.Type != "visualization" && widget.Type != "search" {
			err = fmt.Errorf("%s's Type must be either 'search' or 'visualization'", key)
			return
		}

		if widget.Type == "visualization" && widget.Columns != nil {
			err = fmt.Errorf("%s is a visualization and may not have Columns defined", key)
			return
		}

		if widget.Type == "visualization" && widget.Sort != nil {
			err = fmt.Errorf("%s is a visualization and may not have Sort defined", key)
			return
		}

		if widget.Col <= 0 || widget.Col > 11 {
			err = fmt.Errorf("%s's Col %d is not between 1 and 11 inclusively", key, widget.Col)
			return
		}

		if widget.Row == 0 {
			err = fmt.Errorf("%s's Row %d must be non-zero", key, widget.Row)
			return
		}

		if widget.PanelIndex == 0 {
			err = fmt.Errorf("%s's PanelIndex %d must be non-zero", key, widget.PanelIndex)
			return
		}

		if widget.SizeX+widget.Col > 13 {
			err = fmt.Errorf("%s's SizeX %d plus Col index %d must be less than or equal %d", key, widget.SizeX, widget.Col, 13)
			return
		}

		if numberPattern.MatchString(widget.ID) {
			err = fmt.Errorf("%s's ID %s shouldn't remain a number, it should be descriptive", key, widget.ID)
			return
		}
	}
	return
}

func (widgetMap WidgetMap) ToArray() (widgets Widgets) {
	widgets = make(Widgets, 0)

	for _, widget := range widgetMap {
		widgets = append(widgets, widget)
	}

	return
}

func (widgetMap WidgetMap) ToPanels() (panels string, err error) {
	widgets := widgetMap.ToArray()
	bytez, err := json.Marshal(widgets)
	if err != nil {
		return
	}

	panels = string(bytez)

	return
}

func (widgetMap WidgetMap) ToDoc(title, description string) (doc DashboardDoc, err error) {

	ksom := document.KSOM{
		SearchSourceJSON: `{"filter":[{"query":{"query_string":{"query":"*","analyze_wildcard":true}}}]}`,
	}

	panels, err := widgetMap.ToPanels()
	if err != nil {
		return
	}

	source := document.Source{
		Title:                 title,
		Description:           description,
		PanelsJSON:            panels,
		OptionsJSON:           `{"darkTheme":true}`,
		UIStateJSON:           `{"P-1":{"vis":{"params":{"sort":{"columnIndex":null,"direction":null}}}}}`,
		Version:               1,
		TimeRestore:           false,
		KibanaSavedObjectMeta: ksom,
	}

	doc = DashboardDoc{
		Index:  `.kibana`,
		Type:   "dashboard",
		Id:     title,
		Source: source,
	}

	return
}
