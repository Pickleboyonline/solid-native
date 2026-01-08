use sncore::{HostDelegate, SolidRenderer};
use std::sync::{Arc, Mutex};

// Mock delegate for testing
struct TestDelegate {
    created_nodes: Arc<Mutex<Vec<(String, String)>>>,
    removed_nodes: Arc<Mutex<Vec<String>>>,
    children_changes: Arc<Mutex<Vec<(String, Vec<String>)>>>,
    update_counts: Arc<Mutex<Vec<String>>>,
}

impl TestDelegate {
    fn new() -> Self {
        Self {
            created_nodes: Arc::new(Mutex::new(Vec::new())),
            removed_nodes: Arc::new(Mutex::new(Vec::new())),
            children_changes: Arc::new(Mutex::new(Vec::new())),
            update_counts: Arc::new(Mutex::new(Vec::new())),
        }
    }

    fn get_created_nodes(&self) -> Vec<(String, String)> {
        self.created_nodes.lock().unwrap().clone()
    }

    fn get_removed_nodes(&self) -> Vec<String> {
        self.removed_nodes.lock().unwrap().clone()
    }

    fn get_children_changes(&self) -> Vec<(String, Vec<String>)> {
        self.children_changes.lock().unwrap().clone()
    }
}

impl HostDelegate for TestDelegate {
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
fn test_js_can_call_renderer_api() {
    let delegate = Arc::new(TestDelegate::new());
    let renderer = Arc::new(SolidRenderer::new(delegate.clone()).unwrap());

    // Simulate what would happen when JS calls the renderer API
    let parent_id = renderer.create_element("div".to_string());
    let child_id = renderer.create_text_node("Hello World".to_string());

    renderer.insert_node(parent_id.clone(), child_id.clone(), None);

    // Verify the tree was manipulated correctly
    let tree = renderer.get_tree();
    let tree_lock = tree.lock().unwrap();
    let parent_key = tree_lock.get_key_by_id(&parent_id).unwrap();
    let parent = tree_lock.get_node(parent_key).unwrap();

    assert_eq!(parent.children.len(), 1);
    assert_eq!(parent.get_tag(), Some("div"));

    // Verify delegate callbacks were triggered
    let created_nodes = delegate.get_created_nodes();
    assert_eq!(created_nodes.len(), 2);
    assert_eq!(created_nodes[0], (parent_id.clone(), "div".to_string()));
    assert_eq!(created_nodes[1], (child_id.clone(), "text".to_string()));
}

#[test]
fn test_js_solidjs_renderer_workflow() {
    // This test simulates a typical SolidJS rendering workflow
    let delegate = Arc::new(TestDelegate::new());
    let renderer = Arc::new(SolidRenderer::new(delegate.clone()).unwrap());

    // Step 1: Create root container
    let root_id = renderer.create_element("div".to_string());
    renderer.set_property(root_id.clone(), "id".to_string(), "root".to_string());

    // Step 2: Create a component structure
    // <div id="root">
    //   <h1>Title</h1>
    //   <p>Paragraph text</p>
    // </div>
    let h1_id = renderer.create_element("h1".to_string());
    let h1_text_id = renderer.create_text_node("Title".to_string());
    renderer.insert_node(h1_id.clone(), h1_text_id.clone(), None);

    let p_id = renderer.create_element("p".to_string());
    let p_text_id = renderer.create_text_node("Paragraph text".to_string());
    renderer.insert_node(p_id.clone(), p_text_id.clone(), None);

    renderer.insert_node(root_id.clone(), h1_id.clone(), None);
    renderer.insert_node(root_id.clone(), p_id.clone(), None);

    // Verify structure
    let tree = renderer.get_tree();
    let tree_lock = tree.lock().unwrap();
    let root_key = tree_lock.get_key_by_id(&root_id).unwrap();
    let root = tree_lock.get_node(root_key).unwrap();

    assert_eq!(root.children.len(), 2);
    assert_eq!(root.properties.get("id"), Some(&"root".to_string()));

    let h1_key = tree_lock.get_key_by_id(&h1_id).unwrap();
    let h1 = tree_lock.get_node(h1_key).unwrap();
    assert_eq!(h1.children.len(), 1);

    // Step 3: Update text (reactive update)
    drop(tree_lock);
    renderer.replace_text(p_text_id.clone(), "Updated paragraph text".to_string());

    let tree_lock = tree.lock().unwrap();
    let p_text_key = tree_lock.get_key_by_id(&p_text_id).unwrap();
    let p_text = tree_lock.get_node(p_text_key).unwrap();
    match &p_text.node_type {
        sncore::NodeType::Text { content } => {
            assert_eq!(content, "Updated paragraph text");
        }
        _ => panic!("Expected text node"),
    }

    // Verify all nodes were created
    let created_nodes = delegate.get_created_nodes();
    assert_eq!(created_nodes.len(), 5); // root, h1, h1_text, p, p_text
}

#[test]
fn test_js_dynamic_list_rendering() {
    // This test simulates rendering a dynamic list like SolidJS's <For> component
    let delegate = Arc::new(TestDelegate::new());
    let renderer = Arc::new(SolidRenderer::new(delegate.clone()).unwrap());

    // Create list container
    let ul_id = renderer.create_element("ul".to_string());

    // Add initial items
    let mut item_ids = Vec::new();
    for i in 0..3 {
        let li_id = renderer.create_element("li".to_string());
        let text_id = renderer.create_text_node(format!("Item {}", i));
        renderer.insert_node(li_id.clone(), text_id.clone(), None);
        renderer.insert_node(ul_id.clone(), li_id.clone(), None);
        item_ids.push(li_id);
    }

    // Verify initial state
    let tree = renderer.get_tree();
    let tree_lock = tree.lock().unwrap();
    let ul_key = tree_lock.get_key_by_id(&ul_id).unwrap();
    let ul = tree_lock.get_node(ul_key).unwrap();
    assert_eq!(ul.children.len(), 3);

    // Simulate removing an item (like filtering a list)
    drop(tree_lock);
    let removed_item = item_ids[1].clone();
    renderer.remove_node(ul_id.clone(), removed_item.clone());

    let tree_lock = tree.lock().unwrap();
    let ul = tree_lock.get_node(ul_key).unwrap();
    assert_eq!(ul.children.len(), 2);

    // Verify removal callback was triggered
    drop(tree_lock);
    let removed_nodes = delegate.get_removed_nodes();
    assert!(removed_nodes.contains(&removed_item));
}

#[test]
fn test_js_insert_with_anchor() {
    // Test inserting nodes at specific positions (important for SolidJS reconciliation)
    let delegate = Arc::new(TestDelegate::new());
    let renderer = Arc::new(SolidRenderer::new(delegate.clone()).unwrap());

    let container_id = renderer.create_element("div".to_string());

    // Add first and last items
    let first_id = renderer.create_text_node("First".to_string());
    let last_id = renderer.create_text_node("Last".to_string());

    renderer.insert_node(container_id.clone(), first_id.clone(), None);
    renderer.insert_node(container_id.clone(), last_id.clone(), None);

    // Insert middle item before "Last"
    let middle_id = renderer.create_text_node("Middle".to_string());
    renderer.insert_node(container_id.clone(), middle_id.clone(), Some(last_id.clone()));

    // Verify order
    let tree = renderer.get_tree();
    let tree_lock = tree.lock().unwrap();
    let container_key = tree_lock.get_key_by_id(&container_id).unwrap();
    let container = tree_lock.get_node(container_key).unwrap();

    assert_eq!(container.children.len(), 3);
    let children_ids = tree_lock.get_children_ids(container_key);
    assert_eq!(children_ids[0], first_id);
    assert_eq!(children_ids[1], middle_id);
    assert_eq!(children_ids[2], last_id);
}

#[test]
fn test_js_tree_traversal() {
    // Test the tree traversal API that SolidJS uses
    let delegate = Arc::new(TestDelegate::new());
    let renderer = Arc::new(SolidRenderer::new(delegate.clone()).unwrap());

    // Build tree:
    // parent
    //   ├─ child1
    //   └─ child2
    let parent_id = renderer.create_element("div".to_string());
    let child1_id = renderer.create_element("span".to_string());
    let child2_id = renderer.create_element("p".to_string());

    renderer.insert_node(parent_id.clone(), child1_id.clone(), None);
    renderer.insert_node(parent_id.clone(), child2_id.clone(), None);

    // Test getParentNode
    assert_eq!(
        renderer.get_parent_node(child1_id.clone()),
        Some(parent_id.clone())
    );
    assert_eq!(
        renderer.get_parent_node(child2_id.clone()),
        Some(parent_id.clone())
    );

    // Test getFirstChild
    assert_eq!(
        renderer.get_first_child(parent_id.clone()),
        Some(child1_id.clone())
    );

    // Test getNextSibling
    assert_eq!(
        renderer.get_next_sibling(child1_id.clone()),
        Some(child2_id.clone())
    );
    assert_eq!(renderer.get_next_sibling(child2_id.clone()), None);

    // Test isTextNode
    assert!(!renderer.is_text_node(parent_id));
    assert!(!renderer.is_text_node(child1_id));
}

#[test]
fn test_js_property_updates() {
    // Test setting properties like SolidJS does for reactive attributes
    let delegate = Arc::new(TestDelegate::new());
    let renderer = Arc::new(SolidRenderer::new(delegate.clone()).unwrap());

    let div_id = renderer.create_element("div".to_string());

    // Set multiple properties
    renderer.set_property(div_id.clone(), "className".to_string(), "container".to_string());
    renderer.set_property(div_id.clone(), "id".to_string(), "main".to_string());
    renderer.set_property(div_id.clone(), "data-testid".to_string(), "test".to_string());

    // Update a property
    renderer.set_property(div_id.clone(), "className".to_string(), "updated-container".to_string());

    let tree = renderer.get_tree();
    let tree_lock = tree.lock().unwrap();
    let div_key = tree_lock.get_key_by_id(&div_id).unwrap();
    let div = tree_lock.get_node(div_key).unwrap();

    assert_eq!(div.properties.len(), 3);
    assert_eq!(
        div.properties.get("className"),
        Some(&"updated-container".to_string())
    );
    assert_eq!(div.properties.get("id"), Some(&"main".to_string()));
    assert_eq!(
        div.properties.get("data-testid"),
        Some(&"test".to_string())
    );
}
