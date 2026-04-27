package geometry

import "math"

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (t Triangle) Area() float64 {
	return t.Height * t.Pamats / 2
}
