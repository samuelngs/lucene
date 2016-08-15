package lucene

import "bytes"

// WildcardQuery struct
type WildcardQuery struct {
	Type  string      `json:"type"`
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Lucene to match Rule interface
func (v *WildcardQuery) Lucene() bool {
	return true
}

// Wildcard query
func Wildcard(ref string, val interface{}) CustomRule {
	return &WildcardQuery{
		Type:  "wildcard",
		Field: ref,
		Value: val,
	}
}

// AnyWildcard query
func AnyWildcard(ref string, val string) CustomRule {
	var b bytes.Buffer
	b.WriteRune('*')
	b.WriteString(val)
	b.WriteRune('*')
	return &WildcardQuery{
		Type:  "wildcard",
		Field: ref,
		Value: b.String(),
	}
}

// PrefixWildcard query
func PrefixWildcard(ref string, val string) CustomRule {
	var b bytes.Buffer
	b.WriteRune('*')
	b.WriteString(val)
	return &WildcardQuery{
		Type:  "wildcard",
		Field: ref,
		Value: b.String(),
	}
}

// SuffixWildcard query
func SuffixWildcard(ref string, val string) CustomRule {
	var b bytes.Buffer
	b.WriteString(val)
	b.WriteRune('*')
	return &WildcardQuery{
		Type:  "wildcard",
		Field: ref,
		Value: b.String(),
	}
}
