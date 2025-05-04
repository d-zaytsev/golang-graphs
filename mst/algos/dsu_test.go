package mst

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewDSU(t *testing.T) {
	dsu := NewDSU(5)
	assert.Equal(t, 5, len(dsu.parent))
	assert.Equal(t, 5, len(dsu.size))
	for i := 0; i < 5; i++ {
		assert.Equal(t, NO_PARENT_ID, dsu.parent[i])
		assert.Equal(t, 1, dsu.size[i])
	}
}

func TestFind(t *testing.T) {
	dsu := NewDSU(5)
	assert.Equal(t, 0, dsu.Find(0))
	assert.Equal(t, 1, dsu.Find(1))
	assert.Equal(t, 2, dsu.Find(2))

	dsu.Union(0, 1)
	assert.Equal(t, dsu.Find(0), dsu.Find(1))

	dsu.Union(1, 2)
	assert.Equal(t, dsu.Find(0), dsu.Find(2))
}

func TestUnionAll(t *testing.T) {
	dsu := NewDSU(5)

	dsu.Union(0, 1)
	assert.Equal(t, dsu.Find(0), dsu.Find(1))
	assert.Equal(t, 2, dsu.size[dsu.Find(0)])
	dsu.Union(2, 3)
	assert.Equal(t, dsu.Find(2), dsu.Find(3))
	assert.Equal(t, 2, dsu.size[dsu.Find(3)])

	dsu.Union(0, 2)
	assert.Equal(t, dsu.Find(0), dsu.Find(2))
	assert.Equal(t, dsu.Find(0), dsu.Find(3))
	assert.Equal(t, dsu.Find(1), dsu.Find(3))

	root := dsu.Find(0)
	assert.Equal(t, 4, dsu.size[root])

	// Union 4 and 4 (no change)
	dsu.Union(4, 4)
	assert.Equal(t, NO_PARENT_ID, dsu.parent[4])
	assert.Equal(t, 1, dsu.size[4])

	// Union 4 and 0
	dsu.Union(4, 0)
	assert.Equal(t, dsu.Find(0), dsu.Find(4))
	assert.Equal(t, 5, dsu.size[root])
}

func TestSize(t *testing.T) {
	dsu := NewDSU(5)
	assert.Equal(t, 1, dsu.Size(0))
	assert.Equal(t, 1, dsu.Size(1))

	dsu.Union(0, 1)
	assert.Equal(t, 2, dsu.Size(0))
	assert.Equal(t, 2, dsu.Size(1))

	dsu.Union(0, 2)
	assert.Equal(t, 3, dsu.Size(0))
	assert.Equal(t, 3, dsu.Size(1))
	assert.Equal(t, 3, dsu.Size(2))

	dsu.Union(3, 4)
	assert.Equal(t, 2, dsu.Size(3))

	dsu.Union(0, 3)
	assert.Equal(t, 5, dsu.Size(0))
	assert.Equal(t, 5, dsu.Size(3))
	assert.Equal(t, 5, dsu.Size(4))
}
