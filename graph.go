package algorithms

import "math"

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
func (graph *Graph) SetEdge(from int, to int, cost int) {
	if cost < 1 {
		cost = 1
	}
	graph.adjacencyMatrix[from][to] = cost
	graph.adjacencyMatrix[to][from] = cost // non-directed graph
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
			if val == 0 {
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
			if val == 0 {
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

// DijkstraShortestPath algorithm implementation
func (graph *Graph) DijkstraShortestPath(source int) map[int]int {
	dist := make(map[int]int)
	sptSet := make(map[int]bool)

	for i := 0; i < graph.numVertex; i++ {
		dist[i] = math.MaxInt32
		sptSet[i] = false
	}
	dist[source] = 0

	var u int
	var newDistance int
	for count := 0; count < graph.numVertex-1; count++ {
		u = graph.nextMinDistance(dist, sptSet)
		sptSet[u] = true
		for v := 0; v < graph.numVertex; v++ {
			if graph.adjacencyMatrix[u][v] != 0 && sptSet[v] == false {
				if dist[u] != math.MaxInt32 {
					newDistance = dist[u] + graph.adjacencyMatrix[u][v]
					if newDistance < dist[v] {
						dist[v] = newDistance
					}
				}
			}
		}
	}

	return dist
}

func (graph *Graph) nextMinDistance(dist map[int]int, sptSet map[int]bool) int {
	min := math.MaxInt32
	minIndex := -1
	for v := 0; v < graph.numVertex; v++ {
		if sptSet[v] == false && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}
	return minIndex
}
