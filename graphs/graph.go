package graphs

type Graph interface {
	AddEdge(vertex1, vertex2 string, weight ...int)
	RemoveEdge(vertex1, vertex2 string)
	GetNeighbors(vertex string) []string
	HasEdge(vertex1, vertex2 string) bool
	GetEdgeWeight(vertex1, vertex2 string) (int, bool)
}
