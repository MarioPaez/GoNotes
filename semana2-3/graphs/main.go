package main

import (
	"fmt"
	"slices"
	"sort"
)

type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	Vertex map[int][]Edge
}

func (g *Graph) AddVertex(value int) {
	if _, exists := g.Vertex[value]; exists {
		fmt.Printf("Ya existe un vértice con value: %d\n", value)
		return
	}
	g.Vertex[value] = []Edge{}
}

func (g *Graph) checkIfVertexExist(from, to int) bool {
	if g.Vertex[from] == nil {
		fmt.Printf("No se puede conectar el vertice %d(from) con el vértice %d (to) porque no existe el vértice %d.\n", from, to, from)
		return false
	}
	if g.Vertex[to] == nil {
		fmt.Printf("No se puede conectar el vertice %d(from) con el vértice %d (to) porque no existe el vértice %d.\n", from, to, to)
		return false
	}
	return true
}

func (g *Graph) AddEdge(from, to, weight int) {
	if g.checkIfVertexExist(from, to) {
		g.Vertex[from] = append(g.Vertex[from], Edge{To: to, Weight: weight})
		fmt.Printf("Se ha conectado el vértice %d con el vértice %d con un peso de %d\n", from, to, weight)
	}
}

func (g *Graph) GetNeightbours(val int) []Edge {
	vertice := g.Vertex[val]

	if vertice == nil {
		fmt.Println("No existe el vértice con value ", val)
		return nil
	}
	return vertice
}

func (g *Graph) GetVertices() []int {
	res := make([]int, 0)
	mapa := g.Vertex

	for i := range mapa {
		res = append(res, i)
	}
	sort.Ints(res)
	return res
}

func (g *Graph) RemoveEdge(from, to int) {
	if g.checkIfVertexExist(from, to) {
		elements := g.Vertex[from]
		for i, edge := range elements {
			if edge.To == to {
				g.Vertex[from] = slices.Delete(elements, i, i+1)
				break
			}
		}
	}
}

func (g *Graph) RemoveVertex(val int) {
	for from := range g.Vertex {
		if from == val {
			continue
		}
		g.RemoveEdge(from, val)
	}
	delete(g.Vertex, val)
}

func (g *Graph) GetGraph() {
	fmt.Println(g.Vertex)
}

func (g *Graph) BFS() {
	isVisited := make(map[int]bool)

	for _, start := range g.GetVertices() {
		if !isVisited[start] {
			queue := []int{start}
			isVisited[start] = true

			for len(queue) > 0 {
				vertex := queue[0]
				queue = queue[1:]

				fmt.Println("Visitando:", vertex)

				for _, edge := range g.Vertex[vertex] {
					if !isVisited[edge.To] {
						isVisited[edge.To] = true
						queue = append(queue, edge.To)
					}
				}
			}
		}
	}
}

func (g *Graph) DFS() {
	isVisited := make(map[int]bool)
	for _, vertex := range g.GetVertices() {
		if !isVisited[vertex] {
			g.auxDFS(vertex, isVisited)
		}
	}
}

func (g *Graph) auxDFS(current int, isVisited map[int]bool) {
	if isVisited[current] {
		return
	}

	fmt.Println("He visitado el vértice", current)
	isVisited[current] = true

	for _, edge := range g.Vertex[current] {
		if !isVisited[edge.To] {
			g.auxDFS(edge.To, isVisited)
		}
	}
}

func main() {
	mapa := make(map[int][]Edge)
	grafo := Graph{mapa}

	grafo.AddVertex(1)
	grafo.AddVertex(2)
	grafo.AddVertex(3)
	grafo.AddVertex(4)
	grafo.AddVertex(5)
	grafo.AddVertex(6)
	grafo.AddVertex(9)
	grafo.AddVertex(9)

	grafo.AddEdge(1, 4, 3)
	grafo.AddEdge(1, 2, 1)
	grafo.AddEdge(2, 3, 1)

	grafo.GetGraph()

	fmt.Println("Vecinos del vertice 1: ", grafo.GetNeightbours(1))
	fmt.Println("Vecinos del vertice 2: ", grafo.GetNeightbours(2))

	fmt.Println(grafo.GetVertices())
	println("RECORRIDO BFS")
	grafo.BFS()
	println("RECORRIDO DFS")
	grafo.DFS()
	fmt.Println(grafo.GetNeightbours(1))

	grafo.RemoveVertex(2)
	fmt.Println("Vecinos del vertice 2: ", grafo.GetNeightbours(2))
	fmt.Println("Vecinos del vertice 1: ", grafo.GetNeightbours(1))

}
