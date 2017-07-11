package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"dashboard"
	"elastic"
	"idxp"
	"search"
)

func main() {
	var flagWrite bool
	var flagConfigPath, flagTemplatePath, flagDashTitle, flagPrefix, flagIndex, flagTimeField string
	flag.BoolVar(&flagWrite, "write", false, "Should we write data to elasticsearch.")
	flag.StringVar(&flagConfigPath, "conf", "/project/etc/app.yaml", "Path for application config for kag.")
	flag.StringVar(&flagTemplatePath, "template", "/project/testdata/dnd/kag/", "Path for template config for kag.")
	flag.StringVar(&flagDashTitle, "dashTitle", "", "Title of the dashboard to be rendered.")
	flag.StringVar(&flagPrefix, "pre", "", "Prefix for all elements generated.")
	flag.StringVar(&flagIndex, "idx", "", "Required for index-pattern creation in kibana. An index name (staticor wildcard).")
	flag.StringVar(&flagTimeField, "timeField", "", "Required for index-pattern creation in kibana. Provides the field name that charts and index through time.")
	flag.Parse()

	confBytes, err := ioutil.ReadFile(flagConfigPath)
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
	dashboardDoc, err := dashboard.RenderDoc(flagDashTitle, flagPrefix, dashboardSkelBytes, dashboardYamlBytes)
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
