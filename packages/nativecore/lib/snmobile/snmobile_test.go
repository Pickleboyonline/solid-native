package snmobile

import (
	"testing"
)

type MockHostReceiver struct{}

func (m *MockHostReceiver) OnNodeCreated(nodeId int, nodeType string)               {}
func (m *MockHostReceiver) DoesNodeRequireMeasuring(nodeType string) bool           { return false }
func (m *MockHostReceiver) MeasureNode(nodeId int) *Size                            { return &Size{Width: 100, Height: 100} }
func (m *MockHostReceiver) GetDeviceScreenSize() *Size                              { return &Size{Width: 375, Height: 667} }
func (m *MockHostReceiver) OnLayoutChange(nodeId int, layoutMetrics *LayoutMetrics) {}
func (m *MockHostReceiver) OnPropUpdated(nodeId int, key string, value *JSValue)    {}
func (m *MockHostReceiver) OnChildrenChange(nodeId int, nodeIds *IntegerArray)      {}
func (m *MockHostReceiver) OnUpdateRevisionCount(nodeId int)                        {}
func (m *MockHostReceiver) IsTextElement(nodeId int) bool                           { return false }

func TestCreateNode(t *testing.T) {
	hostReceiver := &MockHostReceiver{}
	snm := NewSolidNativeMobile(hostReceiver)
	defer snm.FreeMemory()

	js := `
	var id = _SolidNativeRenderer.createNodeByName("sn_view");
	if (typeof id !== 'number') {
	throw new Error('id was not created!')
	}
	`

	err := snm.RunJs(js)

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

	err := snm.RunJs(js)

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

	snm.insertBefore(parentId, childId, nil)

	children := snm.nodeChildren[parentId]
	if len(children) != 1 || children[0] != childId {
		t.Fatalf("Node %d was not inserted correctly under parent %d", childId, parentId)
	}
}

func TestRemoveChild(t *testing.T) {
	hostReceiver := &MockHostReceiver{}
	snm := NewSolidNativeMobile(hostReceiver)
	defer snm.FreeMemory()

	parentId := snm.createNode("parentNode")
	childId := snm.createNode("childNode")

	snm.insertBefore(parentId, childId, nil)
	snm.removeChild(parentId, childId)

	if _, exists := snm.yogaNodes[childId]; exists {
		t.Fatalf("Child node %d was not removed correctly from parent %d", childId, parentId)
	}
}

func TestUpdateLayoutAndNotify(t *testing.T) {
	hostReceiver := &MockHostReceiver{}
	snm := NewSolidNativeMobile(hostReceiver)
	defer snm.FreeMemory()

	rootNodeId := snm.CreateRootNode("rootNode")
	childId := snm.createNode("childNode")

	snm.insertBefore(rootNodeId, childId, nil)

}
