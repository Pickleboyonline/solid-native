use crate::delegate::HostDelegate;
use crate::tree::UITree;
use rquickjs::{Context, Runtime};
use std::sync::{Arc, Mutex};

/// The main SolidJS renderer that manages the JS engine and UI tree
pub struct SolidRenderer {
    runtime: Runtime,
    tree: Arc<Mutex<UITree>>,
    delegate: Arc<dyn HostDelegate + Send + Sync>,
    node_counter: Arc<Mutex<usize>>,
}

impl SolidRenderer {
    /// Creates a new SolidRenderer with the given delegate
    pub fn new(delegate: Arc<dyn HostDelegate + Send + Sync>) -> Result<Self, String> {
        let runtime = Runtime::new().map_err(|e| format!("Failed to create runtime: {:?}", e))?;

        Ok(Self {
            runtime,
            tree: Arc::new(Mutex::new(UITree::new())),
            delegate,
            node_counter: Arc::new(Mutex::new(0)),
        })
    }

    /// Generates a unique node ID
    fn generate_node_id(&self) -> String {
        let mut counter = self.node_counter.lock().unwrap();
        *counter += 1;
        format!("node_{}", counter)
    }

    /// Creates an element node
    pub fn create_element(&self, tag: String) -> String {
        let id = self.generate_node_id();
        let mut tree = self.tree.lock().unwrap();
        tree.create_element(id.clone(), tag.clone());

        // Notify delegate
        self.delegate.on_node_created(&id, &tag);

        id
    }

    /// Creates a text node
    pub fn create_text_node(&self, value: String) -> String {
        let id = self.generate_node_id();
        let mut tree = self.tree.lock().unwrap();
        tree.create_text_node(id.clone(), value);

        // Notify delegate
        self.delegate.on_node_created(&id, "text");

        id
    }

    /// Replaces text content of a text node
    pub fn replace_text(&self, node_id: String, value: String) {
        let tree_ref = self.tree.lock().unwrap();
        if let Some(key) = tree_ref.get_key_by_id(&node_id) {
            drop(tree_ref);
            let mut tree = self.tree.lock().unwrap();
            tree.replace_text(key, value);
            self.delegate.on_update_revision_count(&node_id);
        }
    }

    /// Sets a property on a node
    pub fn set_property(&self, node_id: String, name: String, value: String) {
        let tree_ref = self.tree.lock().unwrap();
        if let Some(key) = tree_ref.get_key_by_id(&node_id) {
            drop(tree_ref);
            let mut tree = self.tree.lock().unwrap();
            tree.set_property(key, name, value);
            self.delegate.on_update_revision_count(&node_id);
        }
    }

    /// Inserts a node into a parent before an anchor node
    pub fn insert_node(&self, parent_id: String, node_id: String, anchor_id: Option<String>) {
        let tree_ref = self.tree.lock().unwrap();

        let parent_key = match tree_ref.get_key_by_id(&parent_id) {
            Some(k) => k,
            None => return,
        };

        let node_key = match tree_ref.get_key_by_id(&node_id) {
            Some(k) => k,
            None => return,
        };

        let anchor_key = anchor_id.and_then(|id| tree_ref.get_key_by_id(&id));

        drop(tree_ref);
        let mut tree = self.tree.lock().unwrap();
        tree.insert_node(parent_key, node_key, anchor_key);

        // Get updated children
        let children_ids = tree.get_children_ids(parent_key);
        drop(tree);

        // Notify delegate
        self.delegate.on_children_change(&parent_id, &children_ids.iter().map(|s| s.clone()).collect::<Vec<_>>());
        self.delegate.on_update_revision_count(&parent_id);
    }

    /// Removes a node from its parent
    pub fn remove_node(&self, parent_id: String, node_id: String) {
        let tree_ref = self.tree.lock().unwrap();

        let parent_key = match tree_ref.get_key_by_id(&parent_id) {
            Some(k) => k,
            None => return,
        };

        let node_key = match tree_ref.get_key_by_id(&node_id) {
            Some(k) => k,
            None => return,
        };

        drop(tree_ref);
        let mut tree = self.tree.lock().unwrap();
        tree.remove_node(parent_key, node_key);

        // Get updated children
        let children_ids = tree.get_children_ids(parent_key);
        drop(tree);

        // Notify delegate
        self.delegate.on_node_removed(&node_id);
        self.delegate.on_children_change(&parent_id, &children_ids.iter().map(|s| s.clone()).collect::<Vec<_>>());
        self.delegate.on_update_revision_count(&parent_id);
    }

    /// Checks if a node is a text node
    pub fn is_text_node(&self, node_id: String) -> bool {
        let tree = self.tree.lock().unwrap();
        if let Some(key) = tree.get_key_by_id(&node_id) {
            if let Some(node) = tree.get_node(key) {
                return node.is_text_node();
            }
        }
        false
    }

    /// Gets the parent node ID
    pub fn get_parent_node(&self, node_id: String) -> Option<String> {
        let tree = self.tree.lock().unwrap();
        let key = tree.get_key_by_id(&node_id)?;
        let parent_key = tree.get_parent(key)?;
        let parent = tree.get_node(parent_key)?;
        Some(parent.id.clone())
    }

    /// Gets the first child node ID
    pub fn get_first_child(&self, node_id: String) -> Option<String> {
        let tree = self.tree.lock().unwrap();
        let key = tree.get_key_by_id(&node_id)?;
        let child_key = tree.get_first_child(key)?;
        let child = tree.get_node(child_key)?;
        Some(child.id.clone())
    }

    /// Gets the next sibling node ID
    pub fn get_next_sibling(&self, node_id: String) -> Option<String> {
        let tree = self.tree.lock().unwrap();
        let key = tree.get_key_by_id(&node_id)?;
        let sibling_key = tree.get_next_sibling(key)?;
        let sibling = tree.get_node(sibling_key)?;
        Some(sibling.id.clone())
    }

    /// Evaluates JavaScript code
    pub fn eval_js(&self, code: String) -> Result<String, String> {
        let context = Context::full(&self.runtime)
            .map_err(|e| format!("Failed to create context: {:?}", e))?;

        context.with(|ctx| {
            ctx.eval::<String, _>(code)
                .map_err(|e| format!("JS eval error: {:?}", e))
        })
    }

    /// Gets access to the UI tree (for inspection/debugging)
    pub fn get_tree(&self) -> Arc<Mutex<UITree>> {
        Arc::clone(&self.tree)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::sync::Mutex as StdMutex;

    // Mock delegate for testing
    struct MockDelegate {
        created_nodes: Arc<StdMutex<Vec<(String, String)>>>,
        removed_nodes: Arc<StdMutex<Vec<String>>>,
        children_changes: Arc<StdMutex<Vec<(String, Vec<String>)>>>,
        update_counts: Arc<StdMutex<Vec<String>>>,
    }

    impl MockDelegate {
        fn new() -> Self {
            Self {
                created_nodes: Arc::new(StdMutex::new(Vec::new())),
                removed_nodes: Arc::new(StdMutex::new(Vec::new())),
                children_changes: Arc::new(StdMutex::new(Vec::new())),
                update_counts: Arc::new(StdMutex::new(Vec::new())),
            }
        }
    }

    impl HostDelegate for MockDelegate {
        fn on_node_created(&self, node_id: &str, node_type: &str) {
            self.created_nodes
                .lock()
                .unwrap()
                .push((node_id.to_string(), node_type.to_string()));
        }

        fn on_node_removed(&self, node_id: &str) {
            self.removed_nodes
                .lock()
                .unwrap()
                .push(node_id.to_string());
        }

        fn on_children_change(&self, node_id: &str, node_ids: &[String]) {
            self.children_changes
                .lock()
                .unwrap()
                .push((node_id.to_string(), node_ids.to_vec()));
        }

        fn on_update_revision_count(&self, node_id: &str) {
            self.update_counts
                .lock()
                .unwrap()
                .push(node_id.to_string());
        }

        fn is_text_element_by_node_id(&self, _node_id: &str) -> bool {
            false
        }

        fn is_text_element_by_node_type(&self, node_type: &str) -> bool {
            node_type == "text"
        }
    }

    #[test]
    fn test_create_element() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let node_id = renderer.create_element("div".to_string());

        assert!(node_id.starts_with("node_"));

        let created = delegate.created_nodes.lock().unwrap();
        assert_eq!(created.len(), 1);
        assert_eq!(created[0].0, node_id);
        assert_eq!(created[0].1, "div");
    }

    #[test]
    fn test_create_text_node() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let node_id = renderer.create_text_node("Hello World".to_string());

        assert!(node_id.starts_with("node_"));

        let created = delegate.created_nodes.lock().unwrap();
        assert_eq!(created.len(), 1);
        assert_eq!(created[0].0, node_id);
        assert_eq!(created[0].1, "text");
    }

    #[test]
    fn test_replace_text() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let node_id = renderer.create_text_node("Hello".to_string());
        renderer.replace_text(node_id.clone(), "World".to_string());

        let tree = renderer.get_tree();
        let tree_lock = tree.lock().unwrap();
        let key = tree_lock.get_key_by_id(&node_id).unwrap();
        let node = tree_lock.get_node(key).unwrap();

        match &node.node_type {
            crate::node::NodeType::Text { content } => assert_eq!(content, "World"),
            _ => panic!("Expected text node"),
        }

        let updates = delegate.update_counts.lock().unwrap();
        assert!(updates.contains(&node_id));
    }

    #[test]
    fn test_set_property() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let node_id = renderer.create_element("div".to_string());
        renderer.set_property(node_id.clone(), "className".to_string(), "container".to_string());

        let tree = renderer.get_tree();
        let tree_lock = tree.lock().unwrap();
        let key = tree_lock.get_key_by_id(&node_id).unwrap();
        let node = tree_lock.get_node(key).unwrap();

        assert_eq!(node.properties.get("className"), Some(&"container".to_string()));
    }

    #[test]
    fn test_insert_node() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let parent_id = renderer.create_element("div".to_string());
        let child_id = renderer.create_element("span".to_string());

        renderer.insert_node(parent_id.clone(), child_id.clone(), None);

        let tree = renderer.get_tree();
        let tree_lock = tree.lock().unwrap();
        let parent_key = tree_lock.get_key_by_id(&parent_id).unwrap();
        let parent = tree_lock.get_node(parent_key).unwrap();

        assert_eq!(parent.children.len(), 1);

        let children_changes = delegate.children_changes.lock().unwrap();
        assert!(children_changes.iter().any(|(id, children)| {
            id == &parent_id && children.len() == 1 && children[0] == child_id
        }));
    }

    #[test]
    fn test_remove_node() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let parent_id = renderer.create_element("div".to_string());
        let child_id = renderer.create_element("span".to_string());

        renderer.insert_node(parent_id.clone(), child_id.clone(), None);
        renderer.remove_node(parent_id.clone(), child_id.clone());

        let tree = renderer.get_tree();
        let tree_lock = tree.lock().unwrap();
        let parent_key = tree_lock.get_key_by_id(&parent_id).unwrap();
        let parent = tree_lock.get_node(parent_key).unwrap();

        assert_eq!(parent.children.len(), 0);

        let removed = delegate.removed_nodes.lock().unwrap();
        assert!(removed.contains(&child_id));
    }

    #[test]
    fn test_is_text_node() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let element_id = renderer.create_element("div".to_string());
        let text_id = renderer.create_text_node("Hello".to_string());

        assert!(!renderer.is_text_node(element_id));
        assert!(renderer.is_text_node(text_id));
    }

    #[test]
    fn test_get_parent_node() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let parent_id = renderer.create_element("div".to_string());
        let child_id = renderer.create_element("span".to_string());

        renderer.insert_node(parent_id.clone(), child_id.clone(), None);

        let found_parent = renderer.get_parent_node(child_id);
        assert_eq!(found_parent, Some(parent_id));
    }

    #[test]
    fn test_get_first_child() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let parent_id = renderer.create_element("div".to_string());
        let child1_id = renderer.create_element("span".to_string());
        let child2_id = renderer.create_element("p".to_string());

        renderer.insert_node(parent_id.clone(), child1_id.clone(), None);
        renderer.insert_node(parent_id.clone(), child2_id.clone(), None);

        let first_child = renderer.get_first_child(parent_id);
        assert_eq!(first_child, Some(child1_id));
    }

    #[test]
    fn test_get_next_sibling() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let parent_id = renderer.create_element("div".to_string());
        let child1_id = renderer.create_element("span".to_string());
        let child2_id = renderer.create_element("p".to_string());

        renderer.insert_node(parent_id.clone(), child1_id.clone(), None);
        renderer.insert_node(parent_id.clone(), child2_id.clone(), None);

        let next_sibling = renderer.get_next_sibling(child1_id);
        assert_eq!(next_sibling, Some(child2_id));
    }

    #[test]
    fn test_eval_js() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let result = renderer.eval_js("(2 + 3).toString()".to_string());
        assert_eq!(result, Ok("5".to_string()));
    }

    #[test]
    fn test_node_id_generation() {
        let delegate = Arc::new(MockDelegate::new());
        let renderer = SolidRenderer::new(delegate.clone()).unwrap();

        let id1 = renderer.create_element("div".to_string());
        let id2 = renderer.create_element("span".to_string());
        let id3 = renderer.create_text_node("text".to_string());

        assert_ne!(id1, id2);
        assert_ne!(id2, id3);
        assert_ne!(id1, id3);
    }
}
