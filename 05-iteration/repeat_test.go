package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("got: %q; want: %q", repeated, expected)
	}
}

var result string

func BenchmarkRepeat(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = Repeat("a", 6)
	}

	result = r
}
