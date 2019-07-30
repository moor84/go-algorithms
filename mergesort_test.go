package main

import "testing"
import "reflect"

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{"1 to 12", []int{12, 1, 2, 5, 3, 4, 6, 7, 9, 8, 11, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
		{"Empty", []int{}, []int{}},
		{"Single", []int{1}, []int{1}},
		{"775", []int{7, 7, 5}, []int{5, 7, 7}},
		{"1000 1", []int{1000, 1}, []int{1, 1000}},
		{"1 1000", []int{1, 1000}, []int{1, 1000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.arr)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Error(tt.name)
			}
		})
	}
}
