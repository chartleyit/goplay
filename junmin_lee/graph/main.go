package main

import "fmt"

// Graph structure
type Graph struct {
	vertices []*Vertex
}

// Vertex structure

type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a Vertex to the Graph
func (g *Graph) AddVertex(k int) {
	g.vertices = append(g.vertices, &Vertex{key: k})
}

// Add Edge

// getVertex
// contains

// Print will print the adjacent list for each vertex of the grpah
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v", v.key)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddVertex(0)

	test.Print()
}
