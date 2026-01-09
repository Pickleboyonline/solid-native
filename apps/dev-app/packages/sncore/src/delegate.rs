use crate::jsvalue::JSValue;

/// Trait for receiving callbacks from the renderer to the host platform
#[uniffi::export(with_foreign)]
pub trait HostDelegate: Send + Sync {
    /// When JS creates a node (or even the Mobile side)
    /// this callback is executed
    fn on_node_created(&self, node_id: String, node_type: String);

    /// Will/MUST be called after the children change and are notified
    fn on_node_removed(&self, node_id: String);

    /// Called when a property is updated on a node
    /// The value is a JSValue that can represent different JavaScript types
    fn on_prop_updated(&self, node_id: String, key: String, value: JSValue);

    /// Called when children of a node change
    fn on_children_change(&self, node_id: String, node_ids: Vec<String>);

    // fn on_node_text_descriptors_change(&self, node_id: &str, text_descriptors: &[TextDescriptor]);

    /// Signifies when its time to update JetpackCompose/SwiftUI
    fn on_update_revision_count(&self, node_id: String);

    /// TODO: Change to node type instead
    fn is_text_element_by_node_id(&self, node_id: String) -> bool;

    /// TODO: Change to node type instead
    fn is_text_element_by_node_type(&self, node_type: String) -> bool;
}
