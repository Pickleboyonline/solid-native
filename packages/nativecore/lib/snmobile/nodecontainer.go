package snmobile

import (
	"nativecore/lib/yoga"

	"github.com/google/uuid"
)

// TODO: Need to add some value to determine whether its a
type NodeContainer struct {
	id       string
	children []*NodeContainer
	yogaNode *yoga.YGNode
	parent   *NodeContainer
	// Used to keep track of previous keys so we can set them
	// to default values if not set anymore
	yogaStyleKeys Set

	isText bool
	// TODO: Maybe add a `prev` and `next`reference so prev and next
	// TODO: sibling lookups are O(1)
}

// Creates NodeContainer with defaults for the pointer map and array types
// along with a unique ID
func newNodeContainer(isText bool) NodeContainer {
	return NodeContainer{
		id:            uuid.NewString(),
		yogaNode:      yoga.NewNode(),
		children:      make([]*NodeContainer, 0),
		yogaStyleKeys: Set{},
		isText:        isText,
	}
}
