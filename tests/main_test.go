package tests

import (
	"testing"
	"test/pkg" // Import the pkg package where Add is defined
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add positive numbers", 1, 2, 3},
		{"Add negative numbers", -1, -2, -3},
		{"Add mixed numbers", -1, 2, 1},
		{"Add zeros", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pkg.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
