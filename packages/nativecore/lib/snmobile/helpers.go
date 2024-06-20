package snmobile

import (
	"fmt"
	"nativecore/lib/yoga"

	"github.com/google/uuid"
)

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
func (s *SolidNativeMobile) updateLayoutAndNotify(modifiedNodes map[int]struct{}) error {
	if s.rootNodeId == nil {
		return fmt.Errorf("root node does not exist! cannot update layout")
	}
	rootNodeId := *s.rootNodeId
	yogaRootNode := s.yogaNodes[rootNodeId]

	yogaRootNode.CalculateLayout(s.deviceScreenSize.Width, s.deviceScreenSize.Height, yoga.DirectionLTR)

	s.applyLayout(rootNodeId)

	return nil
}

func (s *SolidNativeMobile) applyLayout(nodeId int) {
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
func (s *SolidNativeMobile) convertJSToKeysAndObjects(value *JSValue) (map[string]JSValue, error) {
	// Check whether its an object or array:
	s.dukContext.PushGlobalStash() // Push stash => [ stash ]

	s.dukContext.GetPropString(-1, value.stashKeyName) // => [ stash value ]

	valueType := s.dukContext.GetType(-1) // => [ stash value ]

	if !valueType.IsObject() {
		return nil, fmt.Errorf("js value with key %s is not an object or does not exist", value.stashKeyName)
	}

	jsValueMap := make(map[string]JSValue)

	// TODO: Check if 0 works or if i have to put in a flag
	s.dukContext.Enum(-1, 0) // => [ stash value enum ]

	for s.dukContext.Next(-1, true) {
		// => [ ... enum key value ]
		key := s.dukContext.GetString(-2)
		keyStashName := uuid.New().String()

		s.dukContext.PushGlobalStash() // => [ ... enum key value stash ]

		s.dukContext.Replace(-3) // => [ ... enum stash value ]

		s.dukContext.PutPropString(-2, keyStashName) // => [ ... enum stash ]

		s.dukContext.Pop() // => [ ... enum ]

		jsValueMap[key] = *NewJsValue(valueType, keyStashName, s)
	}

	return jsValueMap, nil
}

func (s *SolidNativeMobile) downloadAndRunJs() {
	s.dukContext.EvalString("globalThis._SolidNativeRenderer.createNodeByName('sn_view')")
}
