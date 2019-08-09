package algorithms

// Graph implementation
type Graph struct {
	numVertex       int
	adjacencyMatrix [][]int
	vertices        map[int]string
}

// NewGraph creates a graph
func NewGraph(numVertex int) *Graph {
	graph := &Graph{
		numVertex:       numVertex,
		adjacencyMatrix: make([][]int, numVertex),
		vertices:        make(map[int]string),
	}
	for i := 0; i < numVertex; i++ {
		graph.adjacencyMatrix[i] = make([]int, numVertex)
	}
	return graph
}

// SetVertex adds data to a vertex
func (graph *Graph) SetVertex(id int, data string) {
	if id < 0 || id > graph.numVertex-1 {
		return
	}
	graph.vertices[id] = data
}

// SetEdge adds an edge
func (graph *Graph) SetEdge(from int, to int) {
	graph.adjacencyMatrix[from][to] = 1
	graph.adjacencyMatrix[to][from] = 1 // non-directed graph
}

// BreadthFirst traversal
func (graph *Graph) BreadthFirst(start int) []string {
	result := make([]string, graph.numVertex)
	if start < 0 || start > graph.numVertex-1 {
		return result
	}
	visited := make([]bool, graph.numVertex)
	queue := Queue{}
	current := start
	n := 0

	visited[current] = true
	queue.Push(current)

	for queue.Size() != 0 {
		current = queue.Pop()
		result[n] = graph.vertices[current]
		n++
		for i, val := range graph.adjacencyMatrix[current] {
			if val != 1 {
				continue
			}
			if !visited[i] {
				visited[i] = true
				queue.Push(i)
			}
		}
	}

	return result
}

// DepthFirst traversal
func (graph *Graph) DepthFirst(start int) []string {
	result := make([]string, graph.numVertex)
	if start < 0 || start > graph.numVertex-1 {
		return result
	}
	visited := make([]bool, graph.numVertex)
	stack := Stack{}
	current := start
	n := 0

	visited[current] = true
	stack.Push(current)

	for stack.Size() != 0 {
		current = stack.Pop()
		result[n] = graph.vertices[current]
		n++
		for i, val := range graph.adjacencyMatrix[current] {
			if val != 1 {
				continue
			}
			if !visited[i] {
				visited[i] = true
				stack.Push(i)
			}
		}
	}

	return result
}
