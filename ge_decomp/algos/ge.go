package ge_decomp

import (
    "container/list"
    "fmt"
)

type Graph struct {
    n   int
    adj [][]int
}

func NewGraph(n int) *Graph {
    adj := make([][]int, n)
    for i := 0; i < n; i++ {
        adj[i] = []int{}
    }
    return &Graph{n: n, adj: adj}
}

func (g *Graph) AddEdge(u, v int) {
    g.adj[u] = append(g.adj[u], v)
    g.adj[v] = append(g.adj[v], u)
}

type Blossom struct {
    n      int
    graph  [][]int
    match  []int
    parent []int
    base   []int
    used   []bool
    blossom []bool
}

func NewBlossom(n int, graph [][]int) *Blossom {
    match := make([]int, n)
    parent := make([]int, n)
    base := make([]int, n)
    used := make([]bool, n)
    blossom := make([]bool, n)
    for i := 0; i < n; i++ {
        match[i] = -1
        parent[i] = -1
        base[i] = i
        used[i] = false
        blossom[i] = false
    }
    return &Blossom{
        n:       n,
        graph:   graph,
        match:   match,
        parent:  parent,
        base:    base,
        used:    used,
        blossom: blossom,
    }
}

func (b *Blossom) lca(a, b1 int) int {
    used := make([]bool, b.n)
    for {
        a = b.base[a]
        used[a] = true
        if b.match[a] == -1 {
            break
        }
        a = b.parent[b.match[a]]
    }
    for {
        b1 = b.base[b1]
        if used[b1] {
            return b1
        }
        b1 = b.parent[b.match[b1]]
    }
}

func (b *Blossom) markPath(v, bbase, x int) {
    for b.base[v] != bbase {
        b.blossom[b.base[v]] = true
        b.blossom[b.base[b.match[v]]] = true
        b.parent[v] = x
        x = b.match[v]
        v = b.parent[b.match[v]]
    }
}

func (b *Blossom) findPath(root int) int {
    for i := 0; i < b.n; i++ {
        b.used[i] = false
        b.parent[i] = -1
        b.base[i] = i
    }
    queue := list.New()
    queue.PushBack(root)
    b.used[root] = true

    for queue.Len() > 0 {
        v := queue.Remove(queue.Front()).(int)
        for _, u := range b.graph[v] {
            if b.base[v] == b.base[u] || b.match[v] == u {
                continue
            }
            if u == root || (b.match[u] != -1 && b.parent[b.match[u]] != -1) {
                curBase := b.lca(v, u)
                for i := 0; i < b.n; i++ {
                    b.blossom[i] = false
                }
                b.markPath(v, curBase, u)
                b.markPath(u, curBase, v)
                for i := 0; i < b.n; i++ {
                    if b.blossom[b.base[i]] {
                        b.base[i] = curBase
                        if !b.used[i] {
                            b.used[i] = true
                            queue.PushBack(i)
                        }
                    }
                }
            } else if b.parent[u] == -1 {
                b.parent[u] = v
                if b.match[u] == -1 {
                    return u // найден дополняющий путь
                }
                if !b.used[b.match[u]] {
                    b.used[b.match[u]] = true
                    queue.PushBack(b.match[u])
                }
            }
        }
    }
    return -1
}

func (b *Blossom) augmentPath(start int) {
    v := start
    for v != -1 {
        pv := b.parent[v]
        w := -1
        if pv != -1 {
            w = b.match[pv]
        }
        b.match[v] = pv
        b.match[pv] = v
        v = w
    }
}

func (b *Blossom) Solve() []int {
    for i := 0; i < b.n; i++ {
        if b.match[i] == -1 {
            if endpoint := b.findPath(i); endpoint != -1 {
                b.augmentPath(endpoint)
            }
        }
    }
    return b.match
}

func edmondsBlossomMatching(g *Graph) []int {
    solver := NewBlossom(g.n, g.adj)
    return solver.Solve()
}

func maximumMatching(g *Graph) []int {
    return edmondsBlossomMatching(g)
}

func edmondsMaximumMatchingSize(g *Graph) int {
    m := maximumMatching(g)
    count := 0
    for _, v := range m {
        if v != -1 {
            count++
        }
    }
    return count / 2
}

func gallaiEdmondsDecomposition(g *Graph) (map[int]bool, map[int]bool, map[int]bool) {
    n := g.n
    match := maximumMatching(g)
    originalSize := 0
    for _, v := range match {
        if v != -1 {
            originalSize++
        }
    }
    originalSize /= 2

    D := make(map[int]bool)
    for v := 0; v < n; v++ {
        if match[v] == -1 {
            D[v] = true
        }
    }

    for v := 0; v < n; v++ {
        if D[v] {
            continue
        }
        tempGraph := NewGraph(n)
        for u := 0; u < n; u++ {
            if u == v {
                continue
            }
            for _, w := range g.adj[u] {
                if w == v || u >= w {
                    continue
                }
                tempGraph.AddEdge(u, w)
            }
        }
        tempMatch := maximumMatching(tempGraph)
        tempSize := 0
        for _, x := range tempMatch {
            if x != -1 {
                tempSize++
            }
        }
        tempSize /= 2
        if tempSize == originalSize {
            D[v] = true
        }
    }
    A := make(map[int]bool)
    for d := range D {
        for _, neighbor := range g.adj[d] {
            if !D[neighbor] {
                A[neighbor] = true
            }
        }
    }
    all := make(map[int]bool)
    for i := 0; i < n; i++ {
        all[i] = true
    }
    C := make(map[int]bool)
    for i := 0; i < n; i++ {
        if !D[i] && !A[i] {
            C[i] = true
        }
    }
    return D, A, C
}

func printMatching(m []int) {
    printed := make(map[[2]int]bool)
    for v, u := range m {
        if u != -1 {
            key1 := [2]int{v, u}
            key2 := [2]int{u, v}
            if !printed[key1] && !printed[key2] {
                fmt.Printf("%d - %d\n", v, u)
                printed[key1] = true
            }
        }
    }
}
