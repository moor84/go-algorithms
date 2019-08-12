package algorithms

import (
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	graph := NewGraph(6)
	graph.SetVertex(0, "a")
	graph.SetVertex(1, "b")
	graph.SetVertex(2, "c")
	graph.SetVertex(3, "d")
	graph.SetVertex(4, "e")
	graph.SetVertex(5, "f")
	graph.SetEdge(0, 4, 1)
	graph.SetEdge(0, 2, 1)
	graph.SetEdge(2, 1, 1)
	graph.SetEdge(1, 4, 1)
	graph.SetEdge(4, 3, 1)
	graph.SetEdge(5, 4, 1)
	res := graph.BreadthFirst(0)
	if !reflect.DeepEqual(res, []string{"a", "c", "e", "b", "d", "f"}) {
		t.Errorf("BFS: %v", res)
	}
	res = graph.DepthFirst(0)
	if !reflect.DeepEqual(res, []string{"a", "e", "f", "d", "b", "c"}) {
		t.Errorf("DFS: %v", res)
	}
}

func TestGraph_DijkstraShortestPath(t *testing.T) {
	graph := NewGraph(9)
	graph.SetVertex(0, "0")
	graph.SetVertex(1, "1")
	graph.SetVertex(2, "2")
	graph.SetVertex(3, "3")
	graph.SetVertex(4, "4")
	graph.SetVertex(5, "5")
	graph.SetVertex(6, "6")
	graph.SetVertex(7, "7")
	graph.SetVertex(8, "8")
	graph.SetEdge(0, 1, 4)
	graph.SetEdge(0, 7, 8)
	graph.SetEdge(1, 7, 11)
	graph.SetEdge(1, 2, 8)
	graph.SetEdge(7, 6, 1)
	graph.SetEdge(7, 8, 7)
	graph.SetEdge(2, 8, 2)
	graph.SetEdge(8, 6, 6)
	graph.SetEdge(2, 3, 7)
	graph.SetEdge(6, 5, 2)
	graph.SetEdge(2, 5, 4)
	graph.SetEdge(3, 5, 14)
	graph.SetEdge(3, 4, 9)
	graph.SetEdge(5, 4, 10)
	res := graph.DijkstraShortestPath(0)
	if !reflect.DeepEqual(res, map[int]int{
		0: 0,
		1: 4,
		2: 12,
		3: 19,
		4: 21,
		5: 11,
		6: 9,
		7: 8,
		8: 14,
	}) {
		t.Errorf("Dijkstra: %v", res)
	}
}
