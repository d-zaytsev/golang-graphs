package mst

const NO_PARENT_ID = -1

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) DSU {
	dsu := DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := range dsu.parent {
		dsu.parent[i] = NO_PARENT_ID
		dsu.size[i] = 1
	}
	return dsu
}

func (dsu *DSU) Find(v int) int {
	if dsu.parent[v] == NO_PARENT_ID {
		return v
	}
	dsu.parent[v] = dsu.Find(dsu.parent[v])
	return dsu.parent[v]
}

func (dsu *DSU) Union(v1, v2 int) {
	v1_parent := dsu.Find(v1)
	v2_parent := dsu.Find(v2)
	if v1_parent == v2_parent {
		return
	}
	if dsu.size[v1_parent] < dsu.size[v2_parent] {
		v1_parent, v2_parent = v2_parent, v1_parent
		v2 = v1
	}
	for {
		if v2 == v2_parent {
			dsu.parent[v2_parent] = v1_parent
			break
		}
		dsu.parent[v2], v2 = v1_parent, dsu.parent[v2]
	}
	dsu.size[v1_parent] += dsu.size[v2_parent]
}

func (d *DSU) Size(v int) int {
	return d.size[d.Find(v)]
}
