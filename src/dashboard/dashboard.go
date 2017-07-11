package dashboard

import (
	"bufio"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"

	"elastic"
)

var layoutPattern *regexp.Regexp = regexp.MustCompile(`(?:[0-9]|_|=|\||\^|\<|\>){2}(?:\.(?:[0-9]|_|=|\||\^|\<|\>){2}){11}`)
var numberPattern *regexp.Regexp = regexp.MustCompile(`[0-9]{2}`)

type (
	WidgetMap map[string]Widget
	Widgets   []Widget
	Widget    struct {
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

func (s Widgets) Len() int {
	return len(s)
}
func (s Widgets) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Widgets) Less(i, j int) bool {
	return s[i].PanelIndex < s[j].PanelIndex
}

func Skeleton(layout []byte) (widgetMap WidgetMap, err error) {
	widgetMap = make(WidgetMap, 0)
	row := 1
	scanner := bufio.NewScanner(strings.NewReader(string(layout)))
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
					widgetMap[part] = Widget{PanelIndex: (12*(row-1) + column), Col: column, Row: row, ID: part}
				} else {
					widg := widgetMap[part]
					if widg.SizeX != 0 || widg.SizeY != 0 {
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

func (widgetMap WidgetMap) Supplement(prefix string, yml []byte) (err error) {
	supplements := make(WidgetMap, 0)
	err = yaml.Unmarshal(yml, &supplements)
	if err != nil {
		return
	}

	for key, supplement := range supplements {
		if widg, ok := widgetMap[key]; ok {
			widg.ID = prefix + supplement.ID
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

		if widget.Col <= 0 || widget.Col > 12 {
			err = fmt.Errorf("%s's Col %d is not between 1 and 12 inclusively", key, widget.Col)
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

		if widget.SizeX == 0 {
			err = fmt.Errorf("%s's SizeX %d must be non-zero", key, widget.SizeX)
			return
		}

		if widget.SizeY == 0 {
			err = fmt.Errorf("%s's SizeY %d must be non-zero", key, widget.SizeY)
			return
		}

		if widget.SizeX+widget.Col > 13 {
			err = fmt.Errorf("%s's SizeX %d plus Col index %d must be less than or equal %d", key, widget.SizeX, widget.Col, 13)
			return
		}

		if numberPattern.MatchString(widget.ID) {
			err = fmt.Errorf("%s's ID %s should remain a number, it should be descriptive", key, widget.ID)
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

	sort.Sort(widgets)

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

func (widgetMap WidgetMap) ToDoc(title, prefix, description string) (doc elastic.Doc, err error) {
	ksom := map[string]interface{}{
		"searchSourceJSON": `{"filter":[{"query":{"query_string":{"query":"*","analyze_wildcard":true}}}]}`,
	}

	panels, err := widgetMap.ToPanels()
	if err != nil {
		return
	}

	source := elastic.KibanaSource{
		Title:                 prefix + title,
		Description:           description,
		PanelsJSON:            panels,
		OptionsJSON:           `{"darkTheme":true}`,
		UIStateJSON:           `{"P-1":{"vis":{"params":{"sort":{"columnIndex":null,"direction":null}}}}}`,
		Version:               1,
		TimeRestore:           false,
		KibanaSavedObjectMeta: ksom,
	}

	doc = elastic.Doc{
		Index:  `.kibana`,
		Type:   "dashboard",
		Id:     source.Title,
		Source: source,
	}

	return
}

func RenderDoc(name, prefix string, skeleton, yaml []byte) (doc elastic.Doc, err error) {
	widgetMap, err := Skeleton(skeleton)
	if err != nil {
		return
	}

	err = widgetMap.Supplement(prefix, yaml)
	if err != nil {
		return
	}

	err = widgetMap.Validate()
	if err != nil {
		return
	}

	doc, err = widgetMap.ToDoc(name, prefix, "")
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
