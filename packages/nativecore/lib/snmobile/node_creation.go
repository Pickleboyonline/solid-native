package snmobile

import (
	"nativecore/lib/yoga"

	"github.com/google/uuid"
)

// Used to access arrays since gomobile can't
// expose arrays other than bytes
// TODO: Understand how memory management works between Go/Mobile
type StringArray struct {
	strings []string
}

func (e *StringArray) Length() int {
	if e.strings == nil {
		return 0
	}
	return len(e.strings)
}

func (e *StringArray) Get(index int) string {
	return e.strings[index]
}

// Internal usage. Internally, we do not need to keep track of the node type
// TODO: But i do need some mechanism for the measure function
func (s *SolidNativeMobile) createNodeAndDoNotNotifyHost(nodeType string) string {
	id := uuid.New().String()
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
	s.nodeChildren[id] = make([]string, 0)
	s.nodeStyleKeys[id] = make(Set)
	return id
}
