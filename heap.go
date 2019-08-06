package algorithms

const initialCapacity = 5

// MinHeap implementation
type MinHeap struct {
	data     []int
	size     int
	capacity int
}

// NewMinHeap : Create a new MinHeap
func NewMinHeap() *MinHeap {
	minHeap := &MinHeap{
		data:     make([]int, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
	return minHeap
}

// Insert - add element
func (h *MinHeap) Insert(elem int) {
	if h.size == len(h.data) {
		newData := make([]int, h.capacity*2)
		copy(newData, h.data)
		h.data = newData
		h.capacity = h.capacity * 2
	}
	h.data[h.size] = elem
	h.size++
	h.heapifyUp()
}

// GetMin - peek min element
func (h *MinHeap) GetMin() int {
	if h.size == 0 {
		return 0
	}
	return h.data[0]
}

// ExtractMin - remove and return min element
func (h *MinHeap) ExtractMin() int {
	if h.size == 0 {
		return 0
	}
	res := h.data[0]
	h.data[0] = h.data[h.size-1]
	h.size--
	h.heapifyDown()
	return res
}

func (h *MinHeap) heapifyUp() {
	idx := h.size - 1
	for h.parentIdx(idx) >= 0 && h.data[h.parentIdx(idx)] > h.data[idx] {
		parentIdx := h.parentIdx(idx)
		h.data[parentIdx], h.data[idx] = h.data[idx], h.data[parentIdx]
		idx = h.parentIdx(idx)
	}
}

func (h *MinHeap) heapifyDown() {
	idx := 0
	var smallerChildIdx int
	for h.leftIdx(idx) < h.size {
		smallerChildIdx = h.leftIdx(idx)
		if h.rightIdx(idx) < h.size && h.data[h.rightIdx(idx)] < h.data[h.leftIdx(idx)] {
			smallerChildIdx = h.rightIdx(idx)
		}

		if h.data[idx] < h.data[smallerChildIdx] {
			break
		} else {
			h.data[idx], h.data[smallerChildIdx] = h.data[smallerChildIdx], h.data[idx]
		}

		idx = smallerChildIdx
	}
}

func (h *MinHeap) parentIdx(childIdx int) int {
	return (childIdx - 1) / 2
}

func (h *MinHeap) leftIdx(idx int) int {
	return 2*idx + 1
}

func (h *MinHeap) rightIdx(idx int) int {
	return 2*idx + 2
}
