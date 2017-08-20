package visualization_test

import (
	. "visualization"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
)

var _ = Describe("Visualization", func() {
	Describe("aggs can be created from dsl", func() {
		var (
			dsls    []AggDSL
			agg     Agg
			results Aggs

			err error
		)
		Context("where samples are give", func() {
			BeforeEach(func() {
				dsls = []AggDSL{
					`count`,
					`cardinality(City)`,
					`avg(LowestBid)`,
					`max(LowestBid)`,
					`percentiles(LowestBid)[1,3,20,100]`,
					`date_histogram<x>(DateRequested)`,
					`filters<x>[all]`,
					`filters<slice>[all]`,
					`filters<chart>[all]`,
					`terms<x>(City)`,
					`terms<slice>(City)`,
					`terms<slice>(City)`,
				}
				results = Aggs{
					Agg{ID: "0", Enabled: true, Type: "count", Schema: "metric", Params: nil},
					Agg{ID: "1", Enabled: true, Type: "cardinality", Schema: "metric", Params: map[string]interface{}{"field": "City"}},
					Agg{ID: "2", Enabled: true, Type: "avg", Schema: "metric", Params: map[string]interface{}{"field": "LowestBid"}},
					Agg{ID: "3", Enabled: true, Type: "max", Schema: "metric", Params: map[string]interface{}{"field": "LowestBid"}},
					Agg{ID: "4", Enabled: true, Type: "percentiles", Schema: "metric", Params: map[string]interface{}{"field": "coin", "percents": []string{"1", "3", "20", "100"}}},
					Agg{ID: "5", Enabled: true, Type: "date_histogram", Schema: "segment", Params: map[string]interface{}{"field": "DateRequested", "customInterval": "2h", "interval": "auto", "min_doc_count": 1, "extended_bounds": map[string]interface{}{}}},
					Agg{ID: "6", Enabled: true, Type: "filters", Schema: "segment", Params: map[string]interface{}{"filters": []map[string]interface{}{map[string]interface{}{"input": map[string]interface{}{"query": map[string]interface{}{"query_string": map[string]interface{}{"query": "all", "analyze_wildcard": true}}}, "label": "all"}}}},
					Agg{ID: "7", Enabled: true, Type: "filters", Schema: "", Params: map[string]interface{}{"filters": []map[string]interface{}{map[string]interface{}{"input": map[string]interface{}{"query": map[string]interface{}{"query_string": map[string]interface{}{"query": "all", "analyze_wildcard": true}}}, "label": "all"}}}},
					Agg{ID: "8", Enabled: true, Type: "filters", Schema: "split", Params: map[string]interface{}{"filters": []map[string]interface{}{map[string]interface{}{"input": map[string]interface{}{"query": map[string]interface{}{"query_string": map[string]interface{}{"query": "all", "analyze_wildcard": true}}}, "label": "all"}}}},
					Agg{ID: "9", Enabled: true, Type: "terms", Schema: "segment", Params: map[string]interface{}{"field": "City", "size": 5, "order": "desc", "orderBy": "1"}},
					Agg{ID: "10", Enabled: true, Type: "terms", Schema: "", Params: map[string]interface{}{"field": "City", "size": 5, "order": "desc", "orderBy": "1"}},
					Agg{ID: "11", Enabled: true, Type: "terms", Schema: "", Params: map[string]interface{}{"field": "City", "size": 5, "order": "desc", "orderBy": "1"}},
				}
			})

			It("should be proper", func() {
				for i, dsl := range dsls {
					fmt.Println(dsl)
					agg, err = dsl.Parse(i)
					Expect(err).ToNot(HaveOccurred())
					Expect(agg.ID).To(Equal(results[i].ID))
					Expect(agg.Enabled).To(Equal(results[i].Enabled))
					Expect(agg.Type).To(Equal(results[i].Type))
					Expect(agg.Params).To(Equal(results[i].Params))
				}
			})
		})
	})
})
