package algorithms

// Stack - LIFO
type Stack struct {
	LinkedList
}

// Push on the stack
func (stack *Stack) Push(data int) {
	stack.AddFront(data)
}

// Peek - see the next value
func (stack *Stack) Peek() int {
	return stack.GetFirst()
}

// Pop fron the stack
func (stack *Stack) Pop() int {
	head := stack.head
	if head == nil {
		return 0
	}
	stack.head = head.next
	stack.size--
	return head.data
}

// Size of the stack
func (stack Stack) Size() int {
	return stack.size
}
