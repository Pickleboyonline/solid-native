package snmobile

import (
	"nativecore/lib/yoga"

	"github.com/google/uuid"
)

// Used to access arrays since gomobile can't
// expose arrays other than bytes
// TODO: Understand how memory management works between Go/Mobile
type IntegerArray struct {
	integers []int
}

func (e *IntegerArray) Length() int {
	if e.integers == nil {
		return 0
	}
	return len(e.integers)
}

func (e *IntegerArray) Get(index int) int {
	return e.integers[index]
}

// Internal usage. Internally, we do not need to keep track of the node type
// TODO: But i do need some mechanism for the measure function
func (s *SolidNativeMobile) createNodeAndDoNotNotifyHost(nodeType string) int {
	id := int(uuid.New().ID())
	yogaNode := yoga.NewNode()
	// TODO: check if node type needs measure function.

	needsMeasureFunction := s.hostReceiver.DoesNodeRequireMeasuring(nodeType)

	if needsMeasureFunction {
		yogaNode.SetMeasureFunc(func(node *yoga.YGNode, width float32, widthMode yoga.MeasureMode, height float32, heightMode yoga.MeasureMode) yoga.Size {
			newSize := s.hostReceiver.MeasureNode(id)
			return yoga.Size{
				Width:  newSize.Width,
				Height: newSize.Height,
			}
		})
	}

	s.yogaNodes[id] = yogaNode
	s.nodeChildren[id] = make([]int, 0)
	s.nodeStyleKeys[id] = make(Set)
	return id
}
