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

/*
	for a graph instantiated this way
		g.InsertEdge(0, 1)
		g.InsertEdge(1, 3)
		g.InsertEdge(1, 4)
		g.InsertEdge(2, 0)
		g.InsertEdge(2, 5)

	using a start of 0 will not visit all vertices. while using 2 will.
*/
func (g *graph) Bfs(start int) (parents []int) {
	/*
		processing_que and discovered array

		discovery of a vertex assures us that it will not be processed anymore
		as a vertex will only be added to the processing que if it has not been processed
		and likewise will only be added to the processing que if it has been processed.
	*/

	// set initial parent values to -1
	parents = make([]int, maxv)
	for i := 0; i < len(parents); i++ {
		parents[i] = -1
	}

	var v int       // current vertex
	var y int       // successor vertex
	var p *edgenode // adjacency list of the current vertex

	discovered := make([]bool, maxv)
	processed := make([]bool, maxv)

	processing_que := make([]int, 0)
	// enque start
	processing_que = append(processing_que, start)
	discovered[start] = true

	for len(processing_que) != 0 {
		// deque first item in processing que
		v = processing_que[0]
		processing_que = processing_que[1:]

		// prints current vertex being processed
		process_vertex_early(v)

		// get the adjacency list of the current vertex (v)
		p = g.edges[v]
		for p != nil {
			y = p.y
			if !discovered[y] {
				processing_que = append(processing_que, y)
				discovered[y] = true
				parents[y] = v
			}
			/*
				if !processed[y] || g.directed {
					// do something
				}
			*/
			p = p.next
		}

		processed[v] = true
	}

	return parents
}

func FindPath(parents []int, start, end int) {
	if start == end || end == -1 {
		fmt.Printf("%v -> ", start)
	} else {
		FindPath(parents, start, parents[end])
		fmt.Printf("%v -> ", end)
	}
}

func process_vertex_early(v int) {
	fmt.Printf("%v -> ", v)
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
