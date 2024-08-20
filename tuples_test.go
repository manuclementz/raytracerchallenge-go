package raytracerbook

import "testing"

func TestTupleAsPoint(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 1.0}

	if a.x != 4.3 || a.y != -4.2 || a.z != 3.1 || a.w != 1.0 {
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

	if a.x != 4.3 || a.y != -4.2 || a.z != 3.1 || a.w != 0.0 {
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
	t1 := Tuple{x: 1.0, y: 1.0, z: 1.0, w: 1.0}
	t2 := Tuple{x: 1.00000002, y: 1.0, z: 1.0, w: 1.0}
	t3 := Tuple{x: 2.0, y: 2.2, z: 2.0, w: 0.0}
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
