package lucene

// MatchQuery struct
type MatchQuery struct {
	Type  string      `json:"type"`
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Lucene to match Rule interface
func (v *MatchQuery) Lucene() bool {
	return true
}

// Match query
func Match(ref string, val interface{}) Rule {
	return &MatchQuery{
		Type:  "match",
		Field: ref,
		Value: val,
	}
}
