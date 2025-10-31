package iteration

import "testing"

func TestRepeat(t *testing.T) {
	character := "b"
	repeatCount := 6

	repeated := Repeat(character, repeatCount)
	expected := "bbbbbb"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
