package main

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got: %.2f; want: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		s    Shape
		want float64
	}{
		{"rectangle", &Rectangle{12, 6}, 72.0},
		{"circle", &Circle{10}, 314.1592653589793},
		{"triangle", &Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Area()
			if got != tt.want {
				t.Errorf("got: %g; want: %g", got, tt.want)
			}
		})
	}
}
