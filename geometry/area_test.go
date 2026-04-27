package geometry

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape   Shape
		HasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, HasArea: 72.0},
		{shape: Circle{Radius: 10}, HasArea: 314.1592653589793},
		{shape: Triangle{Height: 12, Pamats: 6}, HasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.HasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.HasArea)
		}
	}

}
