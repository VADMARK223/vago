package test

import "testing"

func TestSum(t *testing.T) {
	res := Sum(2, 3)
	if res != 5 {
		t.Fatalf("expected 5, got %d", res)
	}
}

func TestAbs(t *testing.T) {
	testTable := []struct {
		name string
		in   int
		want int
	}{
		{"positive", 5, 5},
		{"negative", -5, 5},
		{"zero", 0, 0},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.in); got != tt.want {
				t.Errorf("Abs(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
