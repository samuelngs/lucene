package lucene

// BooleanQuery struct
type BooleanQuery struct {
	Type   string       `json:"type"`
	Must   []CustomRule `json:"must,omitempty"`
	Should []CustomRule `json:"should,omitempty"`
	Not    []CustomRule `json:"not,omitempty"`
}

// Lucene to match CustomRule interface
func (v *BooleanQuery) Lucene() bool {
	return true
}

// BooleanMust query
func BooleanMust(rs ...CustomRule) CustomRule {
	v := &BooleanQuery{
		Type: "boolean",
		Must: make([]CustomRule, len(rs)),
	}
	for i, r := range rs {
		v.Must[i] = r
	}
	return v
}

// BooleanShould query
func BooleanShould(rs ...CustomRule) CustomRule {
	v := &BooleanQuery{
		Type:   "boolean",
		Should: make([]CustomRule, len(rs)),
	}
	for i, r := range rs {
		v.Should[i] = r
	}
	return v
}

// BooleanNot query
func BooleanNot(rs ...CustomRule) CustomRule {
	v := &BooleanQuery{
		Type: "boolean",
		Not:  make([]CustomRule, len(rs)),
	}
	for i, r := range rs {
		v.Not[i] = r
	}
	return v
}
