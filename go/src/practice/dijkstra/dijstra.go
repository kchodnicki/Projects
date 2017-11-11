package main

import (
	"fmt"
	"math"
)

type edge struct {
	start, end, weight int
}

type graph struct {
	V, E  int
	edges []*edge
}

func makeGraph(vNum int, pEdges []*edge) *graph {
	return &graph{V: vNum, E: len(pEdges), edges: pEdges}
}

func shortPath(root int, pGraph *graph) []int {
	sptSet, distanceSet := []bool{}, []int{}
	for i := 0; i < pGraph.V; i++ {
		distanceSet = append(distanceSet, math.MaxUint32)
		sptSet = append(sptSet, false)
	}
	distanceSet[root] = 0
	for cnt := 0; cnt < pGraph.V-1; cnt++ {
		minDistV := findShortestDist(distanceSet, sptSet)
		sptSet[minDistV] = true
		for _, edgeFromV := range getEdgesFromV(minDistV, pGraph, sptSet) {
			updateDistance(minDistV, edgeFromV, distanceSet)
		}
	}
	return distanceSet
}

func findShortestDist(pDistanceSet []int, pVisitedSet []bool) int {
	minDistV := 0
	minDist := math.MaxUint32
	for i := 0; i < len(pDistanceSet); i++ {
		if !pVisitedSet[i] {
			if pDistanceSet[i] < minDist {
				minDist = pDistanceSet[i]
				minDistV = i
			}
		}
	}
	return minDistV
}

func updateDistance(root int, edgeFromV *edge, pDistanceSet []int) {
	newDist := pDistanceSet[root] + edgeFromV.weight
	if edgeFromV.start == root {
		if newDist < pDistanceSet[edgeFromV.end] {
			pDistanceSet[edgeFromV.end] = newDist
		}
	} else {
		if newDist < pDistanceSet[edgeFromV.start] {
			pDistanceSet[edgeFromV.start] = newDist
		}
	}
}

func getEdgesFromV(root int, pGraph *graph, pVisitedSet []bool) []*edge {
	edgesFromRoot := []*edge{}
	for _, edge := range pGraph.edges {
		if edge.start == root || edge.end == root {
			edgesFromRoot = append(edgesFromRoot, edge)
		}
	}
	return edgesFromRoot
}

func main() {
	var edges []*edge
	edges = append(edges, &edge{start: 0, end: 1, weight: 4})
	edges = append(edges, &edge{start: 0, end: 7, weight: 8})
	edges = append(edges, &edge{start: 1, end: 2, weight: 8})
	edges = append(edges, &edge{start: 1, end: 7, weight: 11})
	edges = append(edges, &edge{start: 2, end: 3, weight: 7})
	edges = append(edges, &edge{start: 2, end: 8, weight: 2})
	edges = append(edges, &edge{start: 2, end: 5, weight: 4})
	edges = append(edges, &edge{start: 3, end: 4, weight: 9})
	edges = append(edges, &edge{start: 3, end: 5, weight: 14})
	edges = append(edges, &edge{start: 4, end: 5, weight: 10})
	edges = append(edges, &edge{start: 5, end: 6, weight: 2})
	edges = append(edges, &edge{start: 6, end: 8, weight: 6})
	edges = append(edges, &edge{start: 6, end: 7, weight: 1})
	edges = append(edges, &edge{start: 7, end: 8, weight: 7})
	newGraph := makeGraph(9, edges)
	fmt.Println(shortPath(edges[0].start, newGraph))
}
