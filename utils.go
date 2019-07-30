package main

import "fmt"

// PrintArray print an array (slice)
func PrintArray(arr []int) {
	l := len(arr)
	for idx, i := range arr {
		fmt.Printf("%d", i)
		if idx < l-1 {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}
}
