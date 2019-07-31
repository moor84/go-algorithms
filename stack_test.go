package algorithms

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if !reflect.DeepEqual(stack.GetArray(), []int{3, 2, 1}) {
		t.Error()
	}
}

func TestStack_Peek(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Peek() != 3 {
		t.Error()
	}
}

func TestStack_Pop(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Pop() != 3 {
		t.Error()
	}
	if stack.Pop() != 2 {
		t.Error()
	}
	if stack.Pop() != 1 {
		t.Error()
	}
	if stack.Pop() != 0 {
		t.Error()
	}
}

func TestStack_Size(t *testing.T) {
	stack := Stack{}
	if stack.Size() != 0 {
		t.Error()
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Error()
	}
	stack.Pop()
	stack.Pop()
	if stack.Size() != 1 {
		t.Error()
	}
}
