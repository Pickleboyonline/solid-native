/*
Package serves as a way to interface with iOS and Android
TODO: Mobile will create an object here.
*/
package snmobile

import (
	"nativecore/lib/collections"
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
	OnPropUpdated(nodeId int, key string, value *JsValue)
	OnChildrenChange(nodeId int)
	// Signifies when its time to update JetpackCompose/SwiftUI
	OnUpdateRevisionCount(nodeId int)
}

// Houses important info.
type SolidNativeMobile struct {
	// Get chidren
	nodeChildren  map[int][]int
	yogaNodes     map[int]*yoga.Node
	nodeStyleKeys map[int]*collections.Set
	hostReceiver  HostReceiver
	dukContext    *duktape.Context
	// Set to -1 initally, need to set before calulcating layouts
	rootNodeId *int
}

func NewSolidNativeMobile(hostReceiver HostReceiver) *SolidNativeMobile {
	ctx := duktape.New()
	// ctx.PushGoFunction()
	return &SolidNativeMobile{
		yogaNodes:    make(map[int]*yoga.Node),
		hostReceiver: hostReceiver,
		dukContext:   ctx,
		rootNodeId:   nil,
		// We use this to keep track of when keys are removed
		// Since Yoga works via mutation
		nodeStyleKeys: make(map[int]*collections.Set),
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
	s.nodeStyleKeys[id] = collections.NewSet()
	s.hostReceiver.OnNodeCreated(id, nodeType)
	return id
}

// Updates the host receiver about the props from the JS side
// TODO: Determine how to send it over. & update to take value
// TODO: Prob will need some JSValue like object to hold any object.
// Value can be a JSValue
// or primative.
// JS Value can be array
func (s *SolidNativeMobile) SetNodeProp(nodeId int, key string, value *JsValue) {
	// TODO: Send new value
	s.hostReceiver.OnPropUpdated(nodeId, key, value)

	// Update flex style and notify
	if key == "style" {
		// TODO: Update yoga layout in node
		s.updateFlexStyle()
		s.updateLayoutAndNotify(collections.NewSet())
		return
	}

	s.hostReceiver.OnUpdateRevisionCount(nodeId)
}

// Anchor is optional.
func (s *SolidNativeMobile) InsertBefore(parentId int, newNodeId int, anchorId *int) {
	// TODO: Update flex style

	// TODO: Update children of something
	s.hostReceiver.OnChildrenChange(parentId)

	s.updateLayoutAndNotify(collections.NewSet())
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
func (s *SolidNativeMobile) updateLayoutAndNotify(modifiedNodes *collections.Set) {

}

func (s *SolidNativeMobile) updateFlexStyle() {}
