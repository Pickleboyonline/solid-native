use rquickjs::{Context, Runtime};

uniffi::setup_scaffolding!();

/// Uses JS Engine for it.
#[uniffi::export]
fn add(a: u32, b: u32) -> u32 {
    let runtime = Runtime::new().unwrap();
    let context = Context::full(&runtime).unwrap();

    let expression = format!("{} + {} + 1", a, b);

    let result = context.with(|ctx| ctx.eval::<u32, _>(expression).unwrap());
    result
}

pub trait HostDelegate {
    fn on_node_created(node_id: &str);
}

/*
What's needed:
- 1 SNCore object that handles the JS engine.
- It should be able to take a reference to the "receiver" object
- We are using the delegate pattern here
- The host platform receives function calls and data
*/

/*
Here was the Go implementation:
// Wrapper object that has callbacks to host platform methods
type HostReceiver interface {
	// When JS creates a node (or even the Mobile side)
	// this callback is executed
	OnNodeCreated(nodeId string, nodeType string)
	// Will/MUST be called after the children change and are notified
	OnNodeRemoved(nodeId string)

	// Some nodes, like text & text input, need to be
	// measured while calculating layout before
	// sending it over the wire
	DoesNodeRequireMeasuring(nodeType string) bool

	// TODO: See if we need any other info to make measure call
	MeasureNode(nodeId string, proposedSize *Size, sizeMode *SizeMode) *Size

	// Need this to setup root node and calculate layout.
	GetDeviceScreenSize() *Size

	OnLayoutChange(nodeId string, layoutMetrics *LayoutMetrics)
	OnPropUpdated(nodeId string, key string, value *JSValue)

	// TODO: Determine how to send the data over.
	// Can work with bytes, but need to determine the size of the int
	// to effectivly decode it.
	OnChildrenChange(nodeId string, nodeIds *core.StringArray)

	OnNodeTextDescriptorsChange(nodeId string, textDescriptors *TextDescriptorArray)

	// Signifies when its time to update JetpackCompose/SwiftUI
	OnUpdateRevisionCount(nodeId string)
	// TODO: Change to node type instead
	IsTextElementByNodeId(nodeId string) bool
	// TODO: Change to node type instead
	IsTextElementByNodeType(nodeType string) bool
}

*/