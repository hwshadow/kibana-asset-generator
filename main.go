package main

import (
	"dash/dashboard"
	"dash/elastic"
	"dash/idxp"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	var flagIndex, flagTimeField string
	flag.StringVar(&flagIndex, "idx", "", "Required for index-pattern creation in kibana. An index name (staticor wildcard).")
	flag.StringVar(&flagTimeField, "timeField", "", "Required for index-pattern creation in kibana. Provides the field name that charts and index through time.")
	flag.Parse()

	confBytes, err := ioutil.ReadFile("./etc/app.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	elastic.GlobalClient, err = elastic.CreateClient(confBytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if flagIndex != "" && flagTimeField != "" {
		indexPatternDoc, err := idxp.RenderDoc(flagIndex, flagTimeField)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		target := elastic.Target{indexPatternDoc.Index, indexPatternDoc.Type, indexPatternDoc.Id}
		err = elastic.GlobalClient.SaveIndexPattern(target, indexPatternDoc.Source)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	skelBytes, err := ioutil.ReadFile("./etc/dashboard.skeleton")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	yamlBytes, err := ioutil.ReadFile("./etc/dashboard.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = dashboard.RenderDoc("dashboard-noc2", skelBytes, yamlBytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
