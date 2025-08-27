package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "one element",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "sorted slice",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "reverse sorted slice",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "random order",
			input: []int{3, 1, 4, 1, 5, 9, 2, 6},
			want:  []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:  "with duplicates",
			input: []int{5, 2, 8, 2, 5, 9, 8},
			want:  []int{2, 2, 5, 5, 8, 8, 9},
		},
		{
			name:  "with negative numbers",
			input: []int{-5, 2, -8, 1, 0},
			want:  []int{-8, -5, 0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}
