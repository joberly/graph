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
// updated.
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

// RemoveEdge removes the connection from this node to the destination node.
func (v *Vertex) RemoveEdge(dst *Vertex) {
	e := v.out[dst]
	if e == nil {
		return
	}

	delete(e.src.out, e.dst)
	delete(e.dst.in, e.src)
}

// RemoveNode removes a node from the graph. It removes all edges associated
// with the node in the process.
func (g *Graph) RemoveNode(v *Vertex) {
	// Remove edges to this node from all sources.
	for src := range v.in {
		delete(src.out, v)
	}
	// Delete the node from graph's map of nodes.
	delete(g.vertex, v)
}