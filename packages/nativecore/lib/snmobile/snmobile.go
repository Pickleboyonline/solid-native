/*
Package serves as a way to interface with iOS and Android
TODO: Mobile will create an object here.
*/
package snmobile

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"nativecore/lib/yoga"
	"net/http"

	"gopkg.in/olebedev/go-duktape.v3"
)

// Houses important info.
type SolidNativeMobile struct {
	// Get chidren
	// nodeChildren  map[string][]string
	// yogaNodes     map[string]*yoga.YGNode
	// nodeStyleKeys map[string]Set
	// nodeParent    map[string]string
	nodeContainers map[string]*NodeContainer
	hostReceiver   HostReceiver
	dukContext     *duktape.Context
	// Set to -1 initally, need to set before calulcating layouts
	rootNodeId       string
	deviceScreenSize *Size
}

func NewSolidNativeMobile(hostReceiver HostReceiver) *SolidNativeMobile {
	ctx := duktape.New()

	// ctx.PushGoFunction()
	return &SolidNativeMobile{
		hostReceiver:     hostReceiver,
		dukContext:       ctx,
		rootNodeId:       "",
		deviceScreenSize: hostReceiver.GetDeviceScreenSize(),
		nodeContainers:   map[string]*NodeContainer{},
	}
}

// TODO: Give iterator type to retrieve all
// native modules with receiver function.
// May need to use flex for type conversion here.
func (s *SolidNativeMobile) RegisterModules() {
	s.dukContext.PushTimers()
	s.dukContext.PushGlobalGoFunction("log", func(c *duktape.Context) int {
		fmt.Println(c.SafeToString(-1))
		return 0
	})
	s.registerRenderer()
}

// Registure core into system, download, and runs js.
func (s *SolidNativeMobile) RunJsFromServer(url string) error {
	s.RegisterModules()
	return s.downloadAndRunJs(url)
}

// Registure core into system, download, and runs js.
func (s *SolidNativeMobile) EvalJs(jsToEval string) error {
	s.RegisterModules()
	return s.dukContext.PevalString(jsToEval)
}

// Yoga and JSContext need to be cleaned up before this object is deallocated
func (s *SolidNativeMobile) FreeMemory() {
	s.dukContext.DestroyHeap()
}

// TODO: determine how this is implimneted and if i can just update the device dimensions
func (s *SolidNativeMobile) OnOrientationChange() {

}

// Create root node and return its ID.
// Not to be called on JS
// This removes the callback like effect and allows the host to create its root node immediatly
// to present it to the screen.
//
// Use the nodetype to tell whether we need to measure it or not
func (s *SolidNativeMobile) CreateRootNode(nodeType string) string {
	nodeContainer := s.createNodeAndDoNotNotifyHost(nodeType)

	s.rootNodeId = nodeContainer.id

	yogaNode := nodeContainer.yogaNode

	// Ensure proper height/width
	yogaNode.SetWidth(s.deviceScreenSize.Width)
	yogaNode.SetHeight(s.deviceScreenSize.Height)

	return nodeContainer.id
}

// Internal function to setup the rendeder
//   - How to throw errors: https://duktape.org/api.html#concepts.9
func (s *SolidNativeMobile) registerRenderer() {
	// Push global object, we'll use it later
	s.dukContext.PushGlobalObject() // => [ globalThis ]

	// ======= Setup functions

	s.dukContext.PushObject() // => [globalThis obj]

	// Start our functions for solid js. Honesly rn not sure how i want tos tructure the API
	// for now, just slap them in a global object called _SolidNativeRenderer

	addGoFunc := func(cbName string, cb func(ctx *duktape.Context) int) {
		// [global obj]
		s.dukContext.PushGoFunction(cb) // => [global obj func]
		s.dukContext.PutPropString(-2, cbName)
		// => [global obj]
	}

	addGoFunc("createNodeByName", func(ctx *duktape.Context) int {
		// => [ nodeType ]
		nodeType := ctx.GetString(-1)

		nodeId := s.createNode(nodeType)

		log.Printf("New Node create of type %v with id %v", nodeType, nodeId)
		ctx.PushString(nodeId) // => [ nodeType nodeId ]

		return 1 // Return top of stack
	})

	addGoFunc("setProp", func(ctx *duktape.Context) int {
		// Stack: [ nodeId key value ]
		// Get second arg type
		nodeId := ctx.GetString(-3)
		key := ctx.GetString(-2)

		if ctx.IsFunction(-1) {
			log.Printf("Key \"%v\" for node %v was set with a function. Skipping as not supported...", key, nodeId)
			return 0
		}

		if ctx.IsUndefined(-1) {
			// We pass in an undefined thing here.
			s.setNodeProp(nodeId, key,
				&JSValue{})
		}

		encodedJson := ctx.JsonEncode(-1)

		var unmarshaledValue interface{}

		json.Unmarshal([]byte(encodedJson), &unmarshaledValue)

		log.Printf("Json of prop: %v", unmarshaledValue)

		s.setNodeProp(nodeId, key,
			&JSValue{
				data: unmarshaledValue,
			})

		return 0
	})

	addGoFunc("insertBefore", func(ctx *duktape.Context) int {
		// => [parentId nodeId anchorId?]
		parentId := ctx.GetString(-3)
		nodeId := ctx.GetString(-2)
		anchorType := ctx.GetType(-1)
		isAnchorAString := anchorType.IsString()

		var anchorId string

		if isAnchorAString {
			anchorId = ctx.GetString(-1)
		}

		log.Printf("Node Insert called with child %v  to be inserted under parent %v", nodeId, parentId)
		s.insertBefore(parentId, nodeId, anchorId)
		log.Printf("Node %v inserted under parent %v", nodeId, parentId)
		return 0
	})

	addGoFunc("isTextElement", func(ctx *duktape.Context) int {
		// => [nodeId]
		nodeId := ctx.GetString(-1)
		isTextElement := s.hostReceiver.IsTextElementByNodeId(nodeId)

		ctx.PushBoolean(isTextElement)

		return 1
	})

	addGoFunc("removeChild", func(ctx *duktape.Context) int {
		// => [parentId nodeId]
		parentId := ctx.GetString(-2)
		nodeId := ctx.GetString(-1)

		s.removeChild(parentId, nodeId)
		return 0
	})

	addGoFunc("getParent", func(ctx *duktape.Context) int {
		// => [nodeId]
		nodeId := ctx.GetString(-1)

		parentId, exists := s.getParent(nodeId)

		if !exists {
			return 0
		}

		ctx.PushString(parentId) // => [nodeId parentId]

		return 1
	})

	addGoFunc("getFirstChild", func(ctx *duktape.Context) int {
		// => [nodeId]
		nodeId := ctx.GetString(-1)

		firstChildId, exists := s.getFirstChild(nodeId)

		if !exists {
			return 0
		}

		ctx.PushString(firstChildId) // => [nodeId firstChildId]

		return 1
	})

	addGoFunc("getNextSibling", func(ctx *duktape.Context) int {
		// => [nodeId]
		nodeId := ctx.GetString(-1)

		nextSiblingId, exists := s.getNextSibling(nodeId)

		if !exists {
			return 0
		}

		ctx.PushString(nextSiblingId) // => [nodeId nextSiblingIndex]

		return 1
	})

	addGoFunc("getRootView", func(ctx *duktape.Context) int {

		ctx.PushString(s.rootNodeId) // => [nodeId rootNodeId]

		return 1
	})

	// Final Step, but the renderer there:
	s.dukContext.PutPropString(-2, "_SolidNativeRenderer") // => [global]
	s.dukContext.Pop()                                     // => []
}

// ===================  JS Stuff =======================
// Used on JS Renderer. Maybe make into its own struct?

// Creates node and notifies mobile host reciever
// to be typically called from JS side.
// Returns Node ID (which is an int)
func (s *SolidNativeMobile) createNode(nodeType string) string {
	nodeContainer := s.createNodeAndDoNotNotifyHost(nodeType)
	s.hostReceiver.OnNodeCreated(nodeContainer.id, nodeType)
	return nodeContainer.id
}

// Updates the host receiver about the props from the JS side
// Value can be a JSValue
// or primative.
// JS Value can be array
// The old JS value associatted does not need to be freed because it has
// a hashed ID. You only need to free JSValues with random temparary ones
func (s *SolidNativeMobile) setNodeProp(nodeId string, key string, value *JSValue) error {
	nodeContainer, exists := s.nodeContainers[nodeId]

	if !exists {
		return fmt.Errorf("node does not exist with id %v", nodeId)
	}

	prevStyleMap := nodeContainer.styleMap

	// Host Receiver will take in new JSValue for usage
	// However, the view doesn't update until we use the
	// `OnUpdateRevisionCount` method
	s.hostReceiver.OnPropUpdated(nodeId, key, value)

	// Update flex style and notify of new layout metrics
	if key == "style" {
		newStyleMap := s.convertJSToKeysAndObjects(value)

		updateNodeStyleAndReturnNewStyleKeys(nodeContainer.yogaNode, newStyleMap, prevStyleMap)

		nodeContainer.styleMap = newStyleMap

		// TODO: Update the text styles here.
		if nodeContainer.isText {
			s.updateHostOfTextDescriptor(nodeContainer)
		}

		// Call the layout function, which will update the layout metrics and send it over
		// to the host. It will also notify dirty yoga nodes and update all the
		// revision counts needed.
		s.updateLayoutAndNotify()
		return nil
	}

	if key == "text" && nodeContainer.isText {
		s.updateHostOfTextDescriptor(nodeContainer)
		s.updateLayoutAndNotify()
		return nil
	}

	s.hostReceiver.OnUpdateRevisionCount(nodeId)
	return nil
}

// ===================  UI Mutation =======================

// Anchor is optional.
func (s *SolidNativeMobile) insertBefore(parentId string, newNodeId string, anchorId string) {
	// If its a text node we need special handling.
	// If the parent is a text node, do NOT attach a Yoga node. but DO keep the node in the tree.
	// Also, we do NOT need to

	// If there's an anchor, insert before the anchor

	parentNodeContainer := s.nodeContainers[parentId]
	newNodeContainer := s.nodeContainers[newNodeId]

	// Init to nil
	var newChildren []*NodeContainer

	currentParentChildren := parentNodeContainer.children
	parentYogaNode := parentNodeContainer.yogaNode
	newYogaNode := newNodeContainer.yogaNode

	// Will not insert yoga node if its a text component
	insertChildYogaNode := func(insertionIndex int) {
		if parentNodeContainer.isText {
			return
		}
		parentYogaNode.InsertChild(newYogaNode, insertionIndex)
	}

	if anchorId != "" {
		for i, n := range currentParentChildren {
			if n.id == anchorId {
				insertChildYogaNode(i)
				newChildren = append(newChildren, newNodeContainer, n)
			} else {
				newChildren = append(newChildren, n)
			}
		}
	} else {
		// Add to the end
		newChildren = append(newChildren, currentParentChildren...)
		newChildren = append(newChildren, newNodeContainer)
		ind := len(currentParentChildren)

		insertChildYogaNode(ind)
	}

	// Update internal children
	parentNodeContainer.children = newChildren
	newNodeContainer.parent = parentNodeContainer

	if !parentNodeContainer.isText {
		newChildrenIds := make([]string, 0, len(newChildren))

		for _, n := range newChildren {
			newChildrenIds = append(newChildrenIds, n.id)
		}

		s.hostReceiver.OnChildrenChange(parentId, &StringArray{
			values: newChildrenIds,
		})
	} else {
		s.updateHostOfTextDescriptor(parentNodeContainer)
	}

	s.updateLayoutAndNotify()
}

func (s *SolidNativeMobile) removeChild(parentId string, childNodeId string) {
	parentNodeContainer := s.nodeContainers[parentId]
	childNodeContainer := s.nodeContainers[childNodeId]

	parentChildren := parentNodeContainer.children
	newChildren := make([]*NodeContainer, 0, len(parentChildren)-1)
	newChildIds := make([]string, 0, len(parentChildren)-1)

	for _, n := range parentChildren {
		if n.id == childNodeId {
			continue
		}
		newChildIds = append(newChildIds, n.id)
		newChildren = append(newChildren, n)
	}

	parentNodeContainer.children = newChildren

	// Cleanup on Yoga
	parentYogaNode := parentNodeContainer.yogaNode
	childYogaNode := childNodeContainer.yogaNode

	// If parent is a text node, dont worry about removing the child, it doesn't exist
	// (see insertBefore)
	if !parentNodeContainer.isText {
		parentYogaNode.RemoveChild(childYogaNode)
	}

	delete(s.nodeContainers, childNodeId)

	childYogaNode.Free()

	// Need to:
	// - Update the text descriptors (if that parent is a text)
	// - Notify the host that the children have changed and the node is removed. (parent)
	//		- NOTE: If the parent is a text node and its parent is not, that means that
	//				Does not have children, so no notification is necessary.
	//				So, basically if its a

	if !parentNodeContainer.isText {
		s.hostReceiver.OnChildrenChange(parentId, &StringArray{
			values: newChildIds,
		})
		s.hostReceiver.OnNodeRemoved(childNodeId)
	} else {
		// Same as
		s.updateHostOfTextDescriptor(parentNodeContainer)
	}

	s.updateLayoutAndNotify()
}

// Returns parentId and whether or not it exists
func (s *SolidNativeMobile) getParent(nodeId string) (string, bool) {

	nodeContainer := s.nodeContainers[nodeId]

	if nodeContainer.parent == nil {
		return "", false
	}

	return nodeContainer.parent.id, true
}

// ===================  UI Mutation =======================

// Returns first child id and whether or not it exists
func (s *SolidNativeMobile) getFirstChild(nodeId string) (string, bool) {
	nodeChildren := s.nodeContainers[nodeId].children

	length := len(nodeChildren)

	if length == 0 {
		return "", false
	}

	firstChild := nodeChildren[0]

	return firstChild.id, true
}

func (s *SolidNativeMobile) getNextSibling(nodeId string) (string, bool) {
	parentId, exists := s.getParent(nodeId)

	if !exists {
		return "", false
	}

	parentChildren := s.nodeContainers[parentId].children
	parentChildrenIdLength := len(parentChildren)
	childIndex := 0

	for i, n := range parentChildren {
		if n.id == nodeId {
			childIndex = i
			break
		}
	}

	nextSiblingIndex := childIndex + 1

	if nextSiblingIndex >= parentChildrenIdLength {
		return "", false
	}

	return parentChildren[nextSiblingIndex].id, true
}

// ================= Private Helper Methods =========================

// Call after a prop is changed related to layout/style
//
// # Note:
//
// Be sure to mark a node as dirty with the `YGNode.MarkDirty` function
// If you've update something that causes the MeasureFunction to return a different result
func (s *SolidNativeMobile) updateLayoutAndNotify() error {
	if s.rootNodeId == "" {
		return fmt.Errorf("root node does not exist! cannot update layout")
	}
	rootNodeId := s.rootNodeId
	yogaRootNode := s.nodeContainers[rootNodeId].yogaNode

	yogaRootNode.CalculateLayout(s.deviceScreenSize.Width, s.deviceScreenSize.Height, yoga.DirectionLTR)

	s.applyLayout(rootNodeId)

	return nil
}

func (s *SolidNativeMobile) applyLayout(nodeId string) {
	node := s.nodeContainers[nodeId]

	yogaNode := node.yogaNode

	if !yogaNode.GetHasNewLayout() {
		return
	}

	yogaNode.SetHasNewLayout(false)

	s.hostReceiver.OnLayoutChange(nodeId, convertYogaLayoutMetricToSNLayoutMetrics(
		yoga.NewLayoutMetrics(yogaNode),
	))
	s.hostReceiver.OnUpdateRevisionCount(nodeId)

	for _, n := range node.children {
		s.applyLayout(n.id)
	}
}

// "Upwraps" JS Value by enumerating over its keys
// and values. Ensure this is an object, otherwise just return nothing.
func (s *SolidNativeMobile) convertJSToKeysAndObjects(value *JSValue) map[string]JSValue {
	jsValueMap := make(map[string]JSValue)

	internalMap, ok := value.data.(map[string]interface{})

	// Value is undefined, return nothing
	if !ok {
		return jsValueMap
	}

	for key, v := range internalMap {
		jsValueMap[key] = JSValue{
			data: v,
		}
	}

	return jsValueMap
}

func (s *SolidNativeMobile) downloadAndRunJs(url string) error {
	// Make a GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	// Print the retrieved text
	jsToEval := string(body)

	// fmt.Print(jsToEval)

	return s.dukContext.PevalString(jsToEval)
}

// Note that `parentNodeContainer` need be a text node, but does not have to be the parent of the
// text sub tree.
func (s *SolidNativeMobile) updateHostOfTextDescriptor(parentNodeContainer *NodeContainer) {
	// We now need to update the text descriptor on the parent text node (within the text node subview tree)
	textDescriptors, parentNodeInTextSubTree := generateTextDescriptor(parentNodeContainer)
	// Need to mark dirty so Yoga will want to recalculate the layout.
	// Once the Host gets text descriptors it's measure function will return the appropriate size
	// when we calculate the yoga layout
	parentNodeInTextSubTree.yogaNode.MarkDirty()
	s.hostReceiver.OnNodeTextDescriptorsChange(parentNodeInTextSubTree.id, &TextDescriptorArray{
		values: textDescriptors,
	})
}
