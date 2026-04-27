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
