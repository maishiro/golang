package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

func WriteGraph(g graph.Graph, name string) {
	// output : dot file
	result, _ := dot.Marshal(g, "", "", "  ")
	fmt.Print(string(result), "\n\n")
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(result)
	if err != nil {
		log.Fatal(err)
	}
}

// ツリーをリストに並べる
func PrintWalkAll(g *simple.DirectedGraph) {
	// 上端を抽出
	var roots []int64
	{
		nodes := g.Nodes()
		for nodes.Next() {
			node := nodes.Node()
			// 上層がいないものを抽出
			if g.To(node.ID()).Len() == 0 {
				roots = append(roots, node.ID())
			}
		}
	}
	// 上端を起点にそれぞれ出力
	for _, root := range roots {
		fmt.Println("root: ", root)
		WalkAll(g, root)
	}
	fmt.Println()
}

func WalkAll(g graph.Graph, root int64) {
	node := g.Node(root)
	Walk(g, node, 0)
}

func Walk(g graph.Graph, node graph.Node, depth int) {
	fmt.Println(strings.Repeat(" ", depth), "*", node.ID())

	nodes := g.From(node.ID())
	for nodes.Next() {
		node := nodes.Node()
		Walk(g, node, depth+1)
	}
}
