package iteration

import (
	"fmt"
	"testing"
)
func TestIterate(t *testing.T) {
	got :=Iterate("a",5 );
	want := "aaaaa"

	if got != want {
		t.Errorf("expected %q but got %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a", 5)
	}
}

func ExampleIterate() {
	sum := Iterate("a", 5)
	fmt.Println(sum)
	// Output: aaaaa
}
