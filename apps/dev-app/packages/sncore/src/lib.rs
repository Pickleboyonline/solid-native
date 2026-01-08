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

// Supporting types for HostDelegate
#[derive(Debug, Clone)]
pub struct Size {
    pub width: f64,
    pub height: f64,
}

#[derive(Debug, Clone)]
pub struct SizeMode {
    pub width_mode: MeasureMode,
    pub height_mode: MeasureMode,
}

#[derive(Debug, Clone)]
pub enum MeasureMode {
    Undefined,
    Exactly,
    AtMost,
}

#[derive(Debug, Clone)]
pub struct LayoutMetrics {
    pub x: f64,
    pub y: f64,
    pub width: f64,
    pub height: f64,
}

#[derive(Debug, Clone)]
pub enum JSValue {
    String(String),
    Number(f64),
    Boolean(bool),
    Null,
    // Add more variants as needed
}

#[derive(Debug, Clone)]
pub struct TextDescriptor {
    pub text: String,
    pub start: usize,
    pub end: usize,
    // Add more fields as needed (font, color, etc.)
}

pub trait HostDelegate {
    /// When JS creates a node (or even the Mobile side)
    /// this callback is executed
    fn on_node_created(&self, node_id: &str, node_type: &str);

    /// Will/MUST be called after the children change and are notified
    fn on_node_removed(&self, node_id: &str);

    fn on_prop_updated(&self, node_id: &str, key: &str, value: &JSValue);

    /// TODO: Determine how to send the data over.
    /// Can work with bytes, but need to determine the size of the int
    /// to effectively decode it.
    fn on_children_change(&self, node_id: &str, node_ids: &[String]);

    fn on_node_text_descriptors_change(&self, node_id: &str, text_descriptors: &[TextDescriptor]);

    /// Signifies when its time to update JetpackCompose/SwiftUI
    fn on_update_revision_count(&self, node_id: &str);

    /// TODO: Change to node type instead
    fn is_text_element_by_node_id(&self, node_id: &str) -> bool;

    /// TODO: Change to node type instead
    fn is_text_element_by_node_type(&self, node_type: &str) -> bool;
}

/*
What's needed:
- 1 SNCore object that handles the JS engine.
- It should be able to take a reference to the "receiver" object
- We are using the delegate pattern here
- The host platform receives function calls and data
*/
