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
	nodeChildren  map[string][]string
	yogaNodes     map[string]*yoga.YGNode
	nodeStyleKeys map[string]Set
	nodeParent    map[string]string
	hostReceiver  HostReceiver
	dukContext    *duktape.Context
	// Set to -1 initally, need to set before calulcating layouts
	rootNodeId       string
	deviceScreenSize *Size
}

func NewSolidNativeMobile(hostReceiver HostReceiver) *SolidNativeMobile {
	ctx := duktape.New()

	// ctx.PushGoFunction()
	return &SolidNativeMobile{
		yogaNodes:    make(map[string]*yoga.YGNode),
		hostReceiver: hostReceiver,
		dukContext:   ctx,
		rootNodeId:   "",
		nodeParent:   make(map[string]string),
		nodeChildren: make(map[string][]string),
		// We use this to keep track of when keys are removed
		// Since Yoga works via mutation
		nodeStyleKeys:    make(map[string]Set),
		deviceScreenSize: hostReceiver.GetDeviceScreenSize(),
	}
}

// TODO: Give iterator type to retreive all
// native modules with reciever function.
// May need to use flex for type conversion here.
func (s *SolidNativeMobile) RegistureModules() {
	s.dukContext.PushTimers()
	s.dukContext.PushGlobalGoFunction("log", func(c *duktape.Context) int {
		fmt.Println(c.SafeToString(-1))
		return 0
	})
	s.registureRenderer()
}

// Registure core into system, download, and runs js.
func (s *SolidNativeMobile) RunJsFromServer(url string) error {
	s.RegistureModules()
	return s.downloadAndRunJs(url)
}

// Registure core into system, download, and runs js.
func (s *SolidNativeMobile) EvalJs(jsToEval string) error {
	s.RegistureModules()
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
	nodeId := s.createNodeAndDoNotNotifyHost(nodeType)

	s.rootNodeId = nodeId

	yogaNode := s.yogaNodes[nodeId]

	// Ensure proper height/width
	yogaNode.SetWidth(s.deviceScreenSize.Width)
	yogaNode.SetHeight(s.deviceScreenSize.Height)

	return nodeId
}

// Internal function to setup the rendeder
//   - How to throw errors: https://duktape.org/api.html#concepts.9
func (s *SolidNativeMobile) registureRenderer() {
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
		isTextElement := s.hostReceiver.IsTextElement(nodeId)

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
	nodeId := s.createNodeAndDoNotNotifyHost(nodeType)
	s.hostReceiver.OnNodeCreated(nodeId, nodeType)
	return nodeId
}

// Updates the host receiver about the props from the JS side
// Value can be a JSValue
// or primative.
// JS Value can be array
// The old JS value associatted does not need to be freed because it has
// a hashed ID. You only need to free JSValues with random temparary ones
func (s *SolidNativeMobile) setNodeProp(nodeId string, key string, value *JSValue) error {
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
		styleMap := s.convertJSToKeysAndObjects(value)

		newStyleKeys := updateNodeStyleAndReturnNewStyleKeys(node, styleMap, prevKeys)

		s.nodeStyleKeys[nodeId] = newStyleKeys

		// Call the layout function, which will update the layout metrics and send it over
		// to the host. It will also notify dirty yoga nodes and update all the
		// revision counts needed.
		s.updateLayoutAndNotify(map[string]struct{}{
			nodeId: {},
		})
		return nil
	}

	s.hostReceiver.OnUpdateRevisionCount(nodeId)
	return nil
}

// Anchor is optional.
func (s *SolidNativeMobile) insertBefore(parentId string, newNodeId string, anchorId string) {
	// If there's an anchor, insert before the anchor

	// Init to nil
	var newChildrenIds []string
	currentChildrenIds := s.nodeChildren[parentId]
	parentYogaNode := s.yogaNodes[parentId]
	newYogaNode := s.yogaNodes[newNodeId]

	if anchorId != "" {
		for ind, nodeId := range currentChildrenIds {
			if nodeId == anchorId {
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
	s.hostReceiver.OnChildrenChange(parentId, &StringArray{
		strings: newChildrenIds,
	})

	s.updateLayoutAndNotify(map[string]struct{}{})
}

func (s *SolidNativeMobile) removeChild(parentId string, childNodeId string) {
	parentChildIds := s.nodeChildren[parentId]
	newChildIds := make([]string, 0)

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

	s.updateLayoutAndNotify(map[string]struct{}{})
}

func (s *SolidNativeMobile) getParent(nodeId string) (parentId string, exists bool) {
	parentId, exists = s.nodeParent[nodeId]
	return parentId, exists
}

func (s *SolidNativeMobile) getFirstChild(nodeId string) (firstChildId string, exists bool) {
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

func (s *SolidNativeMobile) getNextSibling(nodeId string) (string, bool) {
	parentId, exists := s.getParent(nodeId)

	if !exists {
		return "", false
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

	nextSiblingIndex := childIndex + 1

	if nextSiblingIndex >= parentChildrenIdLength {
		return "", false
	}

	return parentChildrenIds[nextSiblingIndex], true
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
func (s *SolidNativeMobile) updateLayoutAndNotify(modifiedNodes map[string]struct{}) error {
	if s.rootNodeId == "" {
		return fmt.Errorf("root node does not exist! cannot update layout")
	}
	rootNodeId := s.rootNodeId
	yogaRootNode := s.yogaNodes[rootNodeId]

	yogaRootNode.CalculateLayout(s.deviceScreenSize.Width, s.deviceScreenSize.Height, yoga.DirectionLTR)

	s.applyLayout(rootNodeId)

	return nil
}

func (s *SolidNativeMobile) applyLayout(nodeId string) {
	node := s.yogaNodes[nodeId]

	if !node.GetHasNewLayout() {
		return
	}

	node.SetHasNewLayout(false)

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
