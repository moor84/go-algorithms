package main

import (
	"fmt"
)

// MergeSort mergesort implementation for a slice of integers
func MergeSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}
	mid := l / 2

	left := arr[mid:]
	right := arr[:mid]

	left = MergeSort(left)
	right = MergeSort(right)

	var i, j int
	var result []int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func main() {
	fmt.Printf("Mergesort\n")
	arr := []int{12, 1, 2, 5, 3, 4, 6, 7, 9, 8, 11, 10}
	arr = MergeSort(arr)
	PrintArray(arr)
}
