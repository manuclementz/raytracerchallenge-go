package core

import (
	"math"
	"testing"
)

func TestTupleAsPoint(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 1.0}

	if a.X != 4.3 || a.Y != -4.2 || a.Z != 3.1 || a.W != 1.0 {
		t.Errorf("Expected a to be (4.3, -4.2, 3.1, 1.0), got %v", a)
	}

	if !a.IsPoint() {
		t.Error("Expected a to be a point")
	}

	if a.IsVector() {
		t.Error("Expected a to not be a vector")
	}
}

func TestTupleAsVector(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 0.0}

	if a.X != 4.3 || a.Y != -4.2 || a.Z != 3.1 || a.W != 0.0 {
		t.Errorf("Expected a to be (4.3, -4.2, 3.1, 0.0), got %v", a)
	}

	if a.IsPoint() {
		t.Error("Expected a to not be a point")
	}

	if !a.IsVector() {
		t.Error("Expected a to be a vector")
	}
}

func TestPointMaker(t *testing.T) {
	ptg := NewPoint(4.3, -4.2, 3.1)
	if ptg.IsVector() {
		t.Error("Expected ptg to be a point")
	}
}

func TestVectorMaker(t *testing.T) {
	vct := NewVector(4.3, -4.2, 3.1)
	if vct.IsPoint() {
		t.Error("Expected vct to be a vector")
	}
}

func TestEquals(t *testing.T) {
	t1 := Tuple{X: 1.0, Y: 1.0, Z: 1.0, W: 1.0}
	t2 := Tuple{X: 1.00000002, Y: 1.0, Z: 1.0, W: 1.0}
	t3 := Tuple{X: 2.0, Y: 2.2, Z: 2.0, W: 0.0}
	if !t1.Equals(t2) {
		t.Errorf("Expected t1 to be equal to t2")
	}
	if t1.Equals(t3) {
		t.Errorf("Expected t1 not to be equal to t3")
	}
}

func TestAdd(t *testing.T) {
	t1 := NewVector(1, 1, 1)
	t2 := NewVector(2, 2, 2)
	expected := NewVector(3, 3, 3)
	actual := t1.Add(t2)
	if !expected.Equals(actual) {
		t.Errorf("Expected t1 + t2 to equal Vector(3, 3, 3)")
	}
}

// Scenario: Subtracting two points
// Given p1 ← point(3, 2, 1)
// And p2 ← point(5, 6, 7)
// Then p1 - p2 = vector(-2, -4, -6)
func TestSubPoints(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	p2 := NewPoint(5, 6, 7)
	res := p1.Sub(p2)
	expected := NewVector(-2, -4, -6)
	if !res.Equals(expected) {
		t.Errorf("Expected t %s to be equal to %s", res, expected)
	}
}

// Scenario: Subtracting a vector from a point
// Given p ← point(3, 2, 1)
// And v ← vector(5, 6, 7)
// Then p - v = point(-2, -4, -6)
func TestSubPointVector(t *testing.T) {
	p := NewPoint(3, 2, 1)
	v := NewVector(5, 6, 7)
	res := p.Sub(v)
	expected := NewPoint(-2, -4, -6)
	if !res.Equals(expected) {
		t.Errorf("Expected %s-%s to be equal to %s", p, v, expected)
	}
}

// Scenario: Subtracting a vector from a point
// Given p ← point(3, 2, 1)
// And v ← vector(5, 6, 7)
// Then p - v = point(-2, -4, -6)
func TestSubVectorPoint(t *testing.T) {
	v1 := NewVector(3, 2, 1)
	v2 := NewVector(5, 6, 7)
	res := v1.Sub(v2)
	expected := NewVector(-2, -4, -6)
	if !res.Equals(expected) {
		t.Errorf("Expected %s-%s to be equal to %s", v1, v2, expected)
	}
}

// Scenario: Subtracting a vector from the zero vector
// Given zero ← vector(0, 0, 0)
// And v ← vector(1, -2, 3)
// Then zero - v = vector(-1, 2, -3)
func TestSubVectorZero(t *testing.T) {
	zero := NewVector(0, 0, 0)
	v := NewVector(1, -2, 3)
	res := zero.Sub(v)
	expected := NewVector(-1, 2, -3)
	if !res.Equals(expected) {
		t.Errorf("Expected zero vector minus %s to be equal to %s", v, expected)
	}
}

// Scenario: Negating a tuple
// Given a ← tuple(1, -2, 3, -4)
// Then -a = tuple(-1, 2, -3, 4)
func TestNegateTuple(t *testing.T) {
	a := Tuple{X: 1, Y: -2, Z: 3, W: -4}
	res := a.Negate()
	expected := Tuple{X: -1, Y: 2, Z: -3, W: 4}
	if !res.Equals(expected) {
		t.Errorf("Expected %s.Negate() to be equal to %s", a, expected)
	}
}

func TestTupleMult(t *testing.T) {
	testCases := []struct {
		tuple    Tuple
		factor   float64
		expected Tuple
	}{
		{Tuple{1, -2, 3, -4}, 3.5, Tuple{3.5, -7, 10.5, -14}},
		{Tuple{0, 0, 0, 0}, 2, Tuple{0, 0, 0, 0}},
		{Tuple{1, 2, 3, 4}, 0.5, Tuple{0.5, 1, 1.5, 2}},
		{Tuple{1, 2, 3, 4}, 1, Tuple{1, 2, 3, 4}},
	}

	for _, tc := range testCases {
		result := tc.tuple.Mult(tc.factor)
		if !result.Equals(tc.expected) {
			t.Errorf("Expected %v multiplied by %v to be %v, but got %v", tc.tuple, tc.factor, tc.expected, result)
		}
	}
}

func TestTupleDiv(t *testing.T) {
	testCases := []struct {
		tuple    Tuple
		factor   float64
		expected Tuple
	}{
		{Tuple{1, -2, 3, -4}, 2, Tuple{0.5, -1, 1.5, -2}},
		{Tuple{0, 0, 0, 0}, 2, Tuple{0, 0, 0, 0}},
		{Tuple{1, 2, 3, 4}, 1, Tuple{1, 2, 3, 4}},
		{Tuple{1, 2, 3, 4}, 4, Tuple{0.25, 0.5, 0.75, 1}},
	}

	for _, tc := range testCases {
		result := tc.tuple.Div(tc.factor)
		if !result.Equals(tc.expected) {
			t.Errorf("Expected %v divided by %v to be %v, but got %v", tc.tuple, tc.factor, tc.expected, result)
		}
	}
}

func TestVectorMagnitude(t *testing.T) {
	testCases := []struct {
		vector   Tuple
		expected float64
	}{
		{NewVector(1, 0, 0), 1},
		{NewVector(0, 1, 0), 1},
		{NewVector(0, 0, 1), 1},
		{NewVector(1, 2, 3), math.Sqrt(14)},
		{NewVector(-1, -2, -3), math.Sqrt(14)},
	}

	for _, tc := range testCases {
		result := tc.vector.Magnitude()
		if !FpEquals(result, tc.expected) {
			t.Errorf("Expected magnitude of %v to be %v, but got %v", tc.vector, tc.expected, result)
		}
	}
}

// Scenario: Normalizing vector(4, 0, 0) gives (1, 0, 0)
// Given v ← vector(4, 0, 0)
// Then normalize(v) = vector(1, 0, 0)
// Scenario: Normalizing vector(1, 2, 3)
// Given v ← vector(1, 2, 3)
// Then normalize(v) = approximately vector(0.26726, 0.53452, 0.80178)
// Scenario: The magnitude of a normalized vector Given v ← vector(1, 2, 3)
// When norm ← normalize(v)
// Then magnitude(norm) = 1
func TestVectorNormalize(t *testing.T) {
	testCases := []struct {
		vector   Tuple
		expected Tuple
	}{
		{NewVector(4, 0, 0), NewVector(1, 0, 0)},
		{NewVector(1, 2, 3), NewVector(0.26726, 0.53452, 0.80178)},
	}
	for _, tc := range testCases {
		result := tc.vector.Normalize()
		if !result.Equals(tc.expected) {
			t.Errorf("Expected normalized %v to be %v, but got %v", tc.vector, tc.expected, result)
		}
		if !FpEquals(1, result.Magnitude()) {
			t.Errorf("Expected magnitude of normalized %v to be %v, but got %v", tc.vector, 1, result.Magnitude())
		}
	}
}

// Scenario: The dot product of two tuples
// Given a ← vector(1, 2, 3)
// And b ← vector(2, 3, 4)
// Then dot(a, b) = 20
func TestVectorDotProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)
	expected := 20.0
	result := a.DotProduct(b)
	if !FpEquals(expected, result) {
		t.Errorf("Expected dot product of %s and %s to be %v, got %v", a, b, expected, result)
	}
}

// Scenario: The cross product of two vectors
// Given a ← vector(1, 2, 3)
// And b ← vector(2, 3, 4)
// Then cross(a, b) = vector(-1, 2, -1)
// And cross(b, a) = vector(1, -2, 1)
func TestVectorCrossProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)
	expectedAB := NewVector(-1, 2, -1)
	expectedBA := NewVector(1, -2, 1)
	resAB := a.CrossProduct(b)
	resBA := b.CrossProduct(a)
	if !resAB.Equals(expectedAB) {
		t.Errorf("Expected cross product of %s and %s to be %v, got %v", a, b, expectedAB, resAB)
	}
	if !resBA.Equals(expectedBA) {
		t.Errorf("Expected cross product of %s and %s to be %v, got %v", b, a, expectedBA, resBA)
	}
}
