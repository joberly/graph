package graph

import "testing"

func TestEdges(t *testing.T) {
	g := NewGraph()

	a := g.NewVertex("a")
	d1 := g.NewVertex("d1")
	d2 := g.NewVertex("d2")
	d3 := g.NewVertex("d3")

	b := g.NewVertex("b")

	edgeinfos := []struct {
		v *Vertex
		c int64
	}{
		{v: d1, c: 1},
		{v: d2, c: 2},
		{v: d3, c: 3},
	}

	// Add edges to destinations
	for _, edgeinfo := range edgeinfos {
		a.AddEdge(edgeinfo.v, edgeinfo.c)
	}

	// Check cost to destinations
	for _, edgeinfo := range edgeinfos {
		c := a.FindCost(edgeinfo.v)
		if c == nil {
			t.Errorf("FindCost from %v returned unexpected nil cost for %v", *a, *edgeinfo.v)
		}
		if c.(int64) != edgeinfo.c {
			t.Errorf("FindCost returned unexpected cost value %v for dest with cost %v", c.(int), edgeinfo.c)
		}
	}

	// Remove destinations and check for nil cost
	for _, edgeinfo := range edgeinfos {
		cold := a.RemoveEdge(edgeinfo.v)
		if cold == nil {
			t.Errorf("RemoveEdge from %v returned unexpected nil cost for %v", *a, *edgeinfo.v)
		} else if coldi64 := cold.(int64); coldi64 != edgeinfo.c {
			t.Errorf("RemoveEdge from %v returned unexpected cost %v", coldi64, edgeinfo.c)
		}
		if c := a.FindCost(edgeinfo.v); c != nil {
			t.Errorf("FindCost returned non-nil cost %v for removed edge %v", c, *edgeinfo.v)
		}
	}

	// Check edge that never existed
	if c := a.FindCost(b); c != nil {
		t.Errorf("FindCost returned non-nil cost %v for non-existent edge", c)
	}

	// Try to remove non-existent edge. Should not panic.
	cnil := a.RemoveEdge(b)
	if cnil != nil {
		t.Errorf("RemoveEdge returned non-nil cost %v for non-existent edge", cnil)
	}

	// Add edges to a vertex then remove the vertex
	for _, edgeinfo := range edgeinfos {
		edgeinfo.v.AddEdge(b, edgeinfo.c)
	}
	g.RemoveVertex(b)

	// Now check the destinations and ensure no edges to b.
	for _, edgeinfo := range edgeinfos {
		if c := edgeinfo.v.FindCost(b); c != nil {
			t.Errorf("FindCost returned non-nil cost %v to removed vertex from %v", c, *edgeinfo.v)
		}
	}
}
