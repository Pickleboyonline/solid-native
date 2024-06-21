package snmobile

import (
	"log"
	"testing"
)

type MockHostReceiver struct{}

func (m *MockHostReceiver) OnNodeCreated(nodeId string, nodeType string)               {}
func (m *MockHostReceiver) DoesNodeRequireMeasuring(nodeType string) bool              { return false }
func (m *MockHostReceiver) MeasureNode(nodeId string) *Size                            { return &Size{Width: 100, Height: 100} }
func (m *MockHostReceiver) GetDeviceScreenSize() *Size                                 { return &Size{Width: 375, Height: 667} }
func (m *MockHostReceiver) OnLayoutChange(nodeId string, layoutMetrics *LayoutMetrics) {}
func (m *MockHostReceiver) OnPropUpdated(nodeId string, key string, value *JSValue)    {}
func (m *MockHostReceiver) OnChildrenChange(nodeId string, nodeIds *StringArray)       {}
func (m *MockHostReceiver) OnUpdateRevisionCount(nodeId string)                        {}
func (m *MockHostReceiver) IsTextElement(nodeId string) bool                           { return false }

func TestCreateNode(t *testing.T) {
	hostReceiver := &MockHostReceiver{}
	snm := NewSolidNativeMobile(hostReceiver)
	defer snm.FreeMemory()

	id := snm.CreateRootNode("sn_view")

	log.Println("Root node with id ", id)

	snm.RegistureModules()

	err := snm.EvalJs("log('' + _SolidNativeRenderer.getRootView())")

	if err != nil {
		t.Fatalf("jslog error: %v", err)
	}

	err = snm.downloadAndRunJs("http://localhost:8080")

	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func CreateRootNode(t *testing.T) {
	hostReceiver := &MockHostReceiver{}
	snm := NewSolidNativeMobile(hostReceiver)
	defer snm.FreeMemory()

	js := `
	var id = _SolidNativeRenderer.createNodeByName("sn_view");
	if (typeof id !== 'number') {
	throw new Error('id was not created!') 
	}
	`

	err := snm.EvalJs(js)

	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestSetNodeProp(t *testing.T) {
	// hostReceiver := &MockHostReceiver{}
	// snm := NewSolidNativeMobile(hostReceiver)

	// nodeId := snm.CreateNode("testNode")

	// jsValue := NewJsValue(duktape.TypeString, "testValueKey", snm)
	// err := snm.SetNodeProp(nodeId, "style", jsValue)

	// if err != nil {
	// 	t.Fatalf("Setting node property failed: %v", err)
	// }

	// if _, exists := snm.nodeStyleKeys[nodeId]; !exists {
	// 	t.Fatalf("Node style key was not updated for node ID %d", nodeId)
	// }
}

func TestInsertBefore(t *testing.T) {
	hostReceiver := &MockHostReceiver{}
	snm := NewSolidNativeMobile(hostReceiver)
	defer snm.FreeMemory()

	parentId := snm.createNode("parentNode")
	childId := snm.createNode("childNode")

	snm.insertBefore(parentId, childId, "")

	children := snm.nodeChildren[parentId]
	if len(children) != 1 || children[0] != childId {
		t.Fatalf("Node %s was not inserted correctly under parent %s", childId, parentId)
	}
}

func TestRemoveChild(t *testing.T) {
	hostReceiver := &MockHostReceiver{}
	snm := NewSolidNativeMobile(hostReceiver)
	defer snm.FreeMemory()

	parentId := snm.createNode("parentNode")
	childId := snm.createNode("childNode")

	snm.insertBefore(parentId, childId, "")
	snm.removeChild(parentId, childId)

	if _, exists := snm.yogaNodes[childId]; exists {
		t.Fatalf("Child node %s was not removed correctly from parent %s", childId, parentId)
	}
}

func TestUpdateLayoutAndNotify(t *testing.T) {
	hostReceiver := &MockHostReceiver{}
	snm := NewSolidNativeMobile(hostReceiver)
	defer snm.FreeMemory()

	rootNodeId := snm.CreateRootNode("rootNode")
	childId := snm.createNode("childNode")

	snm.insertBefore(rootNodeId, childId, "")

}
