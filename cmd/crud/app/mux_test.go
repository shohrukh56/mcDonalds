package app

import "testing"

func Test_calculateWeight(t *testing.T) {
	tests := []struct {
		name string
		pattern string
		want int
	}{
		{"zero for root", "/", 0},
		{"one for root + subpath", "/catalog", 1},
		{"two for root + subpath (with trailing slash)", "/catalog/", 2},
		{"three for root + hierarchical subpath", "/catalog/1", 3},
		{"four for root + hierarchical subpath (with trailing slash)", "/catalog/1/", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWeight(tt.pattern); got != tt.want {
				t.Errorf("calculateWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}