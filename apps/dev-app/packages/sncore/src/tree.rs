use crate::node::{Node, NodeKey, NodeType};
use slotmap::SlotMap;
use std::collections::HashMap;

/// Manages the UI tree structure using a normalized slotmap
#[derive(Debug)]
pub struct UITree {
    nodes: SlotMap<NodeKey, Node>,
    id_to_key: HashMap<String, NodeKey>,
    root: Option<NodeKey>,
}

impl UITree {
    pub fn new() -> Self {
        Self {
            nodes: SlotMap::new(),
            id_to_key: HashMap::new(),
            root: None,
        }
    }

    /// Creates a new element node
    pub fn create_element(&mut self, id: String, tag: String) -> NodeKey {
        let node = Node::new_element(id.clone(), tag);
        let key = self.nodes.insert(node);
        self.id_to_key.insert(id, key);
        key
    }

    /// Creates a new text node
    pub fn create_text_node(&mut self, id: String, content: String) -> NodeKey {
        let node = Node::new_text(id.clone(), content);
        let key = self.nodes.insert(node);
        self.id_to_key.insert(id, key);
        key
    }

    /// Gets a node by its key
    pub fn get_node(&self, key: NodeKey) -> Option<&Node> {
        self.nodes.get(key)
    }

    /// Gets a mutable node by its key
    pub fn get_node_mut(&mut self, key: NodeKey) -> Option<&mut Node> {
        self.nodes.get_mut(key)
    }

    /// Gets a node key by its ID
    pub fn get_key_by_id(&self, id: &str) -> Option<NodeKey> {
        self.id_to_key.get(id).copied()
    }

    /// Inserts a node into a parent at a specific position
    pub fn insert_node(&mut self, parent_key: NodeKey, child_key: NodeKey, anchor: Option<NodeKey>) {
        // Set parent relationship
        if let Some(child) = self.nodes.get_mut(child_key) {
            child.parent = Some(parent_key);
        }

        // Add to parent's children
        if let Some(parent) = self.nodes.get_mut(parent_key) {
            if let Some(anchor_key) = anchor {
                if let Some(pos) = parent.children.iter().position(|&k| k == anchor_key) {
                    parent.children.insert(pos, child_key);
                } else {
                    parent.children.push(child_key);
                }
            } else {
                parent.children.push(child_key);
            }
        }
    }

    /// Removes a node from its parent
    pub fn remove_node(&mut self, parent_key: NodeKey, child_key: NodeKey) {
        // Remove parent relationship
        if let Some(child) = self.nodes.get_mut(child_key) {
            child.parent = None;
        }

        // Remove from parent's children
        if let Some(parent) = self.nodes.get_mut(parent_key) {
            parent.children.retain(|&k| k != child_key);
        }
    }

    /// Deletes a node and all its descendants
    pub fn delete_node(&mut self, key: NodeKey) {
        // Collect all data we need before any mutable operations
        let (id, children, parent_key) = {
            if let Some(node) = self.nodes.get(key) {
                (node.id.clone(), node.children.clone(), node.parent)
            } else {
                return;
            }
        };

        // Recursively delete children
        for child_key in children {
            self.delete_node(child_key);
        }

        // Remove from parent if exists
        if let Some(parent_key) = parent_key {
            if let Some(parent) = self.nodes.get_mut(parent_key) {
                parent.children.retain(|&k| k != key);
            }
        }

        // Remove from maps
        self.id_to_key.remove(&id);
        self.nodes.remove(key);
    }

    /// Updates text content of a text node
    pub fn replace_text(&mut self, key: NodeKey, content: String) {
        if let Some(node) = self.nodes.get_mut(key) {
            if let NodeType::Text { content: old_content } = &mut node.node_type {
                *old_content = content;
            }
        }
    }

    /// Sets a property on a node
    pub fn set_property(&mut self, key: NodeKey, name: String, value: String) {
        if let Some(node) = self.nodes.get_mut(key) {
            node.properties.insert(name, value);
        }
    }

    /// Gets the parent of a node
    pub fn get_parent(&self, key: NodeKey) -> Option<NodeKey> {
        self.nodes.get(key).and_then(|n| n.parent)
    }

    /// Gets the first child of a node
    pub fn get_first_child(&self, key: NodeKey) -> Option<NodeKey> {
        self.nodes.get(key).and_then(|n| n.children.first().copied())
    }

    /// Gets the next sibling of a node
    pub fn get_next_sibling(&self, key: NodeKey) -> Option<NodeKey> {
        let parent_key = self.get_parent(key)?;
        let parent = self.nodes.get(parent_key)?;

        let pos = parent.children.iter().position(|&k| k == key)?;
        parent.children.get(pos + 1).copied()
    }

    /// Gets all children IDs of a node
    pub fn get_children_ids(&self, key: NodeKey) -> Vec<String> {
        if let Some(node) = self.nodes.get(key) {
            node.children
                .iter()
                .filter_map(|&child_key| {
                    self.nodes.get(child_key).map(|n| n.id.clone())
                })
                .collect()
        } else {
            Vec::new()
        }
    }

    /// Sets the root node of the tree
    pub fn set_root(&mut self, key: NodeKey) {
        self.root = Some(key);
    }

    /// Gets the root node key
    pub fn get_root(&self) -> Option<NodeKey> {
        self.root
    }

    /// Gets the root node ID
    pub fn get_root_id(&self) -> Option<String> {
        self.root
            .and_then(|key| self.nodes.get(key))
            .map(|node| node.id.clone())
    }
}

impl Default for UITree {
    fn default() -> Self {
        Self::new()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_create_element() {
        let mut tree = UITree::new();
        let key = tree.create_element("node1".to_string(), "div".to_string());

        let node = tree.get_node(key).unwrap();
        assert_eq!(node.id, "node1");
        assert_eq!(node.get_tag(), Some("div"));
        assert!(!node.is_text_node());
    }

    #[test]
    fn test_create_text_node() {
        let mut tree = UITree::new();
        let key = tree.create_text_node("text1".to_string(), "Hello".to_string());

        let node = tree.get_node(key).unwrap();
        assert_eq!(node.id, "text1");
        assert!(node.is_text_node());
    }

    #[test]
    fn test_get_key_by_id() {
        let mut tree = UITree::new();
        let key = tree.create_element("test_id".to_string(), "span".to_string());

        let found_key = tree.get_key_by_id("test_id");
        assert_eq!(found_key, Some(key));

        let not_found = tree.get_key_by_id("nonexistent");
        assert_eq!(not_found, None);
    }

    #[test]
    fn test_insert_node() {
        let mut tree = UITree::new();
        let parent_key = tree.create_element("parent".to_string(), "div".to_string());
        let child_key = tree.create_element("child".to_string(), "span".to_string());

        tree.insert_node(parent_key, child_key, None);

        let parent = tree.get_node(parent_key).unwrap();
        assert_eq!(parent.children.len(), 1);
        assert_eq!(parent.children[0], child_key);

        let child = tree.get_node(child_key).unwrap();
        assert_eq!(child.parent, Some(parent_key));
    }

    #[test]
    fn test_insert_node_with_anchor() {
        let mut tree = UITree::new();
        let parent_key = tree.create_element("parent".to_string(), "div".to_string());
        let child1_key = tree.create_element("child1".to_string(), "span".to_string());
        let child2_key = tree.create_element("child2".to_string(), "p".to_string());
        let child3_key = tree.create_element("child3".to_string(), "a".to_string());

        tree.insert_node(parent_key, child1_key, None);
        tree.insert_node(parent_key, child3_key, None);
        tree.insert_node(parent_key, child2_key, Some(child3_key)); // Insert before child3

        let parent = tree.get_node(parent_key).unwrap();
        assert_eq!(parent.children.len(), 3);
        assert_eq!(parent.children[0], child1_key);
        assert_eq!(parent.children[1], child2_key);
        assert_eq!(parent.children[2], child3_key);
    }

    #[test]
    fn test_remove_node() {
        let mut tree = UITree::new();
        let parent_key = tree.create_element("parent".to_string(), "div".to_string());
        let child_key = tree.create_element("child".to_string(), "span".to_string());

        tree.insert_node(parent_key, child_key, None);
        tree.remove_node(parent_key, child_key);

        let parent = tree.get_node(parent_key).unwrap();
        assert_eq!(parent.children.len(), 0);

        let child = tree.get_node(child_key).unwrap();
        assert_eq!(child.parent, None);
    }

    #[test]
    fn test_delete_node() {
        let mut tree = UITree::new();
        let parent_key = tree.create_element("parent".to_string(), "div".to_string());
        let child_key = tree.create_element("child".to_string(), "span".to_string());

        tree.insert_node(parent_key, child_key, None);
        tree.delete_node(child_key);

        assert!(tree.get_node(child_key).is_none());
        assert!(tree.get_key_by_id("child").is_none());

        let parent = tree.get_node(parent_key).unwrap();
        assert_eq!(parent.children.len(), 0);
    }

    #[test]
    fn test_delete_node_recursive() {
        let mut tree = UITree::new();
        let parent_key = tree.create_element("parent".to_string(), "div".to_string());
        let child_key = tree.create_element("child".to_string(), "span".to_string());
        let grandchild_key = tree.create_element("grandchild".to_string(), "p".to_string());

        tree.insert_node(parent_key, child_key, None);
        tree.insert_node(child_key, grandchild_key, None);

        tree.delete_node(child_key);

        assert!(tree.get_node(child_key).is_none());
        assert!(tree.get_node(grandchild_key).is_none());
        assert!(tree.get_key_by_id("child").is_none());
        assert!(tree.get_key_by_id("grandchild").is_none());
    }

    #[test]
    fn test_replace_text() {
        let mut tree = UITree::new();
        let key = tree.create_text_node("text1".to_string(), "Hello".to_string());

        tree.replace_text(key, "World".to_string());

        let node = tree.get_node(key).unwrap();
        match &node.node_type {
            NodeType::Text { content } => assert_eq!(content, "World"),
            _ => panic!("Expected text node"),
        }
    }

    #[test]
    fn test_set_property() {
        let mut tree = UITree::new();
        let key = tree.create_element("node1".to_string(), "div".to_string());

        tree.set_property(key, "className".to_string(), "container".to_string());
        tree.set_property(key, "id".to_string(), "main".to_string());

        let node = tree.get_node(key).unwrap();
        assert_eq!(node.properties.get("className"), Some(&"container".to_string()));
        assert_eq!(node.properties.get("id"), Some(&"main".to_string()));
    }

    #[test]
    fn test_get_parent() {
        let mut tree = UITree::new();
        let parent_key = tree.create_element("parent".to_string(), "div".to_string());
        let child_key = tree.create_element("child".to_string(), "span".to_string());

        tree.insert_node(parent_key, child_key, None);

        assert_eq!(tree.get_parent(child_key), Some(parent_key));
        assert_eq!(tree.get_parent(parent_key), None);
    }

    #[test]
    fn test_get_first_child() {
        let mut tree = UITree::new();
        let parent_key = tree.create_element("parent".to_string(), "div".to_string());
        let child1_key = tree.create_element("child1".to_string(), "span".to_string());
        let child2_key = tree.create_element("child2".to_string(), "p".to_string());

        tree.insert_node(parent_key, child1_key, None);
        tree.insert_node(parent_key, child2_key, None);

        assert_eq!(tree.get_first_child(parent_key), Some(child1_key));
        assert_eq!(tree.get_first_child(child1_key), None);
    }

    #[test]
    fn test_get_next_sibling() {
        let mut tree = UITree::new();
        let parent_key = tree.create_element("parent".to_string(), "div".to_string());
        let child1_key = tree.create_element("child1".to_string(), "span".to_string());
        let child2_key = tree.create_element("child2".to_string(), "p".to_string());
        let child3_key = tree.create_element("child3".to_string(), "a".to_string());

        tree.insert_node(parent_key, child1_key, None);
        tree.insert_node(parent_key, child2_key, None);
        tree.insert_node(parent_key, child3_key, None);

        assert_eq!(tree.get_next_sibling(child1_key), Some(child2_key));
        assert_eq!(tree.get_next_sibling(child2_key), Some(child3_key));
        assert_eq!(tree.get_next_sibling(child3_key), None);
    }

    #[test]
    fn test_get_children_ids() {
        let mut tree = UITree::new();
        let parent_key = tree.create_element("parent".to_string(), "div".to_string());
        let child1_key = tree.create_element("child1".to_string(), "span".to_string());
        let child2_key = tree.create_element("child2".to_string(), "p".to_string());

        tree.insert_node(parent_key, child1_key, None);
        tree.insert_node(parent_key, child2_key, None);

        let children_ids = tree.get_children_ids(parent_key);
        assert_eq!(children_ids.len(), 2);
        assert_eq!(children_ids[0], "child1");
        assert_eq!(children_ids[1], "child2");
    }
}
