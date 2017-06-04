package query

type (
	Queries map[string]Query
	Query   struct {
		Lucene string
		Label  string
		Color  string
	}
)
