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

	doc, err := widgetMap.ToDoc("dashboard-noc2", "")
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

var dashYamlz string = `00:
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
