package graphics

import (
	"fmt"
	rtc "raytracerchallenge/core"
)

type Color struct {
	R float64
	G float64
	B float64
}

func (c Color) String() string {
	return fmt.Sprintf("Color {R:%f G:%f B:%f}", c.R, c.G, c.B)
}

func (c Color) Equals(c2 Color) bool {
	return rtc.FpEquals(c.R, c2.R) && rtc.FpEquals(c.G, c2.G) && rtc.FpEquals(c.B, c2.B)
}

func (c Color) Add(c2 Color) Color {
	return Color{c.R + c2.R, c.G + c2.G, c.B + c2.B}
}

func (c Color) Sub(c2 Color) Color {
	return Color{c.R - c2.R, c.G - c2.G, c.B - c2.B}
}

func (c Color) ScalarMult(factor float64) Color {
	return Color{c.R * factor, c.G * factor, c.B * factor}
}

func (c Color) Mult(c2 Color) Color {
	return Color{c.R * c2.R, c.G * c2.G, c.B * c2.B}
}
