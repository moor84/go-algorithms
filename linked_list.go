package algorithms

// Node is a linked list node
type Node struct {
	data int
	next *Node
}

// LinkedList linked list implementation
type LinkedList struct {
	head *Node
	size int
}

// AddFront adds an element to the front
func (l *LinkedList) AddFront(data int) {
	node := &Node{data: data}

	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}

	l.size++
}

// GetFirst returns the first element
func (l LinkedList) GetFirst() int {
	if l.head == nil {
		return 0
	}
	return l.head.data
}

// GetLast get last element
func (l LinkedList) GetLast() int {
	tail := l.GetTail()
	if tail == nil {
		return 0
	}
	return tail.data
}

// GetTail get tail node
func (l LinkedList) GetTail() *Node {
	if l.head == nil {
		return nil
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current
}

// AddBack add to the tail
func (l *LinkedList) AddBack(data int) {
	node := &Node{data: data}
	if l.head == nil {
		l.head = node
	} else {
		tail := l.GetTail()
		tail.next = node
	}
	l.size++
}

// GetArray get slice of all data
func (l LinkedList) GetArray() []int {
	arr := make([]int, l.size)
	if l.head == nil {
		return arr
	}
	var idx int
	arr[idx] = l.head.data
	current := l.head
	for current.next != nil {
		current = current.next
		idx++
		arr[idx] = current.data
	}
	return arr
}

// DeleteValue delete first matching value
func (l *LinkedList) DeleteValue(data int) bool {
	if l.head == nil {
		return false
	}
	if l.head.data == data {
		l.head = l.head.next
		l.size--
		return true
	}
	current := l.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			l.size--
			return true
		}
		current = current.next
	}
	return false
}

// Reverse - reverse the linked list
func (l *LinkedList) Reverse() {
	var prev *Node
	var next *Node
	current := l.head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

// DetectLoop using Floyd’s Cycle-Finding Algorithm
func (l *LinkedList) DetectLoop() *Node {
	slow := l.head
	fast := l.head
	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return slow
		}
	}
	return nil
}
