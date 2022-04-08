package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	repeated := Repeat("v", 3)
	fmt.Println(repeated)
	// Output: vvv
}

func TestRepeatv2(t *testing.T) {
	got := strings.Repeat("go", 3)
	want := "gogogo"

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
func BenchmarkRepeatv2(b *testing.B) {
	strings.Repeat("a", b.N)
}
