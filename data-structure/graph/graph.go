package main

import "fmt"

type Graph struct {
	vertices map[int][]int
}

func NewGraph() *Graph {
	return &Graph{make(map[int][]int)}
}
func (g *Graph) AddVertex(value int) {
	if _, exist := g.vertices[value]; !exist {
		g.vertices[value] = []int{}
	}
}
func (g *Graph) AddEdges(v1, v2 int) {
	if _, exist := g.vertices[v1]; exist {
		g.vertices[v1] = append(g.vertices[v1], v2)
	}
	if _, exist := g.vertices[v2]; exist {
		g.vertices[v2] = append(g.vertices[v2], v1)
	}
}
func main() {
	graph := NewGraph()
	graph.AddVertex(4)
	graph.AddVertex(7)
	graph.AddVertex(8)
	graph.AddVertex(2)
	graph.AddEdges(4, 8)
	graph.AddEdges(4, 2)
	graph.AddEdges(2, 7)
	// graph.PrintGraph()
	graph.DFS(4, make(map[int]bool))
	graph.BFS(4)

}
func (g *Graph) PrintGraph() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%d -> %v\n", vertex, edges)
	}
}
func (g *Graph) DFS(start int, visited map[int]bool) {
	if _, exist := g.vertices[start]; !exist {
		fmt.Println("vertices not found")
		return
	}
	visited[start] = true
	fmt.Println(start)
	for _, neighbor := range g.vertices[start] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
		}
	}
}
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		fmt.Println(vertex)
		for _, neighbor := range g.vertices[vertex] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

}
