package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"dashboard"
	"elastic"
	"idxp"
	"search"
	"visualization"
)

func main() {
	var flagWrite bool
	var flagConfigPath, flagOperations, flagTemplatePath, flagDashTitle, flagPrefix, flagIndex, flagTimeField string
	flag.BoolVar(&flagWrite, "write", false, "Should we write data to elasticsearch.")
	flag.StringVar(&flagOperations, "ops", "", "What ops should be performed. d=dashboard,i=indexpattern,s=search,v=visualizations")
	flag.StringVar(&flagConfigPath, "conf", "/project/etc/app.yaml", "Path for application config for kag.")
	flag.StringVar(&flagTemplatePath, "template", "/project/testdata/dnd/kag/", "Path for template config for kag.")
	flag.StringVar(&flagDashTitle, "dashTitle", "", "Title of the dashboard to be rendered.")
	flag.StringVar(&flagPrefix, "prefix", "", "Prefix for all elements generated.")
	flag.StringVar(&flagIndex, "index", "", "Required for index-pattern creation in kibana. An index name (staticor wildcard).")
	flag.StringVar(&flagTimeField, "timeField", "", "Required for index-pattern creation in kibana. Provides the field name that charts and index through time.")
	flag.Parse()

	confBytes, err := ioutil.ReadFile(flagConfigPath)
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
	var searchDocs, visualizationDocs elastic.Docs

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
			searchYamlBytes, err := ioutil.ReadFile(flagTemplatePath + "search.yaml")
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

		if strings.Contains(flagOperations, "v") {
			//VISUALIZATION
			visualizationYamlBytes, err := ioutil.ReadFile(flagTemplatePath + "visualization.yaml")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			visualizationDocs, err = visualization.RenderDocs(flagIndex, flagPrefix, visualizationYamlBytes)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}

	if strings.Contains(flagOperations, "d") {
		//DASHBOARD
		dashboardSkelBytes, err := ioutil.ReadFile(flagTemplatePath + "dashboard.skeleton")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		dashboardYamlBytes, err := ioutil.ReadFile(flagTemplatePath + "dashboard.yaml")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		dashboardDoc, err = dashboard.RenderDoc(flagDashTitle, flagPrefix, dashboardSkelBytes, dashboardYamlBytes)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// WRITE TO ELASTIC
	allDocs := make(elastic.Docs, 0)

	if flagWrite {
		if strings.Contains(flagOperations, "i") {
			allDocs = append(allDocs, indexPatternDoc)
		}
		if strings.Contains(flagOperations, "s") {
			allDocs = append(allDocs, searchDocs...)
		}
		if strings.Contains(flagOperations, "v") {
			allDocs = append(allDocs, visualizationDocs...)
		}
		if strings.Contains(flagOperations, "d") {
			allDocs = append(allDocs, dashboardDoc)
		}
		err = allDocs.Save()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
