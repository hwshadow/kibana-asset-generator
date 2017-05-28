package dashboard

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSkeletonToWidgetMapMultiFormat5_1_1(t *testing.T) {
	assert := assert.New(t)
	skeletons := []string{dash1Skeleton, dash2Skeleton, dash3Skeleton}
	for _, skeleton := range skeletons {
		t.Log(skeleton)
		widgetMap, err := Skeleton(skeleton)
		assert.Nil(err)
		bytez, err := json.Marshal(widgetMap)
		assert.Nil(err)
		t.Log(string(bytez))

		var expectedWidgetMap WidgetMap
		err = json.Unmarshal([]byte(dashWidgetMapFrame), &expectedWidgetMap)
		assert.Nil(err)
		t.Log(expectedWidgetMap)
		assert.Equal(expectedWidgetMap, widgetMap)
	}

	return
}

func TestFromSkeletonToWidgetMapBadInput5_1_1(t *testing.T) {
	assert := assert.New(t)
	skeletons := []string{dash4Skeleton, dash5Skeleton, dash6Skeleton, dash7Skeleton, dash8Skeleton}
	for _, skeleton := range skeletons {
		t.Log(skeleton)
		_, err := Skeleton(skeleton)
		assert.NotNil(err)
	}

	return
}

func TestEnrichWidgetMap5_1_1(t *testing.T) {
	assert := assert.New(t)
	t.Log(dashWidgetMapFrame)
	var widgetMap WidgetMap
	err := json.Unmarshal([]byte(dashWidgetMapFrame), &widgetMap)
	assert.Nil(err)
	err = widgetMap.Supplement([]byte(dashYamlz))
	assert.Nil(err)

	bytez, err := json.Marshal(widgetMap)
	assert.Nil(err)
	t.Log(string(bytez))

	var expectedWidgetMap WidgetMap
	err = json.Unmarshal([]byte(dashWidgetMapEnriched), &expectedWidgetMap)
	assert.Nil(err)
	t.Log(expectedWidgetMap)
	assert.Equal(expectedWidgetMap, widgetMap)

	return
}

func TestFromWidgetMapToWidgets5_1_1(t *testing.T) {
	assert := assert.New(t)
	t.Log(dashWidgetMapEnriched)
	var widgetMap WidgetMap
	err := json.Unmarshal([]byte(dashWidgetMapEnriched), &widgetMap)
	assert.Nil(err)
	widgets := widgetMap.ToArray()

	bytez, err := json.Marshal(widgets)
	assert.Nil(err)
	t.Log(string(bytez))

	var expectedWidgets Widgets
	err = json.Unmarshal([]byte(dashWidgets), &expectedWidgets)
	assert.Nil(err)
	sort.Sort(expectedWidgets)
	t.Log(expectedWidgets)
	assert.Equal(expectedWidgets, widgets)

	return
}

var dashWidgetMapFrame string = `{"00":{"col":1,"id":"00","panelIndex":1,"row":1,"size_x":5,"size_y":4,"type":""},"01":{"col":8,"id":"01","panelIndex":8,"row":1,"size_x":5,"size_y":4,"type":""},"02":{"col":1,"id":"02","panelIndex":49,"row":5,"size_x":5,"size_y":4,"type":""},"03":{"col":8,"id":"03","panelIndex":56,"row":5,"size_x":5,"size_y":4,"type":""},"20":{"col":6,"id":"20","panelIndex":6,"row":1,"size_x":2,"size_y":4,"type":""},"21":{"col":6,"id":"21","panelIndex":54,"row":5,"size_x":2,"size_y":2,"type":""},"22":{"col":6,"id":"22","panelIndex":78,"row":7,"size_x":2,"size_y":2,"type":""}}`

var dashWidgetMapEnriched = `{"00":{"col":1,"id":"sick_vs_nasty","panelIndex":1,"row":1,"size_x":5,"size_y":4,"type":"visualization"},"01":{"col":8,"columns":["sick","nasty"],"id":"state_of_the world","panelIndex":8,"row":1,"size_x":5,"size_y":4,"sort":["size"],"type":"search"},"02":{"col":1,"id":"age_ratios","panelIndex":49,"row":5,"size_x":5,"size_y":4,"type":"visualization"},"03":{"col":8,"columns":["first_name","last_name"],"id":"people","panelIndex":56,"row":5,"size_x":5,"size_y":4,"sort":["age"],"type":"search"},"20":{"col":6,"id":"count_nasty","panelIndex":6,"row":1,"size_x":2,"size_y":4,"type":"visualization"},"21":{"col":6,"id":"count_size","panelIndex":54,"row":5,"size_x":2,"size_y":2,"type":"visualization"},"22":{"col":6,"id":"count_snakebites","panelIndex":78,"row":7,"size_x":2,"size_y":2,"type":"visualization"}}`

var dashWidgets = `[{"col":1,"id":"sick_vs_nasty","panelIndex":1,"row":1,"size_x":5,"size_y":4,"type":"visualization"},{"col":8,"columns":["sick","nasty"],"id":"state_of_the world","panelIndex":8,"row":1,"size_x":5,"size_y":4,"sort":["size"],"type":"search"},{"col":1,"id":"age_ratios","panelIndex":49,"row":5,"size_x":5,"size_y":4,"type":"visualization"},{"col":8,"columns":["first_name","last_name"],"id":"people","panelIndex":56,"row":5,"size_x":5,"size_y":4,"sort":["age"],"type":"search"},{"col":6,"id":"count_nasty","panelIndex":6,"row":1,"size_x":2,"size_y":4,"type":"visualization"},{"col":6,"id":"count_size","panelIndex":54,"row":5,"size_x":2,"size_y":2,"type":"visualization"},{"col":6,"id":"count_snakebites","panelIndex":78,"row":7,"size_x":2,"size_y":2,"type":"visualization"}]`

var dash1Skeleton string = `00.__.__.__.__.20.__.01.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.00.__.20.__.__.__.__.01
02.__.__.__.__.21.__.03.__.__.__.__
__.__.__.__.__.__.21.__.__.__.__.__
__.__.__.__.__.22.__.__.__.__.__.__
__.__.__.__.02.__.22.__.__.__.__.03`

var dash2Skeleton string = `00.||.||.||.||.20.||.01.||.||.||.||
||.__.__.__.||.||.||.||.__.__.__.||
||.__.__.__.||.||.||.||.__.__.__.||
||.||.||.||.00.||.20.||.||.||.||.01
02.||.||.||.||.21.||.03.||.||.||.||
||.__.__.__.||.||.21.||.__.__.__.||
||.__.__.__.||.22.||.||.__.__.__.||
||.||.||.||.02.||.22.||.||.||.||.03`

var dash3Skeleton string = `00.==.==.==.<<.20.<<.01.==.==.==.<<
||.__.__.__.||.||.||.||.__.__.__.||
||.__.__.__.||.||.||.||.__.__.__.||
<<.==.==.==.00.<<.20.==.==.==.==.01
02.==.==.==.<<.21.<<.03.==.==.==.<<
||.__.__.__.||.<<.21.||.__.__.__.||
||.__.__.__.||.22.<<.||.__.__.__.||
<<.==.==.==.02.<<.22.==.==.==.==.03`

var dash4Skeleton string = `00.==.==.<<.20.<<.01.==.==.==.<<
||.__.__.__.||.||.||.||.__.__.__.||
||.__.__.__.||.||.||.||.__.__.__.||
<<.==.==.==.00.<<.20.==.==.==.==.01
02.==.==.==.<<.21.<<.03.==.==.==.<<
||.__.__.__.||.<<.21.||.__.__.__.||
||.__.__.__.||.22.<<.||.__.__.__.||
<<.==.==.==.02.<<.22.==.==.==.==.03`

var dash5Skeleton string = `00.==.==.==.<<.20.<<.01.==.==.==.<<
||.__.__.__.||.||.||.||.__.__.__.||
||.__.__.__.||.||.||.||._.__.__.||
<<.==.==.==.00.<<.20.==.==.==.==.01
02.==.==.==.<<.21.<<.03.==.==.==.<<
||.__.__.__.||.<<.21.||.__.__.__.||
||.__.__.__.||.22.<<.||.__.__.__.||
<<.==.==.==.02.<<.22.==.==.==.==.03`

var dash6Skeleton string = `00.==.==.==.<<.20.<<.01.==.==.==.<<
||.__.__.__.||.||.||.||.__.__.__.||
||.__.__.__.||.||.||.||.__.__.__.||
<<.==.==.==.00.<<.20.==.==.==.==.01
02.==.==.==.<<.21.<<.03.==.==.==.<<
||.__.__.__.||.<<.21.||.__.__.__.||
||.__.__.__.||.22.<<.||.__.__.__.||
<<.==.==.==.02.<<.2.==.==.==.==.03`

var dash7Skeleton string = `00.==.==.==.<<.20.<<.01.==.==.==.<<
||.__.__.__.||.||.||.||.__.__.__.||
|||.__.__.__.||.||.||.||.__.__.__.||
<<.==.==.==.00.<<.20.==.==.==.==.01
02.==.==.==.<<.21.<<.03.==.==.==.<<
||.__.__.__.||.<<.21.||.__.__.__.||
||.__.__.__.||.22.<<.||.__.__.__.||
<<.==.==.==.02.<<.2.==.==.==.==.03`

var dash8Skeleton string = `00.==.==.==.<<.20.<<.01.==.==.==.<<
||.__.__.__.||.||.||.||.__.__.__.||
||.__.__.__.||.||.||.||.__.__.__.||
 <<.==.==.==.00.<<.20.==.==.==.==.01
02.==.==.==.<<.21.<<.03.==.==.==.<<
||.__.__.__.||.<<.21.||.__.__.__.||
||.__.__.__.||.22.<<.||.__.__.__.||
<<.==.==.==.02.<<.2.==.==.==.==.03`

var dashYamlz string = `---
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
    type: visualization`
