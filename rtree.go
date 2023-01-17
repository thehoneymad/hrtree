package hrtree

import (
	"math"
)

const defaultFillFactor = 0.4
const defaultMaxEntries = 4
const defaultMinEntries = 2

type RTree struct {
	root              *Node
	max               int64
	min               int64
	count             int64
	defaultFillFactor float64
}

func NewRTree(_max int64) *RTree {
	_max = int64(math.Max(defaultMaxEntries, float64(_max)))
	_min := int64(math.Max(defaultMinEntries, math.Ceil(float64(_max)+defaultFillFactor)))

	rt := &RTree{
		max:               _max,
		min:               _min,
		defaultFillFactor: defaultFillFactor,
	}

	rt.Clear()

	return rt
}

func (rt *RTree) Insert(geom BBoxable) error {
	// Step 1: Create a bounding rectangle from the input
	// Step 2: If this tree is empty, as in no root node, create a root node.
	// Step 3: Start walking down the tree from root node. Start from the root node.
	// Step 4: In the depth first search, for each node, search for leaf nodes. Data only goes into leaf nodes in B+Tree fashion.
	// If it does not have a leaf node, expand the current bounding rectangle of this node to fit the geometry.
	// Find whether the leaf node needs to be updated.

	rt.insert(geom, rt.root.height)
	rt.count++

	return nil
}

func (rt *RTree) Clear() {
	rt.root = NewNode(make([]BBoxable, 0), 0)
	rt.count = 0
}

func (rt *RTree) insert(geom BBoxable, depth int32) {
	// Step 1: Create a bounding rectangle from the input
	// Step 2: If this tree is empty, as in no root node, create a root node.
	// Step 3: Start walking down the tree from root node. Start from the root node.
}
