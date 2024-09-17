package types

import "testing"

func TestOverflow(t *testing.T) {
	of := overflow()
	if of != -128 {
		t.Fail()
	}
}
