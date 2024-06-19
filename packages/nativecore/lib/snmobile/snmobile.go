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

type HostReceiver interface {
	// When JS creates a node (or even the Mobile side)
	// this callback is executed
	OnNodeCreated(nodeId int, nodeType string)
	// Some nodes, like text & text input, need to be
	// measured while calculating layout before
	// sending it over the wire
	DoesNodeRequireMeasuring(nodeType string) bool
	// TODO: Determine how to handle this.
	OnLayoutChange(nodeId int, layoutMetrics *LayoutMetrics)
	OnPropUpdated(nodeId int, key string, value *JSValue)
	// TODO: Determine how to send the data over.
	// Can work with bytes, but need to determine the size of the int
	// to effectivly decode it.
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
	nodeParent    map[int]int
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
		nodeParent:   make(map[int]int),
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

// TODO: Give iterator type to retreive all
// native modules with reciever function.
// May need to use flex for type conversion here.
func (s *SolidNativeMobile) RegistureModules() {}

// ===================  JS Stuff =======================

// Create root node and return its ID.
// Not to be called on JS
// This removes the callback like effect and allows the host to create its root node immediatly
// to present it to the screen.
func (s *SolidNativeMobile) CreateRootNode() int {
	return s.createNodeAndDoNotNotifyHost()
}

// Creates node and notifies mobile host reciever
// to be typically called from JS side.
// Returns Node ID (which is an int)
func (s *SolidNativeMobile) CreateNode(nodeType string) int {
	nodeId := s.createNodeAndDoNotNotifyHost()
	s.hostReceiver.OnNodeCreated(nodeId, nodeType)
	return nodeId
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

	// Host Receiver will take in new JSValue for usage
	// However, the view doesn't update until we use the
	// `OnUpdateRevisionCount` method
	s.hostReceiver.OnPropUpdated(nodeId, key, value)

	// Update flex style and notify of new layout metrics
	if key == "style" {
		styleMap, err := s.convertJSToKeysAndObjects(value)
		if err != nil {
			return err
		}

		newStyleKeys := updateNodeStyleAndReturnNewStyleKeys(node, styleMap, prevKeys)

		s.nodeStyleKeys[nodeId] = newStyleKeys

		// Call the layout function, which will update the layout metrics and send it over
		// to the host. It will also notify dirty yoga nodes and update all the
		// revision counts needed.
		s.updateLayoutAndNotify(map[int]struct{}{
			nodeId: {},
		})
		return nil
	}

	s.hostReceiver.OnUpdateRevisionCount(nodeId)
	return nil
}

// Anchor is optional.
// TODO: Impliment Me!
func (s *SolidNativeMobile) InsertBefore(parentId int, newNodeId int, anchorId *int) {
	// If there's an anchor, insert before the anchor

	// Init to nil
	var newChildrenIds []int
	currentChildrenIds := s.nodeChildren[parentId]
	parentYogaNode := s.yogaNodes[parentId]
	newYogaNode := s.yogaNodes[newNodeId]

	if anchorId != nil {
		for ind, nodeId := range currentChildrenIds {
			if nodeId == *anchorId {
				parentYogaNode.InsertChild(newYogaNode, ind)
				newChildrenIds = append(newChildrenIds, newNodeId, nodeId)
			} else {
				newChildrenIds = append(newChildrenIds, nodeId)
			}
		}
	} else {
		// Add to the end
		newChildrenIds = append(newChildrenIds, currentChildrenIds...)
		newChildrenIds = append(newChildrenIds, newNodeId)
		ind := len(currentChildrenIds)
		parentYogaNode.InsertChild(newYogaNode, ind)
	}

	// Update internal children
	s.nodeChildren[parentId] = newChildrenIds
	s.nodeParent[newNodeId] = parentId
	// TODO: Send node over
	s.hostReceiver.OnChildrenChange(parentId)

	s.updateLayoutAndNotify(map[int]struct{}{})
}

func (s *SolidNativeMobile) RemoveChild(parentId int, childNodeId int) {
	parentChildIds := s.nodeChildren[parentId]
	newChildIds := make([]int, 0)

	for _, nodeId := range parentChildIds {
		if nodeId == childNodeId {
			continue
		}
		newChildIds = append(newChildIds, nodeId)
	}

	s.nodeChildren[parentId] = newChildIds

	// Cleanup on Yoga
	parentYogaNode := s.yogaNodes[parentId]
	childYogaNode := s.yogaNodes[childNodeId]

	parentYogaNode.RemoveChild(childYogaNode)
	delete(s.yogaNodes, childNodeId)
	delete(s.nodeChildren, childNodeId)
	delete(s.nodeStyleKeys, childNodeId)
	delete(s.nodeParent, childNodeId)
	childYogaNode.Free()

	s.updateLayoutAndNotify(map[int]struct{}{})
}

func (s *SolidNativeMobile) GetParent(nodeId int) (parentId int, exists bool) {
	parentId, exists = s.nodeParent[nodeId]
	return parentId, exists
}

func (s *SolidNativeMobile) GetFirstChild(nodeId int) (firstChildId int, exists bool) {
	nodeChildren := s.nodeChildren[nodeId]

	length := len(nodeChildren)

	if length == 0 {
		exists = false
		return firstChildId, exists
	}

	firstChildId = nodeChildren[0]
	exists = true

	return firstChildId, exists
}

func (s *SolidNativeMobile) GetNextSibling(nodeId int) (nextSiblingIndex int, exists bool) {
	parentId, exists := s.GetParent(nodeId)

	if !exists {
		return nextSiblingIndex, false
	}

	parentChildrenIds := s.nodeChildren[parentId]
	parentChildrenIdLength := len(parentChildrenIds)
	childIndex := 0

	for i, n := range parentChildrenIds {
		if n == nodeId {
			childIndex = i
			break
		}
	}

	nextSiblingIndex = childIndex + 1

	if nextSiblingIndex >= parentChildrenIdLength {
		return nextSiblingIndex, false
	}

	return nextSiblingIndex, true
}

// ================= Private Helper Methods =========================

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
func (s *SolidNativeMobile) updateLayoutAndNotify(modifiedNodes map[int]struct{}) error {
	if s.rootNodeId == nil {
		return fmt.Errorf("root node does not exist! cannot update layout")
	}
	rootNodeId := *s.rootNodeId
	yogaRootNode := s.yogaNodes[rootNodeId]

	yogaRootNode.CalculateLayout(yoga.YGUndefined, yoga.YGUndefined, yoga.DirectionLTR)

	s.applyLayout(rootNodeId)

	return nil
}

func (s *SolidNativeMobile) applyLayout(nodeId int) {
	node := s.yogaNodes[nodeId]

	if !node.GetHasNewLayout() {
		return
	}

	node.SetHasNewLayout(false)

	// TODO: Notify of new layout and update layout metrics
	s.hostReceiver.OnLayoutChange(nodeId, convertYogaLayoutMetricToSNLayoutMetrics(
		yoga.NewLayoutMetrics(node),
	))
	s.hostReceiver.OnUpdateRevisionCount(nodeId)

	for _, n := range s.nodeChildren[nodeId] {
		s.applyLayout(n)
	}
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

// Internal usage. Internally, we do not need to keep track of the node type
// TODO: But i do need some mechanism for the measure function
func (s *SolidNativeMobile) createNodeAndDoNotNotifyHost() int {
	id := int(uuid.New().ID())
	yogaNode := yoga.NewNode()
	s.yogaNodes[id] = yogaNode
	s.nodeChildren[id] = make([]int, 0)
	s.nodeStyleKeys[id] = make(Set)
	return id
}

func (s *SolidNativeMobile) downloadAndRunJs() {
	s.dukContext.EvalString("globalThis._SolidNativeRenderer.createNodeByName('sn_view')")
}
