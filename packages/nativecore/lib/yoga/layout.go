package yoga

/*
#include <yoga/Yoga.h>
*/
import "C"

// GetLayoutLeft gets the left layout value of the Yoga node.
func (n *YGNode) GetLayoutLeft() float32 {
	return float32(C.YGNodeLayoutGetLeft(n.node))
}

// GetLayoutTop gets the top layout value of the Yoga node.
func (n *YGNode) GetLayoutTop() float32 {
	return float32(C.YGNodeLayoutGetTop(n.node))
}

// GetLayoutRight gets the right layout value of the Yoga node.
func (n *YGNode) GetLayoutRight() float32 {
	return float32(C.YGNodeLayoutGetRight(n.node))
}

// GetLayoutBottom gets the bottom layout value of the Yoga node.
func (n *YGNode) GetLayoutBottom() float32 {
	return float32(C.YGNodeLayoutGetBottom(n.node))
}

// GetLayoutWidth gets the width layout value of the Yoga node.
func (n *YGNode) GetLayoutWidth() float32 {
	return float32(C.YGNodeLayoutGetWidth(n.node))
}

// GetLayoutHeight gets the height layout value of the Yoga node.
func (n *YGNode) GetLayoutHeight() float32 {
	return float32(C.YGNodeLayoutGetHeight(n.node))
}

// GetLayoutDirection gets the direction layout value of the Yoga node.
func (n *YGNode) GetLayoutDirection() Direction {
	return Direction(C.YGNodeLayoutGetDirection(n.node))
}

// GetLayoutMargin gets the margin layout value of the Yoga node for the given edge.
func (n *YGNode) GetLayoutMargin(edge Edge) float32 {
	return float32(C.YGNodeLayoutGetMargin(n.node, C.YGEdge(edge)))
}

// GetLayoutBorder gets the border layout value of the Yoga node for the given edge.
func (n *YGNode) GetLayoutBorder(edge Edge) float32 {
	return float32(C.YGNodeLayoutGetBorder(n.node, C.YGEdge(edge)))
}

// GetLayoutPadding gets the padding layout value of the Yoga node for the given edge.
func (n *YGNode) GetLayoutPadding(edge Edge) float32 {
	return float32(C.YGNodeLayoutGetPadding(n.node, C.YGEdge(edge)))
}

// TODO: Be able to set undefined if needed.
func (n *YGNode) CalculateLayout(availableWidth, availableHeight float32, direction Direction) {
	// return float32(C.YGNodeLayoutGetPadding(n.node, C.YGEdge(edge)))
	C.YGNodeCalculateLayout(n.node, C.float(availableWidth), C.float(availableHeight), C.YGDirection(direction))
}

func (n *YGNode) GetHasNewLayout() bool {
	return bool(C.YGNodeGetHasNewLayout(n.node))
}

func (n *YGNode) SetHasNewLayout(hasNewLayout bool) {
	C.YGNodeSetHasNewLayout(n.node, C.bool(hasNewLayout))
}
