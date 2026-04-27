package geometry

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Pamats float64
}

type Shape interface {
	Area() float64
}
