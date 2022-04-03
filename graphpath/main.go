package main

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

func main() {
	g := simple.NewUndirectedGraph()
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
	fmt.Print(g.Edges())

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

	// output : dot file
	result, _ := dot.Marshal(g, "", "", "  ")
	fmt.Print(string(result))
	file, err := os.Create("./graph.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(result)
	if err != nil {
		log.Fatal(err)
	}

	// Graphviz
	// dot -Tsvg graph.dot -o output.svg
	// dot -Tpng graph.dot -o output.png
}
