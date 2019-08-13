package algorithms

import (
	"fmt"
	"testing"
)

func TestClimbStairsMemoization(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		n    int
		want int
	}{
		{2, 2},
		{3, 3},
		{12, 233},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := ClimbStairsMemoization(tt.n); got != tt.want {
				t.Errorf("ClimbStairsMemoization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClimbStairsTabulation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		n    int
		want int
	}{
		{2, 2},
		{3, 3},
		{12, 233},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := ClimbStairsTabulation(tt.n); got != tt.want {
				t.Errorf("ClimbStairsTabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
