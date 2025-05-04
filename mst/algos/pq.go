package mst

import (
	"container/heap"

	"golang.org/x/exp/constraints"
)

type Node[V comparable, P constraints.Ordered] struct {
	Value    V
	Priority P
}

type PriorityQueue[V comparable, P constraints.Ordered] struct {
	heap []Node[V, P]
	set  map[V]int // for fast Update func
}

func NewPQ[V comparable, P constraints.Ordered]() *PriorityQueue[V, P] {
	pq := &PriorityQueue[V, P]{
		heap: make([]Node[V, P], 0),
		set:  make(map[V]int)}
	heap.Init(pq)
	return pq
}

// Pop implements heap.Interface.
func (pq *PriorityQueue[V, P]) Pop() any {
	if pq.IsEmpty() {
		panic("Queue is empty")
	}
	n := len(pq.heap)
	item := pq.heap[0]
	pq.Swap(0, n-1)
	pq.heap = pq.heap[:n-1]
	heap.Fix(pq, 0)
	delete(pq.set, item.Value)
	return item
}

func (pq *PriorityQueue[V, P]) Len() int {
	return len(pq.heap)
}

// Less implements heap.Interface.
func (pq *PriorityQueue[V, P]) Less(i int, j int) bool {
	return pq.heap[i].Priority < pq.heap[j].Priority
}

// Swap implements heap.Interface.
func (pq *PriorityQueue[V, P]) Swap(i int, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
	pq.set[pq.heap[i].Value], pq.set[pq.heap[j].Value] = i, j
}

func (pq *PriorityQueue[V, P]) Push(node any) {
	node_casted := node.(Node[V, P])
	if _, exists := pq.set[node_casted.Value]; exists {
		return
	}
	pq.set[node_casted.Value] = len(pq.heap)
	pq.heap = append(pq.heap, node_casted)
	heap.Fix(pq, len(pq.heap)-1)
}

func (pq *PriorityQueue[V, P]) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue[V, P]) Contains(value V) bool {
	_, exist := pq.set[value]
	return exist
}
func (pq *PriorityQueue[V, P]) Update(value V, priority P) {
	if index, exist := pq.set[value]; exist {
		pq.heap[index].Priority = priority
		heap.Fix(pq, index)
	}
}
