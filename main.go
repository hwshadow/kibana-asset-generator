package main

import (
	"dash/dashboard"
	"encoding/json"
	"fmt"
)

func main() {
	err := renderDash()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func renderDash() (err error) {
	widgetMap, err := dashboard.Skeleton(dashSkeleton)
	if err != nil {
		return
	}

	err = widgetMap.Supplement([]byte(dashYamlz))
	if err != nil {
		return
	}

	err = widgetMap.Validate()
	if err != nil {
		return
	}

	doc, err := widgetMap.ToDoc("dashboard-soc2", "")
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

var dashSkeleton string = `00.__.__.__.__.20.__.01.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.__.__.__.__.__.__.__.__
__.__.__.__.00.__.20.__.__.__.__.01
02.__.__.__.__.21.__.03.__.__.__.__
__.__.__.__.__.__.21.__.__.__.__.__
__.__.__.__.__.22.__.__.__.__.__.__
__.__.__.__.02.__.22.__.__.__.__.03`

var dashYamlz string = `---
00:
  id: histo-00
  type: visualization
01:
  id: search-00
  type: search
  columns:
   - sick
   - nasty
  sort:
   - breast_size
02:
  id: histo-01
  type: visualization
03:
  id: search-01
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
  id: count_21_and_older
  type: visualization
22:
  id: count_snakebites
  type: visualization`

var vizYamlz string = `---
table_terms_count:
 - field: email
   size: 0
   perPage: 10
 - field: first_name
   size: 2
   perPage: 2`
