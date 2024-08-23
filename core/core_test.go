package core

import "testing"

func TestFpEquals(t *testing.T) {
	a := 7.123
	b := 7.124
	if FpEquals(a, b) {
		t.Errorf("%f and %f should not be FpEqual", a, b)
	}

	c := 8.000000312312
	d := 8.0000000203299329
	if !FpEquals(c, d) {
		t.Errorf("%f and %f should be FpEqual", c, d)
	}
}
