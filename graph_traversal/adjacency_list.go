package graph_traversal

import "fmt"

const maxv = 10 // maxamount of vertices for a graph

type edgenode struct {
	y      int
	weight int
	next   *edgenode
}

type graph struct {
	edges     [maxv]*edgenode //
	degree    [maxv]int       // number of vertices connected to vertex i
	nvertices int
	nedges    int
	directed  bool
}

func NewGraph(directed bool) *graph {
	return &graph{
		directed: directed,
	}
}

// inserts the edge from x to y
func (g *graph) InsertEdge(x, y int) {
	var p *edgenode

	p = &edgenode{}

	p.y = y

	// set the next edgenode that is connected to p
	// as next
	// g.edges[x] = nil when no edgenode has instantiated yet
	p.next = g.edges[x]

	// set p as the new head of the linked list
	g.edges[x] = p

	// increase the degree of the xth node
	g.degree[x]++

	if !g.directed {
		// insert the edge from y to x
		g.InsertEdge(y, x)
	} else {
		g.nedges++
	}
}

func (g *graph) Bfs(start int) {

}

func (g *graph) Print() {
	init_graph_msg := fmt.Sprintf(
		"number of vertices: %v \nnumber of edges: %v \ndirected: %v",
		g.nvertices,
		g.nedges,
		g.directed,
	)

	for vertex, adjacency_list := range g.edges {
		if g.edges[vertex] == nil {
			// do nothing
		} else {
			init_graph_msg = fmt.Sprint(
				init_graph_msg,
				fmt.Sprintf(
					"\n\nvertex: %v\nvertex degree: %v",
					vertex,
					g.degree[vertex],
				),
			)
			if g.degree[vertex] == 0 {
				// do nothing
			} else {
				// get all edgenodes connected to vertex
				edgenodes := []int{}
				edgenode := adjacency_list
				for edgenode != nil {
					edgenodes = append(edgenodes, edgenode.y)
					edgenode = edgenode.next
				}

				// format adjacency list edgenodes
				adjacency_list_msg := fmt.Sprintf("\n\t%v -> ", vertex)

				for _, edgenode := range edgenodes {
					adjacency_list_msg = fmt.Sprint(adjacency_list_msg, edgenode, " -> ")
				}

				adjacency_list_msg = fmt.Sprint(adjacency_list_msg, "nil")

				init_graph_msg = fmt.Sprint(init_graph_msg, adjacency_list_msg)
			}
		}
	}

	fmt.Print(init_graph_msg, "\n")
}
