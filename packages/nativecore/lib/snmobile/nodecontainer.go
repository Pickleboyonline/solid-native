package snmobile

import (
	"nativecore/lib/yoga"

	"github.com/google/uuid"
)

// TODO: Need to add some value to determine whether its a text or not
// TODO: Add some optimizations for host platform to be unaware of shadow text nodes.
// TODO: In fact, text nodes may need their own nodes later to handle their shadow children
type NodeContainer struct {
	id       string
	children []*NodeContainer
	// Not attached for children text nodes of other children.
	yogaNode *yoga.YGNode
	parent   *NodeContainer
	// Used to keep track of previous keys so we can set them
	// to default values if not set anymore
	// ? Potentially could encapsulate this in a "props" field, but right
	// ? now, we only need to keep track of styles for yoga nodes and text
	// ? style normalization
	// ? There are some text specific props: https://reactnative.dev/docs/text
	// ? But not sure which ones can actually be normalized like that.
	//
	// ! NOTE: text and styleMaps are mutually exclusive.
	styleMap map[string]JSValue

	isText bool
	// Text Value
	// Technically a prop but we have it here for easy access for `TextDescriptor` construction
	text string
	// TODO: Maybe add a `prev` and `next`reference so prev and next
	// TODO: sibling lookups are O(1)
}

// Creates NodeContainer with defaults for the pointer map and array types
// along with a unique ID
func newNodeContainer(isText bool) NodeContainer {
	return NodeContainer{
		id:       uuid.NewString(),
		yogaNode: yoga.NewNode(),
		children: make([]*NodeContainer, 0),
		styleMap: map[string]JSValue{},
		isText:   isText,
	}
}
