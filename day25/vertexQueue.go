package main

type vertexQueue struct {
	verticies []string
}

func InitVertexQueue() *vertexQueue {
	var q *vertexQueue = &vertexQueue{}
	q.verticies = []string{}
	return q
}

func (q *vertexQueue) enqueue(v string) {
	q.verticies = append(q.verticies, v)
}

func (q *vertexQueue) dequeue() string {
	var result string = q.verticies[0]
	if len(q.verticies) == 1 {
		q.verticies = q.verticies[:len(q.verticies)-1]
	} else {
		q.verticies = q.verticies[1:]
	}
	return result
}

func (q *vertexQueue) isEmpty() bool {
	return len(q.verticies) <= 0
}
