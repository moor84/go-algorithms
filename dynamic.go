package algorithms

// Steps to solve a DP
// 1) Identify if it is a DP problem
// 2) Decide a state expression with
//    least parameters
// 3) Formulate state relationship
// 4) Do tabulation (or add memoization)
// https://www.geeksforgeeks.org/solve-dynamic-programming-problem/

// You are climbing a stair case. It takes n steps to reach to the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
// Note: Given n will be a positive integer.

// https://www.geeksforgeeks.org/tabulation-vs-memoization/
// https://www.geeksforgeeks.org/dynamic-programming/

func climb(start int, n int, cache map[int]int) int {
	val, ok := cache[start]
	if ok {
		return val
	}
	if start == n {
		return 1
	}
	if start > n {
		return 0
	}
	count := 0
	count += climb(start+1, n, cache)
	count += climb(start+2, n, cache)
	cache[start] = count
	return count
}

// ClimbStairsMemoization solves the problem using recursion and memoization
func ClimbStairsMemoization(n int) int {
	cache := make(map[int]int)
	return climb(0, n, cache)
}

// ClimbStairsTabulation solves the problem using tabulation
func ClimbStairsTabulation(n int) int {
	data := make([]int, n+1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	data[0] = 0
	data[1] = 1
	data[2] = 2
	for i := 3; i <= n; i++ {
		data[i] = data[i-1] + data[i-2]
	}
	return data[n]
}
