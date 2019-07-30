package algorithms

func partition(arr []int, low int, high int) int {
	i := low - 1
	pivot := arr[high]

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func doQuickSort(arr []int, low int, high int) []int {
	if low >= high {
		return arr
	}
	pi := partition(arr, low, high)
	arr = doQuickSort(arr, low, pi-1)
	arr = doQuickSort(arr, pi+1, high)
	return arr
}

// QuickSort quicksort implementation
func QuickSort(arr []int) []int {
	return doQuickSort(arr, 0, len(arr)-1)
}
