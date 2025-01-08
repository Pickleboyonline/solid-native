package renderer

import (
	"encoding/json"
	"fmt"
	"log"
	"nativecore/lib/core"
	"nativecore/lib/yoga"

	"github.com/buke/quickjs-go"
)

type Renderer struct {
	nodeContainers   map[string]*NodeContainer
	hostReceiver     HostReceiver
	rootNodeId       string
	deviceScreenSize *Size
}

func NewRenderer(hostReceiver HostReceiver) *Renderer {

	return &Renderer{
		hostReceiver:     hostReceiver,
		rootNodeId:       "",
		deviceScreenSize: hostReceiver.GetDeviceScreenSize(),
		nodeContainers:   map[string]*NodeContainer{},
	}
}

// TODO: determine how this is implimneted and if i can just update the device dimensions
func (r *Renderer) OnOrientationChange() {

}

func (r *Renderer) Define(ctx *quickjs.Context) *core.ModuleDefinition {
	module := ctx.Object()

	addGoFunc := func(name string, fn func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value) {
		module.Set(name, ctx.Function(fn))
	}

	addGoFunc("createNodeByName", func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		nodeType := args[0].String()
		nodeId := r.createNode(nodeType)

		log.Printf("New Node create of type %v with id %v", nodeType, nodeId)
		return ctx.String(nodeId)
	})

	addGoFunc("setProp", func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		nodeId := args[0].String()
		key := args[1].String()
		// TODO: Determine if this is a function, which would mean a callback is needed for that key.
		value := args[2]

		if value.IsFunction() {
			log.Printf("Key \"%v\" for node %v was set with a function. Skipping as not supported...", key, nodeId)
			return ctx.Undefined()
		}

		if value.IsUndefined() {
			r.setNodeProp(nodeId, key, &JSValue{})
			return ctx.Undefined()
		}

		encodedJson := value.JSONStringify()

		var unmarshaledValue interface{}

		json.Unmarshal([]byte(encodedJson), &unmarshaledValue)

		log.Printf("Json of prop: %v", unmarshaledValue)

		r.setNodeProp(nodeId, key,
			&JSValue{
				data: unmarshaledValue,
			})

		return ctx.Undefined()
	})

	addGoFunc("insertBefore", func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		parentId := args[0].String()
		nodeId := args[1].String()
		// This may throw, not sure if it does or not
		anchorIdValue := args[2]

		var anchorId string

		if anchorIdValue.IsString() {
			anchorId = anchorIdValue.String()
		}

		log.Printf("Node Insert called with child %v  to be inserted under parent %v", nodeId, parentId)
		r.insertBefore(parentId, nodeId, anchorId)
		log.Printf("Node %v inserted under parent %v", nodeId, parentId)

		return ctx.Undefined()
	})

	addGoFunc("isTextElement", func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		nodeId := args[0].String()

		isTextElement := r.hostReceiver.IsTextElementByNodeId(nodeId)

		return ctx.Bool(isTextElement)
	})

	addGoFunc("removeChild", func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		parentId := args[0].String()
		nodeId := args[1].String()
		r.removeChild(parentId, nodeId)

		return ctx.Undefined()
	})

	addGoFunc("getParent", func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		nodeId := args[0].String()

		parentId, exists := r.getParent(nodeId)

		if !exists {
			return ctx.Undefined()
		}

		return ctx.String(parentId)
	})

	addGoFunc("getFirstChild", func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		nodeId := args[0].String()

		firstChildId, exists := r.getFirstChild(nodeId)

		if !exists {
			return ctx.Undefined()
		}

		return ctx.String(firstChildId)
	})

	addGoFunc("getNextSibling", func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		nodeId := args[0].String()

		nextSiblingId, exists := r.getNextSibling(nodeId)

		if !exists {
			return ctx.Undefined()
		}

		return ctx.String(nextSiblingId)
	})

	addGoFunc("getRootView", func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		return ctx.String(r.rootNodeId)
	})

	return &core.ModuleDefinition{
		Name:  "SNRenderer",
		Value: module,
	}
}

// Create root node and return its ID.
// Not to be called on JS
// This removes the callback like effect and allows the host to create its root node immediatly
// to present it to the screen.
//
// Use the nodetype to tell whether we need to measure it or not
func (r *Renderer) CreateRootNode(nodeType string) string {
	nodeContainer := r.createNodeAndDoNotNotifyHost(nodeType)

	r.rootNodeId = nodeContainer.id

	yogaNode := nodeContainer.yogaNode

	// Ensure proper height/width
	yogaNode.SetWidth(r.deviceScreenSize.Width)
	yogaNode.SetHeight(r.deviceScreenSize.Height)

	return nodeContainer.id
}

// Returns first child id and whether or not it exists
func (r *Renderer) getFirstChild(nodeId string) (string, bool) {
	nodeChildren := r.nodeContainers[nodeId].children

	length := len(nodeChildren)

	if length == 0 {
		return "", false
	}

	firstChild := nodeChildren[0]

	return firstChild.id, true
}

func (r *Renderer) getNextSibling(nodeId string) (string, bool) {
	parentId, exists := r.getParent(nodeId)

	if !exists {
		return "", false
	}

	parentChildren := r.nodeContainers[parentId].children
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

// Anchor is optional.
func (r *Renderer) insertBefore(parentId string, newNodeId string, anchorId string) {
	// If its a text node we need special handling.
	// If the parent is a text node, do NOT attach a Yoga node. but DO keep the node in the tree.
	// Also, we do NOT need to

	// If there's an anchor, insert before the anchor

	parentNodeContainer := r.nodeContainers[parentId]
	newNodeContainer := r.nodeContainers[newNodeId]

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

		r.hostReceiver.OnChildrenChange(parentId, &core.StringArray{
			Values: newChildrenIds,
		})
	} else {
		r.updateHostOfTextDescriptor(parentNodeContainer)
	}

	r.updateLayoutAndNotify()
}

func (r *Renderer) removeChild(parentId string, childNodeId string) {
	parentNodeContainer := r.nodeContainers[parentId]
	childNodeContainer := r.nodeContainers[childNodeId]

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

	delete(r.nodeContainers, childNodeId)

	childYogaNode.Free()

	// Need to:
	// - Update the text descriptors (if that parent is a text)
	// - Notify the host that the children have changed and the node is removed. (parent)
	//		- NOTE: If the parent is a text node and its parent is not, that means that
	//				Does not have children, so no notification is necessary.
	//				So, basically if its a

	if !parentNodeContainer.isText {
		r.hostReceiver.OnChildrenChange(parentId, &core.StringArray{
			Values: newChildIds,
		})
		r.hostReceiver.OnNodeRemoved(childNodeId)
	} else {
		// Same as
		r.updateHostOfTextDescriptor(parentNodeContainer)
	}

	r.updateLayoutAndNotify()
}

// Returns parentId and whether or not it exists
func (r *Renderer) getParent(nodeId string) (string, bool) {

	nodeContainer := r.nodeContainers[nodeId]

	if nodeContainer.parent == nil {
		return "", false
	}

	return nodeContainer.parent.id, true
}

// Updates the host receiver about the props from the JS side
// Value can be a JSValue
// or primative.
// JS Value can be array
// The old JS value associatted does not need to be freed because it has
// a hashed ID. You only need to free JSValues with random temparary ones
func (r *Renderer) setNodeProp(nodeId string, key string, value *JSValue) error {
	nodeContainer, exists := r.nodeContainers[nodeId]

	if !exists {
		return fmt.Errorf("node does not exist with id %v", nodeId)
	}

	prevStyleMap := nodeContainer.styleMap

	// Host Receiver will take in new JSValue for usage
	// However, the view doesn't update until we use the
	// `OnUpdateRevisionCount` method
	r.hostReceiver.OnPropUpdated(nodeId, key, value)

	// Update flex style and notify of new layout metrics
	if key == "style" {
		newStyleMap := r.convertJSToKeysAndObjects(value)

		updateNodeStyleAndReturnNewStyleKeys(nodeContainer.yogaNode, newStyleMap, prevStyleMap)

		nodeContainer.styleMap = newStyleMap

		// TODO: Update the text styles here.
		if nodeContainer.isText {
			r.updateHostOfTextDescriptor(nodeContainer)
		}

		// Call the layout function, which will update the layout metrics and send it over
		// to the host. It will also notify dirty yoga nodes and update all the
		// revision counts needed.
		r.updateLayoutAndNotify()
		return nil
	}

	if key == "text" && nodeContainer.isText {
		r.updateHostOfTextDescriptor(nodeContainer)
		r.updateLayoutAndNotify()
		return nil
	}

	r.hostReceiver.OnUpdateRevisionCount(nodeId)
	return nil
}

// Creates node and notifies mobile host reciever
// to be typically called from JS side.
// Returns Node ID (which is an int)
func (r *Renderer) createNode(nodeType string) string {
	nodeContainer := r.createNodeAndDoNotNotifyHost(nodeType)
	r.hostReceiver.OnNodeCreated(nodeContainer.id, nodeType)
	return nodeContainer.id
}

// Internal usage. Internally, we do not need to keep track of the node type
// Will update NodeContainer map
// TODO: But i do need some mechanism for the measure function
func (r *Renderer) createNodeAndDoNotNotifyHost(nodeType string) *NodeContainer {
	isText := r.hostReceiver.IsTextElementByNodeType(nodeType)
	nodeContainer := newNodeContainer(isText)

	// TODO: check if node type needs measure function.

	needsMeasureFunction := r.hostReceiver.DoesNodeRequireMeasuring(nodeType)

	if needsMeasureFunction {
		nodeContainer.yogaNode.SetMeasureFunc(
			func(node *yoga.YGNode, width float32, widthMode yoga.MeasureMode, height float32, heightMode yoga.MeasureMode) yoga.Size {
				newSize := r.hostReceiver.MeasureNode(nodeContainer.id, NewSize(0, 0), &SizeMode{})
				return yoga.Size{
					Width:  newSize.Width,
					Height: newSize.Height,
				}
			})
	}

	nodeReference := &nodeContainer

	r.nodeContainers[nodeReference.id] = nodeReference

	return nodeReference
}

// Call after a prop is changed related to layout/style
//
// # Note:
//
// Be sure to mark a node as dirty with the `YGNode.MarkDirty` function
// If you've update something that causes the MeasureFunction to return a different result
func (r *Renderer) updateLayoutAndNotify() error {
	if r.rootNodeId == "" {
		return fmt.Errorf("root node does not exist! cannot update layout")
	}
	rootNodeId := r.rootNodeId
	yogaRootNode := r.nodeContainers[rootNodeId].yogaNode

	yogaRootNode.CalculateLayout(r.deviceScreenSize.Width, r.deviceScreenSize.Height, yoga.DirectionLTR)

	r.applyLayout(rootNodeId)

	return nil
}

func (r *Renderer) applyLayout(nodeId string) {
	node := r.nodeContainers[nodeId]

	yogaNode := node.yogaNode

	if !yogaNode.GetHasNewLayout() {
		return
	}

	yogaNode.SetHasNewLayout(false)

	r.hostReceiver.OnLayoutChange(nodeId, convertYogaLayoutMetricToSNLayoutMetrics(
		yoga.NewLayoutMetrics(yogaNode),
	))
	r.hostReceiver.OnUpdateRevisionCount(nodeId)

	for _, n := range node.children {
		r.applyLayout(n.id)
	}
}

// "Upwraps" JS Value by enumerating over its keys
// and values. Ensure this is an object, otherwise just return nothing.
func (r *Renderer) convertJSToKeysAndObjects(value *JSValue) map[string]JSValue {
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

// Note that `parentNodeContainer` need be a text node, but does not have to be the parent of the
// text sub tree.
func (r *Renderer) updateHostOfTextDescriptor(parentNodeContainer *NodeContainer) {
	// We now need to update the text descriptor on the parent text node (within the text node subview tree)
	textDescriptors, parentNodeInTextSubTree := generateTextDescriptor(parentNodeContainer)
	// Need to mark dirty so Yoga will want to recalculate the layout.
	// Once the Host gets text descriptors it's measure function will return the appropriate size
	// when we calculate the yoga layout
	parentNodeInTextSubTree.yogaNode.MarkDirty()
	r.hostReceiver.OnNodeTextDescriptorsChange(parentNodeInTextSubTree.id, &TextDescriptorArray{
		values: textDescriptors,
	})
}
