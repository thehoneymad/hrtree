package hrtree

import "math"

// TODO: This should be something that is geojson BBOX esque. Anything that can have a bbox can be stored in an RTree
type BBoxable interface {
	Bound() BBox
}

type BBox struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func (b *BBox) Extend(other BBox) *BBox {
	return &BBox{
		MinX: math.Min(b.MinX, other.MinX),
		MinY: math.Min(b.MinY, other.MinY),
		MaxX: math.Max(b.MaxY, other.MaxX),
		MaxY: math.Max(b.MaxY, other.MaxY),
	}
}

func emptyBounds() *BBox {
	return &BBox{
		MinX: math.Inf(1),
		MinY: math.Inf(1),
		MaxX: math.Inf(-1),
		MaxY: math.Inf(-1),
	}
}

func computeBBox(items []BBoxable) *BBox {
	bb := emptyBounds()
	for _, item := range items {
		bb = bb.Extend(item.Bound())
	}

	return bb
}
