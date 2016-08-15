package lucene

import (
	"encoding/json"

	"github.com/gocql/gocql"
)

// CustomRule interface
type CustomRule interface {
	Lucene() bool
}

// CustomQuery struct
type CustomQuery struct {
	Filters        []CustomRule
	Queries        []CustomRule
	PerformRefresh bool
}

// MarshalCQL handler
func (v *CustomQuery) MarshalCQL(_ gocql.TypeInfo) ([]byte, error) {
	return json.Marshal(v)
}

// Custom creates custom expr search
func Custom() *CustomQuery {
	return &CustomQuery{
		Filters: make([]CustomRule, 0),
		Queries: make([]CustomRule, 0),
	}
}

// Filter to add filter rules to custom query
func (v *CustomQuery) Filter(opts ...CustomRule) *CustomQuery {
	for _, o := range opts {
		switch {
		case o == nil:
			continue
		default:
			v.Filters = append(v.Filters, o)
		}
	}
	return v
}

// Query to add filter rules to custom query
func (v *CustomQuery) Query(opts ...CustomRule) *CustomQuery {
	for _, o := range opts {
		switch {
		case o == nil:
			continue
		default:
			v.Queries = append(v.Queries, o)
		}
	}
	return v
}

// Refresh to perform refresh
func (v *CustomQuery) Refresh(b ...bool) *CustomQuery {
	var r = true
	for _, o := range b {
		r = o
		break
	}
	v.PerformRefresh = r
	return v
}
