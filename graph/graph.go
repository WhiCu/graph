package graph

import (
	"graph/node"
)

type PrimeGraph[T any] struct {
	nodes []node.Node[T]
}

func New[T any](factory node.Factory[T], init T, count int) PrimeGraph[T] {
	var pg PrimeGraph[T]
	pg.nodes = make([]node.Node[T], count)

	for i := 0; i < count; i++ {
		pg.nodes[i] = factory.Node(init)
		pg.nodes[i].SetNumber(i + 1)
	}
	return pg
}
func (g PrimeGraph[T]) SetAdjacency(first, second int) {
	g.nodes[first-1].SetAdjacent(g.nodes[second-1])
}
func (g PrimeGraph[T]) AdjacencyMatrix() [][]byte {
	adjacencyMatrix := make([][]byte, len(g.nodes))
	for i := 0; i < len(adjacencyMatrix); i++ {
		adjacencyMatrix[i] = make([]byte, len(g.nodes))
		for _, v := range g.nodes[i].AdjacentNodes() {
			adjacencyMatrix[i][v.Number()-1] = 1
		}
	}
	return adjacencyMatrix
}
