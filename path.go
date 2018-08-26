package graph

import (
	"math"

	heap "github.com/joberly/heap/binomial"
)

// PathInfo contains information about the shortest path from the source
// vertex specified when calling FindShortestPaths to the vertex in the
// PathInfo struct.
type PathInfo struct {
	vertex  *Vertex
	prev    *PathInfo
	dist    int64
}

// ShortestPathMap maps PathInfo for the shortest path 
type ShortestPathMap map[*Vertex]*PathInfo

// FindShortestPaths finds the shortest paths to other vertexes in a graph
// staritng from a single source vertex. FindShortestPaths expects all 
// edge costs to be of type int64.
func FindShortestPaths(g *Graph, src *Vertex) ShortestPathMap {
	q := &heap.Heap{}

	numv := 0
	pm := make(map[*Vertex]*PathInfo)
	for v := range g.vertex {
		pi := &PathInfo{vertex: v, dist: math.MaxInt64}
		pm[v] = pi
		numv++
	}
	pisrc := pm[src]
	pisrc.dist = 0

	vpi := pisrc
	for vpi != nil {
		for w, e := range vpi.vertex.out {
			wpi := pm[w]
			wdist := vpi.dist + e.cost.(int64)
			if wdist < wpi.dist {
				wpi.dist = wdist
				wpi.prev = vpi
				q.Add(wpi, pathLess)
			}
		}
		temp := q.RemoveMin(pathLess)
		if temp != nil {
			vpi = temp.(*PathInfo)
		} else {
			vpi = nil
		}
	}

	return pm
}

func pathLess(a, b interface{}) bool {
	return a.(*PathInfo).dist < b.(*PathInfo).dist
}
