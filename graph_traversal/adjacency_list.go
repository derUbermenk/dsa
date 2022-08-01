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
	// disregard parent
	/*
		discovered bool -> arr with size vertex
		processed bool -> arr with size vertex

		processing_que -> que of vertex to process

		enque(processing_que, start)
		discovered[start] = TRUE
		while the processing_que is not empty	{
			current_vertex = deque(processing_que)
			adjacency_list_of_current_vertex = edges[current_vertex]
			while adjacency_list_of_current_vertex != null {
				// check status of successor vertex and act accordingly
				if !discovered[successor_vertex] {
					// discover successor_vertex
					discovered[successor_vertex] = true
					enque(processing_que, successor_vertex)
				}
				// ?? don't know this step
				if !processed[successor_vertex] || g.directed {
					process_edge(v, y)
				}
				// go to next
				adjacency_list = adjacency_list.next
			}
		}
		// set current_vertex as processed
		proccessed[current_vertex] = TRUE
	*/

	var v int       // current vertex
	var y int       // successor vertex
	var p *edgenode // adjacency list of the current vertex

	discovered := make([]bool, g.nvertices)
	processed := make([]bool, g.nvertices)

	processing_que := make([]int, 0)
	// enque start
	processing_que = append(processing_que, start)
	discovered[start] = true

	for len(processing_que) != 0 {
		// deque first item in processing que
		v = processing_que[0]
		processing_que = processing_que[1:]

		// get the adjacency list of the current vertex (v)
		p = g.edges[v]
		for p != nil {
			y = p.y
			if !discovered[y] {
				discovered[y] = true
				processing_que = append(processing_que, y)
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
