package algorithms

import (
	"testing"
)

func TestMinHeap_Insert(t *testing.T) {
	heap := NewMinHeap()
	heap.Insert(15)
	heap.Insert(30)
	heap.Insert(20)
	heap.Insert(400)
	heap.Insert(10)
	heap.Insert(500)
	heap.Insert(100)
	heap.Insert(201)
	if heap.GetMin() != 10 {
		t.Error()
	}
}

func TestMinHeap_ExtractMin(t *testing.T) {
	heap := NewMinHeap()
	heap.Insert(15)
	heap.Insert(30)
	heap.Insert(20)
	heap.Insert(400)
	heap.Insert(10)
	heap.Insert(500)
	heap.Insert(100)
	heap.Insert(201)
	if heap.ExtractMin() != 10 {
		t.Error()
	}
	if heap.ExtractMin() != 15 {
		t.Error()
	}
	if heap.ExtractMin() != 20 {
		t.Error()
	}
	if heap.ExtractMin() != 30 {
		t.Error()
	}
}
