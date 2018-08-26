package graph

import (
	"math"

	heap "github.com/joberly/heap/binomial"
)

// Path contains information about the shortest path from the source
// vertex specified when calling FindShortestPaths to the vertex in the
// Path struct. The prev member of Path points to the Path for the
// previous Vertex in the shortest path from the Vertex originally
// used when calling FindShortestPaths. The shortest path to the original
// source Vertex is found by following the prev Path pointers from
// the destination's Vertex's Path back to the original source Vertex.
type Path struct {
	vertex *Vertex // destination Vertex for this Path
	prev   *Path   // next Path back to source Vertex
	dist   int64   // distance to this Vertex from the source
}

// ShortestPathMap maps destination Vertexes to Paths for the shortest path
// from a specific source Vertex.
type ShortestPathMap map[*Vertex]*Path

// FindShortestPaths finds the shortest paths to other vertexes in a graph
// staritng from a single source vertex. FindShortestPaths expects all
// edge costs to be of type int64.
func (g *Graph) FindShortestPaths(src *Vertex) ShortestPathMap {
	// Use a heap as a priority queue for paths
	q := heap.NewHeap(pathLess)

	numv := 0
	pm := make(map[*Vertex]*Path)
	// Path to each Vertex is initially infinite
	for v := range g.vertex {
		p := &Path{vertex: v, dist: math.MaxInt64}
		pm[v] = p
		numv++
	}
	// Path to source from source has zero distance
	psrc := pm[src]
	psrc.dist = 0

	// Start at the source
	vp := psrc
	// Keep looping until there are no more in the queue.
	for vp != nil {
		// For each outgoing edge e to vertex w, add w's Path to the queue if
		// its distance is less than the current distance vp's Vertex plus
		// the distance to w.
		for w, e := range vp.vertex.out {
			wp := pm[w]
			wdist := vp.dist + e.cost.(int64)
			if wdist < wp.dist {
				wp.dist = wdist
				wp.prev = vp
				q.Add(wp)
			}
		}
		// Remove the next minimum Vertex's Path from queue
		// to go through its destinations.
		temp := q.RemoveMin()
		if temp != nil {
			vp = temp.(*Path)
		} else {
			vp = nil
		}
	}

	return pm
}

func pathLess(a, b interface{}) bool {
	return a.(*Path).dist < b.(*Path).dist
}
