package main

import "container/heap"

type cityBlock struct {
	heatLoss       int
	itemPoint      point
	dir            Direction
	noOfTimesInDir int
	index          int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*cityBlock

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest heat loss
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*cityBlock)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *cityBlock, heatLoss int, p point, dir Direction, noOfTimesInDir int, index int) {
	item.heatLoss = heatLoss
	item.itemPoint = p
	item.dir = dir
	item.noOfTimesInDir = noOfTimesInDir
	heap.Fix(pq, item.index)
}
