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
}

impl Default for UITree {
    fn default() -> Self {
        Self::new()
    }
}
