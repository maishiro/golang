package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

func main() {
	g := simple.NewDirectedGraph()
	p0 := simple.Node(int64(0))
	p1 := simple.Node(int64(1))
	p2 := simple.Node(int64(2))
	p3 := simple.Node(int64(3))
	p4 := simple.Node(int64(4))
	g.AddNode(p0)
	g.AddNode(p1)
	g.AddNode(p2)
	g.AddNode(p3)
	g.AddNode(p4)
	g.SetEdge(simple.Edge{F: p1, T: p2})
	g.SetEdge(simple.Edge{F: p2, T: p3})
	g.SetEdge(simple.Edge{F: p2, T: p4})
	fmt.Println(g.Edges())
	fmt.Println()

	if topo.PathExistsIn(g, p0, p1) {
		fmt.Println("p0 - p1 : path exist")
	} else {
		fmt.Println("p0 - p1 : path none")
	}
	if topo.PathExistsIn(g, p1, p4) {
		fmt.Println("p1 - p4 : path exist")
	} else {
		fmt.Println("p1 - p4 : path none")
	}
	if topo.PathExistsIn(g, p1, p2) {
		fmt.Println("p1 - p2 : path exist")
	} else {
		fmt.Println("p1 - p2 : path none")
	}
	fmt.Println()

	// output : dot file
	WriteGraph(g, "./graph.dot")
	// Graphviz
	// dot -Tsvg graph.dot -o output.svg
	// dot -Tpng graph.dot -o output.png

	// ツリーをリストに並べる
	PrintWalkAll(g)

	var connected []int64
	var leaf []int64

	nodes := g.Nodes()
	// fmt.Println("len:", nodes.Len())
	for nodes.Next() {
		n := nodes.Node()
		if n.ID() == p1.ID() {
			continue
		}

		// p1からつながりがある
		if topo.PathExistsIn(g, p1, n) {
			connected = append(connected, n.ID())

			// かつ 端点
			if g.From(n.ID()).Len() == 0 {
				leaf = append(leaf, n.ID())
			}
		}
	}

	// p1からつながりがある
	fmt.Print("connected: ")
	for _, v := range connected {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")

	// p1からつながりのある端点
	fmt.Print("leaf: ")
	for _, v := range leaf {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}
