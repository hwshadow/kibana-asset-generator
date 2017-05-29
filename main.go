package main

import (
	"dash/dashboard"
	"dash/elastic"
	"dash/idxp"
	"dash/search"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	var flagWrite bool
	var flagPrefix, flagIndex, flagTimeField string
	flag.BoolVar(&flagWrite, "write", false, "Should we write data to elasticsearch.")
	flag.StringVar(&flagPrefix, "pre", "", "Prefix for all elements generated.")
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

	var indexPatternDoc elastic.Doc
	var searchDocs elastic.Docs

	if flagIndex != "" {
		//INDEX-PATTERN
		indexPatternDoc, err = idxp.RenderDoc(flagIndex, flagTimeField)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

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
	dashboardDoc, err := dashboard.RenderDoc("jobs", flagPrefix, dashboardSkelBytes, dashboardYamlBytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// WRITE TO ELASTIC
	if flagWrite {
		err = indexPatternDoc.Save()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = searchDocs.Save()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = dashboardDoc.Save()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
