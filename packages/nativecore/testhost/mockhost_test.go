package testhost

import (
	"log"
	"nativecore/lib/snmobile"
	"testing"
)

type MockHostReceiver struct{}

func (m *MockHostReceiver) OnNodeCreated(nodeId string, nodeType string)  {}
func (m *MockHostReceiver) OnNodeRemoved(nodeId string)                   {}
func (m *MockHostReceiver) DoesNodeRequireMeasuring(nodeType string) bool { return false }
func (m *MockHostReceiver) MeasureNode(nodeId string, size *snmobile.Size, sizeMode *snmobile.SizeMode) *snmobile.Size {
	return &snmobile.Size{Width: 100, Height: 100}
}
func (m *MockHostReceiver) GetDeviceScreenSize() *snmobile.Size {
	return &snmobile.Size{Width: 375, Height: 667}
}
func (m *MockHostReceiver) OnLayoutChange(nodeId string, layoutMetrics *snmobile.LayoutMetrics) {
	log.Printf("Layout for node id %v changed: %#v", nodeId, *layoutMetrics)
}
func (m *MockHostReceiver) OnPropUpdated(nodeId string, key string, value *snmobile.JSValue) {}
func (m *MockHostReceiver) OnChildrenChange(nodeId string, nodeIds *snmobile.StringArray)    {}
func (m *MockHostReceiver) OnUpdateRevisionCount(nodeId string)                              {}
func (m *MockHostReceiver) IsTextElementByNodeId(nodeId string) bool                         { return false }
func (m *MockHostReceiver) IsTextElementByNodeType(nodeType string) bool                     { return false }

func TestMockHost(t *testing.T) {

}
