package yoga

/*
#cgo CFLAGS: -I${SRCDIR}/../../yoga
#cgo LDFLAGS: ${SRCDIR}/../../yoga/build/yoga/libyogacore.a -lc++
#include <yoga/Yoga.h>
*/
import "C"

// Node represents a Yoga node
type Node struct {
	node C.YGNodeRef
}

// NewNode creates a new Yoga node
func NewNode() *Node {
	return &Node{node: C.YGNodeNew()}
}

// Free releases the resources allocated for the Yoga node
func (n *Node) Free() {
	C.YGNodeFree(n.node)
}

// SetWidth sets the width of the Yoga node
func (n *Node) SetWidth(width float32) {
	C.YGNodeStyleSetWidth(n.node, C.float(width))
}

// SetHeight sets the height of the Yoga node
func (n *Node) SetHeight(height float32) {
	C.YGNodeStyleSetHeight(n.node, C.float(height))
}

// SetMargin sets the margin for the specified edge of the Yoga node
func (n *Node) SetMargin(edge Edge, margin float32) {
	C.YGNodeStyleSetMargin(n.node, C.YGEdge(edge), C.float(margin))
}

// SetPadding sets the padding for the specified edge of the Yoga node
func (n *Node) SetPadding(edge Edge, padding float32) {
	C.YGNodeStyleSetPadding(n.node, C.YGEdge(edge), C.float(padding))
}

// CalculateLayout performs the layout calculation for the Yoga node
func (n *Node) CalculateLayout(width, height float32, direction Direction) {
	C.YGNodeCalculateLayout(n.node, C.float(width), C.float(height), C.YGDirection(direction))
}

// GetLayoutWidth gets the calculated layout width of the Yoga node
func (n *Node) GetLayoutWidth() float32 {
	return float32(C.YGNodeLayoutGetWidth(n.node))
}

// GetLayoutHeight gets the calculated layout height of the Yoga node
func (n *Node) GetLayoutHeight() float32 {
	return float32(C.YGNodeLayoutGetHeight(n.node))
}

// Additional functions and types can be added here as needed...

// AddChild adds a child node to the Yoga node
func (n *Node) AddChild(child *Node) {
	C.YGNodeInsertChild(n.node, child.node, C.YGNodeGetChildCount(n.node))
}

// RemoveChild removes a child node from the Yoga node
func (n *Node) RemoveChild(child *Node) {
	C.YGNodeRemoveChild(n.node, child.node)
}

// GetChildCount returns the number of children of the Yoga node
func (n *Node) GetChildCount() int {
	return int(C.YGNodeGetChildCount(n.node))
}

// GetChild returns the child node at the specified index
func (n *Node) GetChild(index int) *Node {
	return &Node{node: C.YGNodeGetChild(n.node, C.ulong(index))}
}
