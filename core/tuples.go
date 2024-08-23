package core

import (
	"fmt"
	"math"
)

type Tuple struct {
	X, Y, Z, W float64
}

func (t Tuple) String() string {
	return fmt.Sprintf("Tuple { %f, %f, %f, %f }", t.X, t.Y, t.Z, t.W)
}

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func (t Tuple) Equals(t2 Tuple) bool {
	return FpEquals(t.X, t2.X) && FpEquals(t.Y, t2.Y) && FpEquals(t.Z, t2.Z) && t.W == t2.W
}

func (t Tuple) Add(t2 Tuple) Tuple {
	return Tuple{
		X: t.X + t2.X,
		Y: t.Y + t2.Y,
		Z: t.Z + t2.Z,
		W: t.W + t2.W,
	}
}

func (t Tuple) Sub(t2 Tuple) Tuple {
	return Tuple{
		X: t.X - t2.X,
		Y: t.Y - t2.Y,
		Z: t.Z - t2.Z,
		W: t.W - t2.W,
	}
}

func (t Tuple) Negate() Tuple {
	return Tuple{
		X: -t.X,
		Y: -t.Y,
		Z: -t.Z,
		W: -t.W,
	}
}

func (t Tuple) Mult(factor float64) Tuple {
	return Tuple{
		X: t.X * factor,
		Y: t.Y * factor,
		Z: t.Z * factor,
		W: t.W * factor,
	}
}

func (t Tuple) Div(factor float64) Tuple {
	return Tuple{
		X: t.X / factor,
		Y: t.Y / factor,
		Z: t.Z / factor,
		W: t.W / factor,
	}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(
		math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(t.W, 2),
	)
}

func (t Tuple) Normalize() Tuple {
	mag := t.Magnitude()
	return Tuple{
		t.X / mag,
		t.Y / mag,
		t.Z / mag,
		t.W / mag,
	}
}

func (t Tuple) DotProduct(t2 Tuple) float64 {
	return t.X*t2.X + t.Y*t2.Y + t.Z*t2.Z + t.W*t2.W
}

// function cross(a, b)
// return vector(a.y * b.z - a.z * b.y,
// end function
// a.z * b.x - a.x * b.z,
// a.x * b.y - a.y * b.x)

func (t Tuple) CrossProduct(t2 Tuple) Tuple {
	return NewVector(t.Y*t2.Z-t.Z*t2.Y, t.Z*t2.X-t.X*t2.Z, t.X*t2.Y-t.Y*t2.X)
}

func NewPoint(x float64, y float64, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 1.0}
}

func NewVector(x float64, y float64, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 0.0}
}
