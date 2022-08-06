package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Vic")

	got := buffer.String()
	want := "Hello, Vic"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
