package yoga

/*
#include <yoga/Yoga.h>
*/
import "C"

func (n *YGNode) InsertChild(childNode *YGNode, atIndex int) {
	C.YGNodeInsertChild(n.node, childNode.node, C.ulong(atIndex))
}

func (n *YGNode) RemoveChild(childNode *YGNode) {
	C.YGNodeRemoveChild(n.node, childNode.node)
}
