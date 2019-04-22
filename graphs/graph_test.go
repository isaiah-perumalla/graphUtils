package graphs

import (
	"os"
	"reflect"
	"testing"
)

func TestGraph_AddEdge(t *testing.T) {
	g := NewDiGraph(4)
	g.AddEdge(0, 2, -400)
	g.AddEdge(0, 3, -100)
	if 2 != g.EdgeCount() {
		t.Error("incorrect edge count in digraph")
	}
}

func TestParseUndirectedSymGraph(t *testing.T) {
	fileName := "testdata/routes.txt"
	f, err := os.Open(fileName)
	if err != nil {
		t.Errorf("could not open file %s", fileName)
	}
	defer f.Close()
	g := ReadUgraph(" ", f)
	edgeCount := g.Ugraph().edgeCount
	if 18 != edgeCount {
		t.Errorf("inccorect edgeCount expected %d but was %d", 18, edgeCount )
	}
	if "ORD" != g.NameFor(2) {
		t.Errorf("index 2 expected to be JFK but was %s", g.NameFor(0))
	}
	adj := g.Adj("JKF")
	expected := []string{"MCO", "ATL", "ORD"}
	if !reflect.DeepEqual(adj, expected) {
		t.Errorf("incorrect adj list for JFK %v, expected %v ", adj, expected )
	}
}

