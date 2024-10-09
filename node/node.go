package node

type Node[T any] interface {
	//Получить номер
	Number() int
	//Задать номер
	SetNumber(int)
	//Получить значение содержащиесе в узле
	Value() T

	//Задать значение содержащиесе в узле
	SetValue(value T)

	//Добавить смежную вершину
	SetAdjacent(Node[T])

	//Получить смежные вершины в slice
	AdjacentNodes() []Node[T]
}

type Factory[T any] interface {
	//Создаёт Node
	Node(init T) Node[T]
}
