package geometry

import "testing"

func TestPerimetr(t *testing.T) {
	rectangle := Rectangle{56.4, 10.2}
	got := Perimetr(rectangle)
	want := 133.2
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{56.4, 10.2}
		checkArea(t, rectangle, 575.28)
	})
	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
