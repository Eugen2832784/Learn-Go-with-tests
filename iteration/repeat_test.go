package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
func ExampleRepeat() {
	repeated := Repeat("s", 6)
	fmt.Println(repeated) // Output: ssssss
}
