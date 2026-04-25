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
	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{34.1, 4}
		got := rectangle.Area()
		want := 136.4
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}
