package algorithms

import (
	"reflect"
	"testing"
)

func TestQueue_Push(t *testing.T) {
	q := Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	if !reflect.DeepEqual(q.GetArray(), []int{1, 2, 3}) {
		t.Error()
	}
}

func TestQueue_Peek(t *testing.T) {
	q := Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	if q.Peek() != 1 {
		t.Error()
	}
}

func TestQueue_Pop(t *testing.T) {
	q := Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	if q.Pop() != 1 {
		t.Error()
	}
	if q.Pop() != 2 {
		t.Error()
	}
	if q.Pop() != 3 {
		t.Error()
	}
	if q.Pop() != 0 {
		t.Error()
	}
}

func TestQueue_Size(t *testing.T) {
	q := Queue{}
	if q.Size() != 0 {
		t.Error()
	}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	if q.Size() != 3 {
		t.Error()
	}
	q.Pop()
	q.Pop()
	if q.Size() != 1 {
		t.Error()
	}
}
