package lucene

import "testing"

func TestCustom(t *testing.T) {
	c := Custom()
	c.Refresh(true)
}
