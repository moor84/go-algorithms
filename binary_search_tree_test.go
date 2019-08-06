package algorithms

import (
	"reflect"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	bst := BinarySearchTree{}
	bst.Insert(36, "one")
	bst.Insert(27, "two")
	bst.Insert(18, "three")
	bst.Insert(59, "four")
	bst.Insert(610, "five")
	bst.Insert(12, "six")
	bst.Insert(2, "seven")
	bst.Insert(34, "eight")
	bst.Insert(46, "nine")
	bst.Insert(51, "ten")
	tests := []struct {
		key      int
		expected string
	}{
		{36, "one"},
		{27, "two"},
		{18, "three"},
		{59, "four"},
		{610, "five"},
		{12, "six"},
		{2, "seven"},
		{34, "eight"},
		{46, "nine"},
		{51, "ten"},
	}
	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result, err := bst.Find(tt.key)
			if result != tt.expected || err != nil {
				t.Errorf("Key: %d, Value: %s", tt.key, result)
			}
		})
	}
}

func TestBinarySearchTree_NotFound(t *testing.T) {
	bst := BinarySearchTree{}
	bst.Insert(10, "bla")
	bst.Insert(7, "bla")
	bst.Insert(12, "bla")
	result, err := bst.Find(331)
	if result != "" {
		t.Error()
	}
	if err == nil {
		t.Error()
	}
}

func TestBinarySearchTree_Delete(t *testing.T) {
	bst := BinarySearchTree{}
	bst.Insert(4, "four")
	bst.Insert(2, "two")
	bst.Insert(6, "six")
	bst.Insert(1, "one")
	bst.Insert(3, "three")
	bst.Insert(5, "five")
	bst.Insert(7, "seven")

	if result, _ := bst.Find(3); result != "three" {
		t.Error()
	}
	bst.Delete(3)
	if result, _ := bst.Find(3); result != "" {
		t.Error()
	}

	bst.Delete(6)
	if result, _ := bst.Find(6); result != "" {
		t.Error()
	}

	bst.Delete(4)
	if result, _ := bst.Find(4); result != "" {
		t.Error()
	}

	if result, _ := bst.Find(7); result != "seven" {
		t.Error()
	}
}

func TestBinarySearchTree_InOrder(t *testing.T) {
	bst := BinarySearchTree{}
	bst.Insert(4, "four")
	bst.Insert(2, "two")
	bst.Insert(6, "six")
	bst.Insert(1, "one")
	bst.Insert(3, "three")
	bst.Insert(5, "five")
	bst.Insert(7, "seven")

	res := bst.InOrder()
	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf("res = %v", res)
	}
}

func TestBinarySearchTree_PreOrder(t *testing.T) {
	bst := BinarySearchTree{}
	bst.Insert(4, "four")
	bst.Insert(2, "two")
	bst.Insert(6, "six")
	bst.Insert(1, "one")
	bst.Insert(3, "three")
	bst.Insert(5, "five")
	bst.Insert(7, "seven")

	res := bst.PreOrder()
	if !reflect.DeepEqual(res, []int{4, 2, 1, 3, 6, 5, 7}) {
		t.Errorf("res = %v", res)
	}
}

func TestBinarySearchTree_PostOrder(t *testing.T) {
	bst := BinarySearchTree{}
	bst.Insert(4, "four")
	bst.Insert(2, "two")
	bst.Insert(6, "six")
	bst.Insert(1, "one")
	bst.Insert(3, "three")
	bst.Insert(5, "five")
	bst.Insert(7, "seven")

	res := bst.PostOrder()
	if !reflect.DeepEqual(res, []int{1, 3, 2, 5, 7, 6, 4}) {
		t.Errorf("res = %v", res)
	}
}

func TestBinarySearchTree_Balance(t *testing.T) {
	bst := BinarySearchTree{}
	bst.Insert(4, "bla")
	bst.Insert(3, "bla")
	bst.Insert(2, "bla")
	bst.Insert(1, "bla")

	bst.Balance()

	if bst.root.key != 2 {
		t.Errorf("Balanced tree, root = %d", bst.root.key)
	}
	res := bst.InOrder()
	if !reflect.DeepEqual(res, []int{1, 2, 3, 4}) {
		t.Errorf("Balanced tree, res = %v", res)
	}
}
