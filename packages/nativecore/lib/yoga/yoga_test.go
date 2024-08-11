package yoga

import (
	"log"
	"testing"
)

func TestNodeMeasurement(t *testing.T) {
	node := NewNode()
	defer node.Free()

	child := NewNode()
	defer child.Free()
	didCallMeasureFunction := false

	child.SetMeasureFunc(func(node *YGNode, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
		log.Println("Called Measure function!")
		didCallMeasureFunction = true
		return Size{}
	})

	node.InsertChild(child, 0)

	node.CalculateLayout(100, 100, DirectionLTR)

	if !didCallMeasureFunction {
		t.Fatalf("Fail: Node did not call measure function!")
	}
}
