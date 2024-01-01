package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

type edge struct {
	vertex   string
	capacity int
}

type graph struct {
	adjacencyList map[string][]*edge
	vertexParent  map[string]string
}

func initGraph() *graph {
	var g *graph = &graph{}
	g.adjacencyList = map[string][]*edge{}
	g.vertexParent = map[string]string{}
	return g
}

func (g *graph) bfs(source string, sink string) bool {
	g.vertexParent = map[string]string{}
	g.vertexParent[source] = source
	var q *vertexQueue = InitVertexQueue()
	q.enqueue(source)
	for !q.isEmpty() {
		var dequeuedVertex string = q.dequeue()
		for _, e := range g.adjacencyList[dequeuedVertex] {
			_, parentDefined := g.vertexParent[e.vertex]
			// if we have capacity and the parent hasn't been defined
			// then we define the parent for that vertext
			if e.capacity > 0 && !parentDefined {
				g.vertexParent[e.vertex] = dequeuedVertex
				q.enqueue(e.vertex)
			}
		}
	}
	_, parentOfSinkDefined := g.vertexParent[sink]
	return parentOfSinkDefined
}

func (g graph) getFlow(source string, sink string) int {
	var flow float64 = math.Inf(1)
	var currNode string = sink
	for currNode != source {
		var parent string = g.vertexParent[currNode]
		var parentEdges []*edge = g.adjacencyList[parent]
		for _, e := range parentEdges {
			if e.vertex == currNode {
				flow = math.Min(flow, float64(e.capacity))
				break
			}
		}
		currNode = parent
	}
	return int(flow)
}

func (g *graph) updateCapacityWithFlow(source, sink string, flow int) {
	var currNode string = sink
	for currNode != source {
		var parent string = g.vertexParent[currNode]
		var parentEdges []*edge = g.adjacencyList[parent]
		// Update capacity going FROM the parent
		for _, edge := range parentEdges {
			if edge.vertex == currNode {
				edge.capacity -= flow
				break
			}
		}
		// Update capacity going TO the parent
		var currNodeEdges []*edge = g.adjacencyList[currNode]
		for _, edge := range currNodeEdges {
			if edge.vertex == parent {
				edge.capacity += flow
				break
			}
		}
		currNode = parent
	}
}

func (g *graph) minCut(source string, sink string) int {
	var maxFlow int = 0
	for g.bfs(source, sink) {
		var flow int = g.getFlow(source, sink)
		maxFlow += flow
		g.updateCapacityWithFlow(source, sink, flow)

	}
	return maxFlow
}

func (g *graph) resetEdges() {
	for _, edges := range g.adjacencyList {
		for _, e := range edges {
			e.capacity = 1
		}
	}
}

func (g *graph) addConnection(labelA string, labelB string) {
	g.addVertexAndEdge(labelA, labelB)
	g.addVertexAndEdge(labelB, labelA)
}

func (g *graph) addVertexAndEdge(labelA string, labelB string) {
	currentConnections, found := g.adjacencyList[labelA]
	if !found {
		currentConnections = []*edge{}
	}
	var newEdge *edge = &edge{vertex: labelB, capacity: 1}
	if !slices.Contains(currentConnections, newEdge) {
		currentConnections = append(currentConnections, newEdge)
	}
	g.adjacencyList[labelA] = currentConnections
}

func multiplyTwoGraphSizes(g *graph) int {
	var graphOneCount int = 0
	for key := range g.adjacencyList {
		_, hasParent := g.vertexParent[key]
		if hasParent {
			graphOneCount++
		}
	}
	return (len(g.adjacencyList) - graphOneCount) * graphOneCount
}

func parseInput(input utils.InputFile) *graph {
	var g *graph = initGraph()
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		var splits []string = strings.Split(line, ":")
		var key string = splits[0]
		var vals []string = strings.Split(splits[1][1:], " ")
		for _, val := range vals {
			g.addConnection(val, key)
		}
	}
	return g
}

func main() {
	input := utils.InitInputFile("wiringData.txt")
	defer input.Close()
	var g *graph = parseInput(input)

	i := 0
	var source string
	for sink := range g.adjacencyList {
		i++
		// Compare the first in the loop against the rest
		if i == 1 {
			source = sink
			continue
		}
		g.resetEdges()
		if g.minCut(source, sink) == 3 {
			break
		}
	}

	var result int = multiplyTwoGraphSizes(g)
	fmt.Println("Final day (!) result: ", result)
}
