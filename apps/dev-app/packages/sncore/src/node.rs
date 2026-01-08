use slotmap::DefaultKey;
use std::collections::HashMap;

/// Type alias for node keys in the slotmap
pub type NodeKey = DefaultKey;

/// Represents different types of nodes in the UI tree
#[derive(Debug, Clone)]
pub enum NodeType {
    Element { tag: String },
    Text { content: String },
}

/// Represents a node in the UI tree
#[derive(Debug, Clone)]
pub struct Node {
    pub id: String,
    pub node_type: NodeType,
    pub parent: Option<NodeKey>,
    pub children: Vec<NodeKey>,
    pub properties: HashMap<String, String>,
}

impl Node {
    pub fn new_element(id: String, tag: String) -> Self {
        Self {
            id,
            node_type: NodeType::Element { tag },
            parent: None,
            children: Vec::new(),
            properties: HashMap::new(),
        }
    }

    pub fn new_text(id: String, content: String) -> Self {
        Self {
            id,
            node_type: NodeType::Text { content },
            parent: None,
            children: Vec::new(),
            properties: HashMap::new(),
        }
    }

    pub fn is_text_node(&self) -> bool {
        matches!(self.node_type, NodeType::Text { .. })
    }

    pub fn get_tag(&self) -> Option<&str> {
        match &self.node_type {
            NodeType::Element { tag } => Some(tag.as_str()),
            NodeType::Text { .. } => None,
        }
    }
}
