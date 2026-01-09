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

    pub fn get_text_content(&self) -> Option<String> {
        match &self.node_type {
            NodeType::Text { content } => Some(content.clone()),
            NodeType::Element { .. } => None,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_new_element() {
        let node = Node::new_element("test_id".to_string(), "div".to_string());
        assert_eq!(node.id, "test_id");
        assert_eq!(node.get_tag(), Some("div"));
        assert!(!node.is_text_node());
        assert!(node.parent.is_none());
        assert!(node.children.is_empty());
        assert!(node.properties.is_empty());
    }

    #[test]
    fn test_new_text() {
        let node = Node::new_text("text_id".to_string(), "Hello World".to_string());
        assert_eq!(node.id, "text_id");
        assert!(node.is_text_node());
        assert_eq!(node.get_tag(), None);
        assert!(node.parent.is_none());
        assert!(node.children.is_empty());
    }

    #[test]
    fn test_is_text_node() {
        let element = Node::new_element("el".to_string(), "span".to_string());
        let text = Node::new_text("txt".to_string(), "content".to_string());

        assert!(!element.is_text_node());
        assert!(text.is_text_node());
    }

    #[test]
    fn test_get_tag() {
        let element = Node::new_element("el".to_string(), "button".to_string());
        let text = Node::new_text("txt".to_string(), "text".to_string());

        assert_eq!(element.get_tag(), Some("button"));
        assert_eq!(text.get_tag(), None);
    }
}
