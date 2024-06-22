package snmobile

import (
	"fmt"
	"io"
	"nativecore/lib/yoga"
	"net/http"
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
