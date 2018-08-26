package graph

import "testing"

func TestFindShortestPaths(t *testing.T) {
	g := NewGraph()

	a := g.NewVertex("a")
	b := g.NewVertex("b")
	c := g.NewVertex("c")
	d := g.NewVertex("d")

	a.AddEdge(b, int64(1))
	a.AddEdge(c, int64(2))

	b.AddEdge(a, int64(1))
	b.AddEdge(c, int64(3))
	b.AddEdge(d, int64(1))

	c.AddEdge(a, int64(1))
	c.AddEdge(b, int64(3))
	c.AddEdge(d, int64(3))

	d.AddEdge(b, int64(1))
	d.AddEdge(c, int64(1))

	// Find the shortest paths to vertex a
	pm := g.FindShortestPaths(a)

	// Check that the PathInfo for d is valid
	pathToD := pm[d]
	if pathToD == nil {
		t.Fatalf("no shortest path info to vertex d\n")
	}
	if pathToD.dist != 2 {
		t.Fatalf("shortest path dist to d incorrect (was %d)\n", pathToD.dist)
	}

	// Check that the path back to d is valid
	vlist := []*Vertex{d, b, a}
	pi := pm[vlist[0]]
	for i, v := range vlist {
		if pi == nil {
			t.Fatalf("Path for vertex %s on shortest path is nil\n", v.Value.(string))
		}
		if pi.vertex != v {
			t.Fatalf("Path for vertex %s does not map to correct vertex\n", v.Value.(string))
		}
		// Check source vertex's
		if i == len(vlist)-1 {
			if pi.prev != nil {
				t.Errorf("Path of source vertex does not end chain\n")
			}
			if pi.dist != 0 {
				t.Errorf("Path of source vertex does not have distance 0 (actual %d)\n", pi.dist)
			}
		} else {
			if pi.prev == nil {
				t.Errorf("Path for vertex %s does not link to previous vertex\n", v.Value.(string))
			}
			pi = pi.prev
		}
	}

}
