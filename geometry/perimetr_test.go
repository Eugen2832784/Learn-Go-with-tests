package geometry

import "testing"

func TestPerimetr(t *testing.T) {
	got := Perimetr(10.2, 56.4)
	want := 133.2
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
