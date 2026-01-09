use crate::jsvalue::JSValue;
use crate::node::NodeKey;
use crate::tree::UITree;
use std::collections::HashMap;

/// Represents a piece of text with its associated styles
/// Similar to React Native's text descriptor pattern
#[derive(Debug, Clone, PartialEq, uniffi::Record)]
pub struct TextDescriptor {
    /// The actual text content
    pub text: String,
    /// Styles applied to this text segment
    pub styles: HashMap<String, JSValue>,
}

impl TextDescriptor {
    /// Creates a new TextDescriptor
    pub fn new(text: String, styles: HashMap<String, JSValue>) -> Self {
        Self { text, styles }
    }

    /// Creates a TextDescriptor with no styles
    pub fn plain(text: String) -> Self {
        Self {
            text,
            styles: HashMap::new(),
        }
    }
}

/// Generates text descriptors for a text node
/// This is called when text nodes change (insert, remove, set prop)
/// Returns the array of text descriptors and the top-level text node key
pub fn generate_text_descriptors(
    tree: &UITree,
    node_key: NodeKey,
) -> Option<(Vec<TextDescriptor>, NodeKey)> {
    let node = tree.get_node(node_key)?;

    // If this isn't a text node, return None
    if !node.is_text_node() {
        return None;
    }

    // Find the top-level text node (walk up until parent is not a text node)
    let mut top_level_key = node_key;
    while let Some(parent_key) = tree.get_parent(top_level_key) {
        if let Some(parent) = tree.get_node(parent_key) {
            if parent.is_text_node() {
                top_level_key = parent_key;
            } else {
                break;
            }
        } else {
            break;
        }
    }

    // Generate descriptors recursively from the top level
    let descriptors = generate_text_descriptors_recursively(tree, top_level_key, None);

    Some((descriptors, top_level_key))
}

/// Recursively generates text descriptors for a node and its children
/// Styles from parent nodes are passed down and can be overridden by children
fn generate_text_descriptors_recursively(
    tree: &UITree,
    node_key: NodeKey,
    parent_styles: Option<&HashMap<String, JSValue>>,
) -> Vec<TextDescriptor> {
    let node = match tree.get_node(node_key) {
        Some(n) => n,
        None => return Vec::new(),
    };

    // Merge parent styles with this node's styles
    // Child styles override parent styles
    let merged_styles = merge_styles(parent_styles, &node.properties);

    // If this is a leaf text node (no children and has text content),
    // return a single descriptor
    if node.children.is_empty() {
        if let Some(text_content) = node.get_text_content() {
            return vec![TextDescriptor::new(text_content, merged_styles)];
        } else {
            return Vec::new();
        }
    }

    // If this node has children, recursively process them
    let mut descriptors = Vec::new();
    for &child_key in &node.children {
        let child_descriptors =
            generate_text_descriptors_recursively(tree, child_key, Some(&merged_styles));
        descriptors.extend(child_descriptors);
    }

    descriptors
}

/// Merges parent styles with node styles
/// Node styles (properties) override parent styles
fn merge_styles(
    parent_styles: Option<&HashMap<String, JSValue>>,
    node_properties: &HashMap<String, String>,
) -> HashMap<String, JSValue> {
    let mut merged = HashMap::new();

    // First add parent styles
    if let Some(parent) = parent_styles {
        for (key, value) in parent {
            merged.insert(key.clone(), value.clone());
        }
    }

    // Then add/override with node properties
    // Note: properties are stored as String, so we wrap them in JSValue::String
    for (key, value) in node_properties {
        merged.insert(key.clone(), JSValue::string(value.clone()));
    }

    merged
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::node::NodeType;

    #[test]
    fn test_text_descriptor_creation() {
        let descriptor = TextDescriptor::plain("Hello".to_string());
        assert_eq!(descriptor.text, "Hello");
        assert!(descriptor.styles.is_empty());

        let mut styles = HashMap::new();
        styles.insert("color".to_string(), JSValue::string("red"));
        let descriptor = TextDescriptor::new("World".to_string(), styles);
        assert_eq!(descriptor.text, "World");
        assert_eq!(descriptor.styles.len(), 1);
    }

    #[test]
    fn test_merge_styles() {
        let mut parent = HashMap::new();
        parent.insert("fontSize".to_string(), JSValue::number(16.0));
        parent.insert("color".to_string(), JSValue::string("black"));

        let mut node = HashMap::new();
        node.insert("color".to_string(), "red".to_string());
        node.insert("fontWeight".to_string(), "bold".to_string());

        let merged = merge_styles(Some(&parent), &node);

        // Should have fontSize from parent
        assert!(matches!(
            merged.get("fontSize"),
            Some(JSValue::Number { value: 16.0 })
        ));

        // Should have color overridden by node
        assert_eq!(
            merged.get("color"),
            Some(&JSValue::string("red"))
        );

        // Should have fontWeight from node
        assert_eq!(
            merged.get("fontWeight"),
            Some(&JSValue::string("bold"))
        );
    }

    #[test]
    fn test_generate_text_descriptors_single_node() {
        let mut tree = UITree::new();
        let key = tree.create_text_node("node1".to_string(), "Hello World".to_string());

        let result = generate_text_descriptors(&tree, key);
        assert!(result.is_some());

        let (descriptors, top_key) = result.unwrap();
        assert_eq!(top_key, key);
        assert_eq!(descriptors.len(), 1);
        assert_eq!(descriptors[0].text, "Hello World");
    }

    #[test]
    fn test_generate_text_descriptors_nested() {
        let mut tree = UITree::new();

        // Create parent text node: "Hello "
        let parent_key = tree.create_text_node("parent".to_string(), "Hello ".to_string());

        // Create child text node: "World"
        let child_key = tree.create_text_node("child".to_string(), "World".to_string());

        // Set style on child
        tree.set_property(child_key, "fontWeight".to_string(), "bold".to_string());

        // Insert child into parent
        tree.insert_node(parent_key, child_key, None);

        // Generate descriptors from child (should walk up to parent)
        let result = generate_text_descriptors(&tree, child_key);
        assert!(result.is_some());

        let (descriptors, top_key) = result.unwrap();
        assert_eq!(top_key, parent_key);
        assert_eq!(descriptors.len(), 2);

        // First descriptor: "Hello " with no styles
        assert_eq!(descriptors[0].text, "Hello ");

        // Second descriptor: "World" with bold style
        assert_eq!(descriptors[1].text, "World");
        assert_eq!(
            descriptors[1].styles.get("fontWeight"),
            Some(&JSValue::string("bold"))
        );
    }
}
