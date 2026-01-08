/// Trait for receiving callbacks from the renderer to the host platform
pub trait HostDelegate {
    /// When JS creates a node (or even the Mobile side)
    /// this callback is executed
    fn on_node_created(&self, node_id: &str, node_type: &str);

    /// Will/MUST be called after the children change and are notified
    fn on_node_removed(&self, node_id: &str);

    // TODO: create/determine wrapper for quckjs js types
    // fn on_prop_updated(&self, node_id: &str, key: &str, value: &JSValue);

    /// TODO: Determine how to send the data over.
    /// Can work with bytes, but need to determine the size of the int
    /// to effectively decode it.
    fn on_children_change(&self, node_id: &str, node_ids: &[String]);

    // fn on_node_text_descriptors_change(&self, node_id: &str, text_descriptors: &[TextDescriptor]);

    /// Signifies when its time to update JetpackCompose/SwiftUI
    fn on_update_revision_count(&self, node_id: &str);

    /// TODO: Change to node type instead
    fn is_text_element_by_node_id(&self, node_id: &str) -> bool;

    /// TODO: Change to node type instead
    fn is_text_element_by_node_type(&self, node_type: &str) -> bool;
}
