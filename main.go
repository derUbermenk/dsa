package main

import (
	"dsa/graph_traversal"
	"fmt"
)

func main() {

}

func GraphTraverse() {
	g := graph_traversal.NewGraph(true)

	g.InsertEdge(0, 1)
	g.InsertEdge(1, 3)
	g.InsertEdge(1, 4)
	g.InsertEdge(2, 0)
	g.InsertEdge(2, 5)

	fmt.Printf("%+v", g)
	g.Print()

	parents := g.Bfs(2)
	fmt.Printf("\n")
	graph_traversal.FindPath(parents, 2, 4)
}
