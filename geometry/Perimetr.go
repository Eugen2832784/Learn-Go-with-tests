package geometry

func Perimetr(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
