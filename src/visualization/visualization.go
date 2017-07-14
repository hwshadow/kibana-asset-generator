package visualization

type (
	Visualizations []Visualization
	Visualization  struct {
		Title  string
		Type   string
		Params map[string]interface{}
		Aggs   Aggs
	}

	Aggs []Agg
	Agg  struct {
		Id      string
		Enabled bool
		Type    string
		Schema  string
		Params  map[string]interface{}
	}

	Definition struct {
		Title      string
		Type       string
		Metrics    Metrics
		Partitions Partitions
	}

	Metrics    []Metric
	Metric     string
	Partitions []Partition
	Partition  string
)

const PatternViews string = `^([a-z]+)(?:<([^<>\[]+)>)?(?:\(([^()]+)\))?(?:\[([^\[\]]+)\])?(?:({.*}))?$`

var (
	HistogramSchemaLookup map[string]string = map[string]string{
		"x":     "segment",
		"slice": "group",
		"chart": "split",
	}

	TableSchemaLookup map[string]string = map[string]string{
		"slice": "bucket",
		"chart": "split",
	}

	AcceptableMetricTypes []string = []string{
		"count",
		"cardinality"
		"max",
		"avg",
	}

	AcceptablePartitionTypes []string = []string{
		"date_histogram",
		"filters",
		"terms",
	}
)
