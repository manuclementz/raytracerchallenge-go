package graphics

import "testing"

// Scenario: Adding colors
// Given c1 ← color(0.9, 0.6, 0.75)
// And c2 ← color(0.7, 0.1, 0.25)
// Then c1 + c2 = color(1.6, 0.7, 1.0)
func TestColorAdd(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	expected := Color{1.6, 0.7, 1.0}
	res := c1.Add(c2)
	if !expected.Equals(res) {
		t.Errorf("Expected %s + %s to be %s, got %s", c1, c2, expected, res)
	}
}

// Scenario: Subtracting colors
// Given c1 ← color(0.9, 0.6, 0.75)
// And c2 ← color(0.7, 0.1, 0.25)
// Then c1 - c2 = color(0.2, 0.5, 0.5)
func TestColorSub(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	expected := Color{0.2, 0.5, 0.5}
	res := c1.Sub(c2)
	if !expected.Equals(res) {
		t.Errorf("Expected %s - %s to be %s, got %s", c1, c2, expected, res)
	}
}

// Scenario: Multiplying a color by a scalar
// Given c ← color(0.2, 0.3, 0.4)
// Then c * 2 = color(0.4, 0.6, 0.8)
func TestColorScalarMult(t *testing.T) {
	c := Color{0.2, 0.3, 0.4}
	expected := Color{0.4, 0.6, 0.8}
	res := c.ScalarMult(2)
	if !expected.Equals(res) {
		t.Errorf("Expected %s * %f to be %s, got %s", c, 2.0, expected, res)
	}
}

// Scenario: Multiplying colors
// Given c1 ← color(1, 0.2, 0.4)
// And c2 ← color(0.9, 1, 0.1)
// Then c1 * c2 = color(0.9, 0.2, 0.04)
func TestColorMult(t *testing.T) {
	c1 := Color{1, 0.2, 0.4}
	c2 := Color{0.9, 1, 0.1}
	expected := Color{0.9, 0.2, 0.04}
	res := c1.Mult(c2)
	if !expected.Equals(res) {
		t.Errorf("Expected %s * %s to be %s, got %s", c1, c2, expected, res)
	}
}
