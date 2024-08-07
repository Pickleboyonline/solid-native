package snmobile

import (
	"nativecore/lib/yoga"

	"github.com/google/uuid"
)

// Internal usage. Internally, we do not need to keep track of the node type
// TODO: But i do need some mechanism for the measure function
func (s *SolidNativeMobile) createNodeAndDoNotNotifyHost(nodeType string) string {
	id := uuid.New().String()
	yogaNode := yoga.NewNode()
	// TODO: check if node type needs measure function.

	needsMeasureFunction := s.hostReceiver.DoesNodeRequireMeasuring(nodeType)

	if needsMeasureFunction {
		yogaNode.SetMeasureFunc(func(node *yoga.YGNode, width float32, widthMode yoga.MeasureMode, height float32, heightMode yoga.MeasureMode) yoga.Size {
			newSize := s.hostReceiver.MeasureNode(id, NewSize(0, 0), &SizeMode{})
			return yoga.Size{
				Width:  newSize.Width,
				Height: newSize.Height,
			}
		})
	}

	s.yogaNodes[id] = yogaNode
	s.nodeChildren[id] = make([]string, 0)
	s.nodeStyleKeys[id] = make(Set)
	return id
}
