# graph
Graph package for Go.

[![GoDoc](https://godoc.org/github.com/joberly/graph?status.svg)](https://godoc.org/github.com/joberly/graph)
[![Build Status](https://travis-ci.org/joberly/graph.svg?branch=master)](https://travis-ci.org/joberly/graph)

## Graph Design

### Graph

A Graph is a simple map of Vertexes. This really just keeps a reference to all
Vertexes in the graph until removed with RemoveVertex. This could index by some
Vertex key, but there hasn't been a need for it yet and the intention is to
keep Graph generic. This leaves indexing a Graph's Vertexes by their keys to
the specific consumer of Graph.

### Vertex

Each Vertex contains two maps which map Vertex pointers to Edges. One maps the
destination vertex to each outgoing edge from the Vertex directed to another 
Vertex. The other maps the source vertex to each incoming edge to the Vertex 
directed from another Vertex. The Value of a Vertex is an empty interface used 
solely by the consumer of the Vertex.

### Edge

An Edge contains pointers to the source and destination Vertexes for that Edge.
It also contains a Cost which designates the cost of traversing the edge. This
is generally intended to be used only by the consumer with the exception of
using FindShortestPaths which requires that the Cost of each Edge be of type
int64.

## FindShortestPaths (Dijkstra's Algorithm)

The FindShortestPaths method implements Dijkstra's Algorithm to find the
shortest paths to all Vertexes in the Graph from the given source vertex. All
Edges in the Graph must have a Cost of type int64 or else the method will
panic. The result is a map of destination vertexes to PathInfo structures.

### Path

A Path struct contains path information to a Vertex from the source Vertex
specified in the call to FindShortestPaths. All Paths from a call to
FindShortestPaths are interconnected. All the Vertexes on the shortest path
from a Vertex back to the original source Vertex can be found by following
the ```prev``` pointers from Path to Path until reaching the source Vertex.
