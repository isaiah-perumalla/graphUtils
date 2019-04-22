package graphs

import (
	"bufio"
	"os"
	"strings"
)

type Edge struct {
	from, to uint32
	wt int32
}

type Digraph struct {
	n, edgeCount uint32 //vertex count
	adjList [][]Edge

}

type LabeledUgraph struct {
	g             *Ugraph
	labels        map[uint32]string
	invertedIndex map[string]uint32
}

func (labeledUgraph LabeledUgraph) Adj(vertex string) []string {
	var adj []string
	return labeledUgraph.CollectAdjTo(vertex, adj)

}

func (labeledUgraph LabeledUgraph) CollectAdjTo(vertex string, collector []string) []string  {
	v := labeledUgraph.invertedIndex[vertex]
	for _, w := range labeledUgraph.g.adjList[v] {
		collector = append(collector, labeledUgraph.labels[w])
	}
	return collector
}

func (labeledUgraph LabeledUgraph) Ugraph() *Ugraph {
	return labeledUgraph.g
}

func (labeledUgraph LabeledUgraph) NameFor(v uint32) string {
	return labeledUgraph.labels[v]
}

func NewDiGraph(n uint32) *Digraph  {
	return &Digraph{n: n, adjList: make([][]Edge, n)};

}

func (g *Digraph) AddEdge(from, to uint32, wt int32) {
	g.adjList[from] = append(g.adjList[from], Edge{from:from, to:to, wt: wt});
	g.edgeCount++
}

func (g *Digraph) EdgeCount() uint32 {
	return g.edgeCount;
}

type Ugraph struct {
	adjList [][]uint32
	edgeCount uint32
}

func NewUGraph(n uint32) *Ugraph {
	return &Ugraph{adjList: make([][]uint32, n)};
}

func (g *Ugraph) AddEdge(u, v uint32) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
	g.edgeCount++
}

func (g *Ugraph) Edges() uint32 {
	return g.edgeCount;
}

func (g *Ugraph) Vertices() int{
	return len(g.adjList)
}

func (g *Ugraph) Adj(v uint32) []uint32 {
	return g.adjList[v]
}


func ReadUgraph(delimiter string, file *os.File) *LabeledUgraph  {
	scanner := bufio.NewScanner(file)
	vertexIdMap := make(map[string] uint32)
	adjListMap := make(map[uint32] []uint32)
	var id uint32
	for scanner.Scan() {
		line := scanner.Text()
		connections := strings.Split(line, delimiter)
		vertex := connections[0]
		for i, v := range connections {
			if _, exists := vertexIdMap[v]; !exists {
				vertexIdMap[v] = id
				id++
			}
			if i == 0 {continue}
			u := vertexIdMap[vertex]
			adjListMap[u] = append(adjListMap[u], vertexIdMap[v])
		}
	}
	numberOfVertices := len(vertexIdMap)
	g := NewUGraph(uint32(numberOfVertices))
	for u, adjList := range adjListMap {
		for _, v := range adjList {
			g.AddEdge(u, v)
		}
	}
	return NewLabeledGraph(vertexIdMap, g)
}

func NewLabeledGraph(labelMap map[string]uint32, ugraph *Ugraph) *LabeledUgraph {
	index := make(map[uint32]string)
	for k, v := range labelMap {
		index[v] = k
	}
	return &LabeledUgraph{invertedIndex: labelMap, g:ugraph, labels: index}
}