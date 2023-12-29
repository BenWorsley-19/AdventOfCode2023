package main

import (
	"math"
	"slices"
)

type nodePointer struct {
	nodeKey        point
	distanceToNode int
}

type graph struct {
	start         point
	end           point
	adjacencyList map[point][]nodePointer
}

func InitGraph(start, end point, graphPoints []point, grid [][]rune, tileDirs map[rune][]direction) graph {
	var g graph = graph{}
	g.start = start
	g.end = end
	g.adjacencyList = map[point][]nodePointer{}
	for _, p := range graphPoints {
		g.adjacencyList[p] = []nodePointer{}
	}

	for _, from := range graphPoints {
		var nodeP nodePointer = nodePointer{nodeKey: from, distanceToNode: 0}
		var stack []nodePointer = []nodePointer{nodeP}
		var seen []point = []point{from}
		for len(stack) > 0 {
			var toNode nodePointer = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if toNode.distanceToNode != 0 && slices.Contains(graphPoints, toNode.nodeKey) {
				g.adjacencyList[from] = append(g.adjacencyList[from], toNode)
				continue
			}

			for _, dir := range tileDirs[grid[toNode.nodeKey.x][toNode.nodeKey.y]] {
				var np point = dir.nextPoint(toNode.nodeKey)
				if !pointIsOffGrid(np, grid) && grid[np.x][np.y] != '#' && !slices.Contains(seen, np) {
					stack = append(stack, nodePointer{nodeKey: np, distanceToNode: toNode.distanceToNode + 1})
					seen = append(seen, np)
				}
			}
		}
	}

	return g
}

func (g graph) dfs() int {
	longestPath, _ := walk(g, []point{}, g.start)
	return longestPath
}

func walk(g graph, seen []point, curr point) (int, bool) {
	if curr == g.end {
		return 0, true
	}
	seen = append(seen, curr)
	var longestPath int = int(math.Inf(-1))
	for _, nodePointer := range g.adjacencyList[curr] {
		var nextPoint point = nodePointer.nodeKey
		if slices.Contains(seen, nextPoint) {
			continue
		}
		foundDistance, endReached := walk(g, seen, nextPoint)
		if endReached {
			var distance int = foundDistance + nodePointer.distanceToNode
			if distance > longestPath {
				longestPath = distance
			}
		}
	}
	indexToDelete := slices.Index(seen, curr)
	seen = slices.Delete(seen, indexToDelete, indexToDelete+1)
	return longestPath, true
}
