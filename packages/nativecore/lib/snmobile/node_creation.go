package snmobile

import (
	"nativecore/lib/yoga"
)

// Internal usage. Internally, we do not need to keep track of the node type
// Will update NodeContainer map
// TODO: But i do need some mechanism for the measure function
func (s *SolidNativeMobile) createNodeAndDoNotNotifyHost(nodeType string) *NodeContainer {
	nodeContainer := newNodeContainer()

	// TODO: check if node type needs measure function.

	needsMeasureFunction := s.hostReceiver.DoesNodeRequireMeasuring(nodeType)

	if needsMeasureFunction {
		nodeContainer.yogaNode.SetMeasureFunc(
			func(node *yoga.YGNode, width float32, widthMode yoga.MeasureMode, height float32, heightMode yoga.MeasureMode) yoga.Size {
				newSize := s.hostReceiver.MeasureNode(nodeContainer.id, NewSize(0, 0), &SizeMode{})
				return yoga.Size{
					Width:  newSize.Width,
					Height: newSize.Height,
				}
			})
	}

	return &nodeContainer
}
