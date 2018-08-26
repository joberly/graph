// Package graph contains data structures for directed graphs.
package graph

// Cost is an abstract type for the cost of an edge between two nodes.
type Cost interface{}

// Graph is a collection of nodes with their edges.
type Graph struct {
	vertex map[*Vertex]bool
}

// Edge is a connection from source to destination with an associated cost.
type edge struct {
	cost Cost
	src *Vertex // Source vertex
	dst *Vertex // Destination vertex
}

// Vertex is an element in a directed graph.
type Vertex struct {
	Value interface{}

	out map[*Vertex]*edge // Edges outgoing from this vertex (key is destination)
	in  map[*Vertex]*edge // Edges incoming to this vertex (key is source)
}

// NewGraph creates a new graph.
func NewGraph() *Graph {
	return &Graph{vertex: make(map[*Vertex]bool)}
}

// NewVertex creates a new node with the given value and places it in the Graph.
func (g *Graph) NewVertex(value interface{}) *Vertex {
	n := &Vertex{
		Value: value,
		in: make(map[*Vertex]*edge),
		out: make(map[*Vertex]*edge),
	}
	g.vertex[n] = true
	return n
}

// AddEdge adds a directed connection from a node to a destination node.
// If the node is already connected to the other node, the cost is
// updated. Note that FindShortestPaths expects all Costs to be of type
// int64.
func (v *Vertex) AddEdge(dst *Vertex, c Cost) {
	if dst == nil {
		panic("edge destination is nil")
	}
	if c == nil {
		panic("cost is nil")
	}

	if e, connected := v.out[dst]; connected {
		e.cost = c
		return
	}

	e := &edge{cost: c, src: v, dst: dst}
	e.src.out[e.dst] = e
	e.dst.in[e.src] = e

	return
}

// FindCost returns the cost from one vertex to another. If there is no edge
// to the destination vertex, FindCost returns nil.
func (v *Vertex) FindCost(dst *Vertex) Cost {
	e := v.out[dst]
	if e == nil {
		return nil
	}

	return e.cost
}

// RemoveEdge removes the connection from this node to the destination node.
// Returns the cost of the edge that is removed. If the edge did not exist,
// RemoveEdge returns nil.
func (v *Vertex) RemoveEdge(dst *Vertex) Cost {
	e := v.out[dst]
	if e == nil {
		return nil
	}

	delete(e.src.out, e.dst)
	delete(e.dst.in, e.src)
	return e.cost
}

// RemoveVertex removes a vertex from the graph. It removes all edges associated
// with the vertex in the process.
func (g *Graph) RemoveVertex(v *Vertex) {
	// Remove edges to this vertex from all sources.
	for src := range v.in {
		delete(src.out, v)
	}
	// Delete the node from graph's map of vertexes.
	delete(g.vertex, v)
}
