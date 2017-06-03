package main

import (
	"dash/dashboard"
	"dash/elastic"
	"dash/idxp"
	"dash/search"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var flagWrite bool
	var flagOperations, flagPrefix, flagIndex, flagTimeField string
	flag.BoolVar(&flagWrite, "write", false, "Should we write data to elasticsearch.")
	flag.StringVar(&flagOperations, "ops", "", "What ops should be performed. d=dashboard,i=indexpattern,s=search,v=visualizations")
	flag.StringVar(&flagPrefix, "prefix", "", "Prefix for all elements generated.")
	flag.StringVar(&flagIndex, "index", "", "Required for index-pattern and/or search creation in kibana. An index name (static or wildcard).")
	flag.StringVar(&flagTimeField, "timeField", "", "Required for index-pattern creation in kibana. Provides the field name that charts and index through time.")
	flag.Parse()

	confBytes, err := ioutil.ReadFile("./etc/app.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = elastic.InitClient(confBytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var indexPatternDoc, dashboardDoc elastic.Doc
	var searchDocs elastic.Docs

	if flagIndex != "" {
		if strings.Contains(flagOperations, "i") {
			//INDEX-PATTERN
			indexPatternDoc, err = idxp.RenderDoc(flagIndex, flagTimeField)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		if strings.Contains(flagOperations, "s") {
			//SEARCH
			searchYamlBytes, err := ioutil.ReadFile("./etc/search.yaml")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			searchDocs, err = search.RenderDocs(flagIndex, flagPrefix, searchYamlBytes)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}

	if strings.Contains(flagOperations, "d") {
		//DASHBOARD
		dashboardSkelBytes, err := ioutil.ReadFile("./etc/dashboard.skeleton")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		dashboardYamlBytes, err := ioutil.ReadFile("./etc/dashboard.yaml")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		dashboardDoc, err = dashboard.RenderDoc("jobs", flagPrefix, dashboardSkelBytes, dashboardYamlBytes)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// WRITE TO ELASTIC
	if flagWrite {
		if strings.Contains(flagOperations, "i") {
			err = indexPatternDoc.Save()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		if strings.Contains(flagOperations, "s") {
			err = searchDocs.Save()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		if strings.Contains(flagOperations, "d") {
			err = dashboardDoc.Save()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}
