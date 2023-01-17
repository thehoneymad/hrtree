package hrtree

type Node struct {
	bbox   BBox
	height int32
	items  []BBoxable
}

func NewNode(items []BBoxable, h int32) *Node {
	n := &Node{
		height: h,
		items:  items,
	}

	n.resetBBox()
	return n
}

func (n *Node) Children() []BBoxable {
	return n.items
}

func (n *Node) Bound() BBox {
	// TODO: Calculate bounding box of all children
	panic("implement me")
}

func (n *Node) IsLeaf() bool {
	return n.height == 1
}

func (n *Node) resetBBox() {
	n.bbox = *computeBBox(n.items)
}
