package snmobile

import (
	"fmt"
)

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
		styleMap, err := s.convertJSToKeysAndObjects(value)
		if err != nil {
			return err
		}
		// Ensure we free the values afterwords since we only need them to compute
		// the flex styles
		for _, value := range styleMap {
			defer value.Free()
		}

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
