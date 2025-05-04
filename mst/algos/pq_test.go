package mst

import (
	"container/heap"
	"math/rand"
	"testing"
	"time"
)

func TestPriorityQueue_Push_Pop(t *testing.T) {
	pq := NewPQ[string, int]()

	pq.Push(Node[string, int]{Value: "A", Priority: 3})
	pq.Push(Node[string, int]{Value: "B", Priority: 1})
	pq.Push(Node[string, int]{Value: "C", Priority: 2})

	if pq.Len() != 3 {
		t.Errorf("Expected length 3, got %d", pq.Len())
	}

	item1, _ := pq.Pop().(Node[string, int])
	if item1.Value != "B" {
		t.Errorf("Expected B, got %s", item1.Value)
	}
	item2, _ := pq.Pop().(Node[string, int])
	if item2.Value != "C" {
		t.Errorf("Expected C, got %s", item2.Value)
	}
	item3, _ := pq.Pop().(Node[string, int])
	if item3.Value != "A" {
		t.Errorf("Expected A, got %s", item3.Value)
	}

	if !pq.IsEmpty() {
		t.Errorf("Expected empty, got length %d", pq.Len())
	}
}

func TestPriorityQueue_Update(t *testing.T) {
	pq := NewPQ[string, int]()

	heap.Push(pq, Node[string, int]{Value: "A", Priority: 3})
	heap.Push(pq, Node[string, int]{Value: "B", Priority: 1})
	heap.Push(pq, Node[string, int]{Value: "C", Priority: 2})

	pq.Update("A", 0)

	item1, _ := pq.Pop().(Node[string, int])
	if item1.Value != "A" {
		t.Errorf("Expected A, got %s", item1.Value)
	}
	item2, _ := pq.Pop().(Node[string, int])
	if item2.Value != "B" {
		t.Errorf("Expected B, got %s", item2.Value)
	}
	item3, _ := pq.Pop().(Node[string, int])
	if item3.Value != "C" {
		t.Errorf("Expected C, got %s", item3.Value)
	}
}

func TestPriorityQueue_Contains(t *testing.T) {
	pq := NewPQ[string, int]()

	heap.Push(pq, Node[string, int]{Value: "A", Priority: 3})
	heap.Push(pq, Node[string, int]{Value: "B", Priority: 1})

	if !pq.Contains("A") {
		t.Errorf("Expected to contain A")
	}
	if !pq.Contains("B") {
		t.Errorf("Expected to contain B")
	}
	if pq.Contains("C") {
		t.Errorf("Expected not to contain C")
	}
}

func TestPriorityQueue_DuplicatePush(t *testing.T) {
	pq := NewPQ[string, int]()

	heap.Push(pq, Node[string, int]{Value: "A", Priority: 3})

	heap.Push(pq, Node[string, int]{Value: "A", Priority: 1}) // Duplicate, should be ignored

	if pq.Len() != 1 {
		t.Errorf("Expected length 1 after duplicate push, got %d", pq.Len())
	}

	item1, _ := pq.Pop().(Node[string, int])

	if item1.Priority != 3 {
		t.Errorf("Expected priority to be 3")
	}
}

func TestPriorityQueue_RandomPriorities(t *testing.T) {
	pq := NewPQ[int, int]()
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	numNodes := 100

	for i := 0; i < numNodes; i++ {
		priority := rand.Intn(1000)
		heap.Push(pq, Node[int, int]{Value: i, Priority: priority})
	}

	// Check if the priorities are in order after popping
	prevPriority := -1
	for i := 0; i < numNodes; i++ {
		item := pq.Pop().(Node[int, int])
		if item.Priority < prevPriority {
			t.Errorf("Priority order is incorrect: %d < %d", item.Priority, prevPriority)
		}
		prevPriority = item.Priority
	}
}

func TestPriorityQueue_StringValue(t *testing.T) {
	pq := NewPQ[string, int]()

	// Push some nodes
	heap.Push(pq, Node[string, int]{Value: "banana", Priority: 3})
	heap.Push(pq, Node[string, int]{Value: "apple", Priority: 1})
	heap.Push(pq, Node[string, int]{Value: "orange", Priority: 2})

	// Check the length
	if pq.Len() != 3 {
		t.Errorf("Expected length 3, got %d", pq.Len())
	}

	// Pop nodes and check priority order
	item1, _ := pq.Pop().(Node[string, int])
	if item1.Value != "apple" {
		t.Errorf("Expected apple, got %s", item1.Value)
	}
	item2, _ := pq.Pop().(Node[string, int])
	if item2.Value != "orange" {
		t.Errorf("Expected orange, got %s", item2.Value)
	}
	item3, _ := pq.Pop().(Node[string, int])
	if item3.Value != "banana" {
		t.Errorf("Expected banana, got %s", item3.Value)
	}
}
