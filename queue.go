package algorithms

// Queue - FIFO
type Queue struct {
	LinkedList
}

// Push on the stack
func (q *Queue) Push(data int) {
	q.AddBack(data)
}

// Peek - see the next value
func (q Queue) Peek() int {
	return q.GetFirst()
}

// Pop fron the queue
func (q *Queue) Pop() int {
	head := q.head
	if head == nil {
		return 0
	}
	q.head = head.next
	q.size--
	return head.data
}

// Size of the stack
func (q Queue) Size() int {
	return q.size
}
