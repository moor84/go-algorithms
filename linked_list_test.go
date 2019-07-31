package algorithms

import (
	"reflect"
	"testing"
)

func TestLinkedList_AddFront(t *testing.T) {
	list := LinkedList{}
	list.AddFront(1)
	list.AddFront(3)
	list.AddFront(5)
	if !reflect.DeepEqual(list.GetArray(), []int{5, 3, 1}) {
		t.Error()
	}
}

func TestLinkedList_GetHead(t *testing.T) {
	list := LinkedList{}
	if !reflect.DeepEqual(list.GetHead(), 0) {
		t.Error()
	}
	list.AddFront(1)
	list.AddFront(3)
	if !reflect.DeepEqual(list.GetHead(), 3) {
		t.Error()
	}
}

func TestLinkedList_GetLast(t *testing.T) {
	list := LinkedList{}
	if list.GetLast() != 0 {
		t.Error()
	}
	list.AddFront(1)
	list.AddFront(3)
	if list.GetLast() != 1 {
		t.Error()
	}
}

func TestLinkedList_GetTail(t *testing.T) {
	list := LinkedList{}
	if list.GetTail() != nil {
		t.Error()
	}
	list.AddFront(1)
	list.AddFront(3)
	if list.GetTail().data != 1 {
		t.Error()
	}
}

func TestLinkedList_AddBack(t *testing.T) {
	list := LinkedList{}
	list.AddBack(1)
	list.AddBack(3)
	list.AddBack(5)
	if !reflect.DeepEqual(list.GetArray(), []int{1, 3, 5}) {
		t.Error()
	}
}

func TestLinkedList_GetArray(t *testing.T) {
	list := LinkedList{}
	if !reflect.DeepEqual(list.GetArray(), []int{}) {
		t.Error()
	}
}

func TestLinkedList_DeleteValue(t *testing.T) {
	list := LinkedList{}

	res := list.DeleteValue(1)
	if res {
		t.Error("Delete from empty list")
	}

	list.AddBack(1)
	list.AddBack(3)
	list.AddBack(5)
	list.AddBack(10)

	res = list.DeleteValue(7)
	if res {
		t.Error("Delete non-existent value")
	}
	if !reflect.DeepEqual(list.GetArray(), []int{1, 3, 5, 10}) {
		t.Error()
	}

	res = list.DeleteValue(3)
	if !res {
		t.Error()
	}
	if !reflect.DeepEqual(list.GetArray(), []int{1, 5, 10}) {
		t.Error()
	}

	res = list.DeleteValue(10)
	if !res {
		t.Error()
	}
	if !reflect.DeepEqual(list.GetArray(), []int{1, 5}) {
		t.Error()
	}
}
