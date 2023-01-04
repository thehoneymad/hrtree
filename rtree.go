package hrtree

import "errors"

type boundingRect struct {
}

// TODO: This should be something that is geojson BBOX esque. Anything that can have a bbox can be stored in an RTree
type BBoxable interface {
	Width() int64
	Height() int64
	X() float64
	Y() float64
}

type RTree struct {
	root boundingRect
	max  int64
}

func NewRTree(root boundingRect, max int64) *RTree {
	return &RTree{root: root, max: max}
}

func (rt *RTree) insert(geom BBoxable) error {
	if geom.X() < 0 {
		return errors.New("expected x coordinate to be greater or equal to 0")
	}

	if geom.Y() < 0 {
		return errors.New("expected y coordinate to be greater or equal to 0")
	}

	if geom.Height() < 0 {
		return errors.New("expected height to be greater or equal to 0")
	}

	if geom.Width() < 0 {
		return errors.New("expected width to be greater or equal to 0")
	}

	// Step 1: Create a bounding rectangle from the input
	// Step 2: If this tree is empty, as in no root node, create a root node.
	// Step 3: Start walking down the tree from root node. Start from the root node.
	// Step 4: In the depth first search, for each node, search for leaf nodes. Data only goes into leaf nodes in B+Tree fashion.
	// If it does not have a leaf node, expand the current bounding rectangle of this node to fit the geometry.
	// Find whether the leaf node needs to be updated.

	return nil
}
