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
