package main

import (
	"fmt"
	"graph/graph"
)

func main() {
	factory := graph.NewPrimeFactory(0)
	g := graph.New(factory, 0, 8)
	g.SetAdjacency(1, 1)
	g.SetAdjacency(2, 3)
	for _, v := range g.AdjacencyMatrix() {
		fmt.Println(v)
	}
	// fmt.Printf("is node %v - %T", graph, graph)
}
