package raytracerbook

import (
	"fmt"
	"math"
)

type Tuple struct {
	x, y, z, w float64
}

func (t Tuple) String() string {
	return fmt.Sprintf("Tuple { %f, %f, %f, %f }", t.x, t.y, t.z, t.w)
}

func (t Tuple) IsPoint() bool {
	return t.w == 1.0
}

func (t Tuple) IsVector() bool {
	return t.w == 0.0
}

func (t Tuple) Equals(t2 Tuple) bool {
	return FpEquals(t.x, t2.x) && FpEquals(t.y, t2.y) && FpEquals(t.z, t2.z) && t.w == t2.w
}

func (t Tuple) Add(t2 Tuple) Tuple {
	return Tuple{
		x: t.x + t2.x,
		y: t.y + t2.y,
		z: t.z + t2.z,
		w: t.w + t2.w,
	}
}

func (t Tuple) Sub(t2 Tuple) Tuple {
	return Tuple{
		x: t.x - t2.x,
		y: t.y - t2.y,
		z: t.z - t2.z,
		w: t.w - t2.w,
	}
}

func (t Tuple) Negate() Tuple {
	return Tuple{
		x: -t.x,
		y: -t.y,
		z: -t.z,
		w: -t.w,
	}
}

func (t Tuple) Mult(factor float64) Tuple {
	return Tuple{
		x: t.x * factor,
		y: t.y * factor,
		z: t.z * factor,
		w: t.w * factor,
	}
}

func (t Tuple) Div(factor float64) Tuple {
	return Tuple{
		x: t.x / factor,
		y: t.y / factor,
		z: t.z / factor,
		w: t.w / factor,
	}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(
		math.Pow(t.x, 2) + math.Pow(t.y, 2) + math.Pow(t.z, 2) + math.Pow(t.w, 2),
	)
}

func (t Tuple) Normalize() Tuple {
	mag := t.Magnitude()
	return Tuple{
		t.x / mag,
		t.y / mag,
		t.z / mag,
		t.w / mag,
	}
}

func NewPoint(x float64, y float64, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 1.0}
}

func NewVector(x float64, y float64, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 0.0}
}
