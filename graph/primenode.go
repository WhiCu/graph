package graph

import (
	"graph/node"
)

type PrimeNode[T any] struct {
	number     int
	value      T
	primeNodes []node.Node[T]
}

func (pn *PrimeNode[T]) Number() int {
	return pn.number
}
func (pn *PrimeNode[T]) SetNumber(n int) {
	pn.number = n
}
func (pn *PrimeNode[T]) SetAdjacent(node node.Node[T]) {
	pn.primeNodes = append(pn.primeNodes, node)
}
func (pn *PrimeNode[T]) Value() T {
	return pn.value
}

func (pn *PrimeNode[T]) SetValue(value T) {
	pn.value = value
}

func (pn *PrimeNode[T]) AdjacentNodes() []node.Node[T] {
	var adjacentNodes = make([]node.Node[T], len(pn.primeNodes))
	for i := 0; i < len(adjacentNodes); i++ {
		adjacentNodes[i] = pn.primeNodes[i]
	}
	return adjacentNodes
}

func NewPrimeNode[T any](init T) node.Node[T] {
	return &PrimeNode[T]{
		value:      init,
		primeNodes: make([]node.Node[T], 0),
	}
}

//factory

type PrimeFactory[T any] func(init T) node.Node[T]

// Node creates a new node using the PrimeFactory.
func (pf PrimeFactory[T]) Node(init T) node.Node[T] {
	return pf(init)
}

func NewPrimeFactory[T any](init T) PrimeFactory[T] {
	var factory PrimeFactory[T] = NewPrimeNode[T]
	return factory

}
