package snmobile

// Wrapper object that has callbacks to host platform methods
type HostReceiver interface {
	// When JS creates a node (or even the Mobile side)
	// this callback is executed
	OnNodeCreated(nodeId int, nodeType string)
	// Some nodes, like text & text input, need to be
	// measured while calculating layout before
	// sending it over the wire
	DoesNodeRequireMeasuring(nodeType string) bool

	// TODO: See if we need any other info to make measure call
	MeasureNode(nodeId int) *Size

	// Need this to setup root node and calculate layout.
	GetDeviceScreenSize() *Size

	OnLayoutChange(nodeId int, layoutMetrics *LayoutMetrics)
	OnPropUpdated(nodeId int, key string, value *JSValue)

	// TODO: Determine how to send the data over.
	// Can work with bytes, but need to determine the size of the int
	// to effectivly decode it.
	OnChildrenChange(nodeId int, nodeIds *IntegerArray)
	// Signifies when its time to update JetpackCompose/SwiftUI
	OnUpdateRevisionCount(nodeId int)
	IsTextElement(nodeId int) bool
}
