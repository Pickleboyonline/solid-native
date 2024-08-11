package yoga

/*
#include <yoga/Yoga.h>

extern YGSize goMeasureCallback(
	YGNodeRef node,
	float width,
	YGMeasureMode widthMode,
	float height,
	YGMeasureMode heightMode
);
*/
import "C"

// MeasureFunc is the type for the measure function callback.
type MeasureFunc func(node *YGNode, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size

type measureFuncRegistryValue struct {
	measureFunc MeasureFunc
	yogaNode    *YGNode
}

var measureFuncRegistry = make(map[C.YGNodeRef]measureFuncRegistryValue)

//export goMeasureCallback
func goMeasureCallback(node C.YGNodeRef, width C.float, widthMode C.YGMeasureMode, height C.float, heightMode C.YGMeasureMode) C.YGSize {
	if result, ok := measureFuncRegistry[node]; ok {
		size := result.measureFunc(result.yogaNode, float32(width), MeasureMode(widthMode), float32(height), MeasureMode(heightMode))
		return C.YGSize{width: C.float(size.Width), height: C.float(size.Height)}
	}
	return C.YGSize{width: 0, height: 0}
}

// SetMeasureFunc sets a custom measure function for the Yoga node.
func (n *YGNode) SetMeasureFunc(f MeasureFunc) {
	if f != nil {
		measureFuncRegistry[n.node] = measureFuncRegistryValue{
			measureFunc: f,
			yogaNode:    n,
		}
		C.YGNodeSetMeasureFunc(n.node, (C.YGMeasureFunc)(C.goMeasureCallback))
	} else {
		delete(measureFuncRegistry, n.node)
		C.YGNodeSetMeasureFunc(n.node, nil)
	}
}
