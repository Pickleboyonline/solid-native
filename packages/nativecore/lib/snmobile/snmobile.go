/*
Package serves as a way to interface with iOS and Android
TODO: Mobile will create an object here.
*/
package snmobile

import (
	"fmt"
	"nativecore/lib/yoga"

	"github.com/google/uuid"
	"gopkg.in/olebedev/go-duktape.v3"
)

type LayoutMetric struct {
	X int
	Y int
}

type HostReceiver interface {
	// When JS creates a node (or even the Mobile side)
	// this callback is executed
	OnNodeCreated(nodeId int, nodeType string)
	// Some nodes, like text & text input, need to be
	// measured while calculating layout before
	// sending it over the wire
	DoesNodeRequireMeasuring(nodeType string) bool
	// TODO: Determine how to handle this.
	OnLayoutChange(nodeId int, layoutMetric LayoutMetric)
	OnPropUpdated(nodeId int, key string, value *JSValue)
	OnChildrenChange(nodeId int)
	// Signifies when its time to update JetpackCompose/SwiftUI
	OnUpdateRevisionCount(nodeId int)
}

// Houses important info.
type SolidNativeMobile struct {
	// Get chidren
	nodeChildren  map[int][]int
	yogaNodes     map[int]*yoga.YGNode
	nodeStyleKeys map[int]Set
	hostReceiver  HostReceiver
	dukContext    *duktape.Context
	// Set to -1 initally, need to set before calulcating layouts
	rootNodeId *int
}

func NewSolidNativeMobile(hostReceiver HostReceiver) *SolidNativeMobile {
	ctx := duktape.New()
	// ctx.PushGoFunction()
	return &SolidNativeMobile{
		yogaNodes:    make(map[int]*yoga.YGNode),
		hostReceiver: hostReceiver,
		dukContext:   ctx,
		rootNodeId:   nil,
		// We use this to keep track of when keys are removed
		// Since Yoga works via mutation
		nodeStyleKeys: make(map[int]Set),
	}
}

// Registure core into system, download, and runs js.
func (s *SolidNativeMobile) RunJs() {
	s.registureRenderer()
	s.downloadAndRunJs()
}

// Yoga and JSContext need to be cleaned up before this object is deallocated
func (s *SolidNativeMobile) FreeMemory() {
	s.dukContext.DestroyHeap()
}

func (s *SolidNativeMobile) downloadAndRunJs() {
	s.dukContext.EvalString("globalThis._SolidNativeRenderer.createNodeByName('sn_view')")
}

// TODO: Give iterator type to retreive all
// native modules with reciever function.
// May need to use flex for type conversion here.
func (s *SolidNativeMobile) RegistureModules() {}

// ===================  JS Stuff =======================

// Create root node and return its ID.
// Not to be called on JS
func (s *SolidNativeMobile) CreateRootNode(nodeType string) {

}

// Creates node and notifies mobile host reciever
// to be typically called from JS side.
func (s *SolidNativeMobile) CreateNode(nodeType string) int {
	id := int(uuid.New().ID())
	yogaNode := yoga.NewNode()
	s.yogaNodes[id] = yogaNode
	s.nodeChildren[id] = make([]int, 0)
	s.nodeStyleKeys[id] = make(Set)
	s.hostReceiver.OnNodeCreated(id, nodeType)
	return id
}

// Updates the host receiver about the props from the JS side
// TODO: Determine how to send it over. & update to take value
// TODO: Prob will need some JSValue like object to hold any object.
// Value can be a JSValue
// or primative.
// JS Value can be array
func (s *SolidNativeMobile) SetNodeProp(nodeId int, key string, value *JSValue) error {
	node, exists := s.yogaNodes[nodeId]

	if !exists {
		return fmt.Errorf("node does not exist with id %v", nodeId)
	}

	prevKeys, exists := s.nodeStyleKeys[nodeId]

	// Silent error, fix it as needed. Should not happen however.
	if !exists {
		prevKeys = make(Set)
		s.nodeStyleKeys[nodeId] = prevKeys
	}

	// TODO: Send new value
	s.hostReceiver.OnPropUpdated(nodeId, key, value)

	// Update flex style and notify
	if key == "style" {
		// TODO: Update yoga layout in node
		// Convert JS value to styles
		styleMap, err := s.convertJSToKeysAndObjects(value)

		if err != nil {
			return err
		}
		updateNodeStyleAndReturnNewStyleKeys(node, styleMap, prevKeys)
		s.updateLayoutAndNotify(map[int]struct{}{
			nodeId: {},
		})
		return nil
	}

	s.hostReceiver.OnUpdateRevisionCount(nodeId)
	return nil
}

// "Upwraps" JS Value by enumerating over its keys
// and values. Ensure this is an object, otherwise just return nothing.
func (s *SolidNativeMobile) convertJSToKeysAndObjects(value *JSValue) (map[string]JSValue, error) {
	// Check whether its an object or array:
	s.dukContext.PushGlobalStash() // Push stash => [ stash ]

	s.dukContext.GetPropString(-1, value.stashKeyName) // => [ stash value ]

	valueType := s.dukContext.GetType(-1) // => [ stash value ]

	if !valueType.IsObject() {
		return nil, fmt.Errorf("js value with key %s is not an object or does not exist", value.stashKeyName)
	}

	jsValueMap := make(map[string]JSValue)

	// TODO: Check if 0 works or if i have to put in a flag
	s.dukContext.Enum(-1, 0) // => [ stash value enum ]

	for s.dukContext.Next(-1, true) {
		// => [ ... enum key value ]
		key := s.dukContext.GetString(-2)
		keyStashName := uuid.New().String()

		s.dukContext.PushGlobalStash() // => [ ... enum key value stash ]

		s.dukContext.Replace(-3) // => [ ... enum stash value ]

		s.dukContext.PutPropString(-2, keyStashName) // => [ ... enum stash ]

		s.dukContext.Pop() // => [ ... enum ]

		jsValueMap[key] = *NewJsValue(valueType, keyStashName, s)
	}

	return jsValueMap, nil
}

// Anchor is optional.
func (s *SolidNativeMobile) InsertBefore(parentId int, newNodeId int, anchorId *int) {
	// TODO: Update flex style

	// TODO: Update children of something
	s.hostReceiver.OnChildrenChange(parentId)

	s.updateLayoutAndNotify(map[int]struct{}{})
}

// Style is the only prop that relates to layout info
// Note: we want to batch layout + any other data change
// Data change happens first
// Then layout must be calc.
// If node is dirty and traversed, send layout + data
// If node was not data and data still needed to be send, send it over.
// ! Generally, this happens over a single thread in the call stack
// ! so SwiftUI/Compose shouldnt notice until call stack is empty, but just be
// ! aware.
// ! Example change:
// 1. Font Size change
// 2. Font size sent to
//

// Call after a prop is changed related to layout/style
//
// # modifiedNodes:
//
// Serves as a way to mark which nodes are dirty since sometimes
// the yoga layout does not change as a result. We still want to dispatch to the
// host that something changed (update revision count)
func (s *SolidNativeMobile) updateLayoutAndNotify(modifiedNodes map[int]struct{}) {

}
