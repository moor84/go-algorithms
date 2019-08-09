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
	graph.SetEdge(0, 4)
	graph.SetEdge(0, 2)
	graph.SetEdge(2, 1)
	graph.SetEdge(1, 4)
	graph.SetEdge(4, 3)
	graph.SetEdge(5, 4)
	res := graph.BreadthFirst(0)
	if !reflect.DeepEqual(res, []string{"a", "c", "e", "b", "d", "f"}) {
		t.Errorf("BFS: %v", res)
	}
	res = graph.DepthFirst(0)
	if !reflect.DeepEqual(res, []string{"a", "e", "f", "d", "b", "c"}) {
		t.Errorf("DFS: %v", res)
	}
}
