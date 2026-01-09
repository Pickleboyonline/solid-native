use crate::delegate::HostDelegate;
use crate::error::SNCoreError;
use crate::js_bindings::bind_renderer_to_context;
use crate::renderer::SolidRenderer;
use rquickjs::{Context, Runtime};
use std::sync::{Arc, Mutex};

/// Main entry point for SolidNative Core
/// This struct is exposed to the host platform via Uniffi
/// and manages both the JavaScript runtime and the renderer
#[derive(uniffi::Object)]
pub struct SolidNativeCore {
    _runtime: Arc<Runtime>, // Keep runtime alive, underscore to silence unused warning
    context: Arc<Mutex<Context>>,
    renderer: Arc<SolidRenderer>,
}

#[uniffi::export]
impl SolidNativeCore {
    /// Creates a new SNCore instance with the given delegate
    /// The delegate receives callbacks when the UI tree changes
    #[uniffi::constructor]
    pub fn new(delegate: Arc<dyn HostDelegate>) -> Result<Arc<Self>, SNCoreError> {
        // Create runtime and wrap in Arc for shared ownership
        let runtime =
            Arc::new(Runtime::new().map_err(|e| SNCoreError::runtime(format!("Failed to create runtime: {:?}", e)))?);

        let renderer = Arc::new(
            SolidRenderer::new(delegate)
                .map_err(|e| SNCoreError::renderer(format!("Failed to create renderer: {}", e)))?,
        );

        // Create persistent context
        let context =
            Context::full(&*runtime).map_err(|e| SNCoreError::context(format!("Failed to create context: {:?}", e)))?;

        // Bind renderer to context once
        bind_renderer_to_context(&context, renderer.clone())
            .map_err(|e| SNCoreError::context(format!("Failed to bind renderer: {:?}", e)))?;

        Ok(Arc::new(Self {
            _runtime: runtime,
            context: Arc::new(Mutex::new(context)),
            renderer,
        }))
    }

    // ==================== Host API Methods ====================
    // These methods allow the host platform to directly manipulate the tree

    /// Creates an element node
    pub fn create_element(&self, tag: String) -> String {
        self.renderer.create_element(tag)
    }

    /// Creates a text node
    pub fn create_text_node(&self, value: String) -> String {
        self.renderer.create_text_node(value)
    }

    /// Inserts a node into a parent before an anchor node
    pub fn insert_node(&self, parent_id: String, node_id: String, anchor_id: Option<String>) {
        self.renderer.insert_node(parent_id, node_id, anchor_id)
    }

    /// Removes a node from its parent
    pub fn remove_node(&self, parent_id: String, node_id: String) {
        self.renderer.remove_node(parent_id, node_id)
    }

    /// Sets a property on a node
    pub fn set_property(&self, node_id: String, name: String, value: String) {
        self.renderer.set_property(node_id, name, value)
    }

    /// Replaces text content of a text node
    pub fn replace_text(&self, node_id: String, value: String) {
        self.renderer.replace_text(node_id, value)
    }

    /// Checks if a node is a text node
    pub fn is_text_node(&self, node_id: String) -> bool {
        self.renderer.is_text_node(node_id)
    }

    /// Gets the parent node ID
    pub fn get_parent_node(&self, node_id: String) -> Option<String> {
        self.renderer.get_parent_node(node_id)
    }

    /// Gets the first child node ID
    pub fn get_first_child(&self, node_id: String) -> Option<String> {
        self.renderer.get_first_child(node_id)
    }

    /// Gets the next sibling node ID
    pub fn get_next_sibling(&self, node_id: String) -> Option<String> {
        self.renderer.get_next_sibling(node_id)
    }

    // ==================== JavaScript Execution ====================

    /// Evaluates JavaScript code
    /// The renderer is automatically bound to globalThis.solidNative
    /// Uses a persistent context, so JavaScript state persists across calls
    pub fn eval_js(&self, code: String) -> Result<String, SNCoreError> {
        let context = self.context.lock().unwrap();

        // Execute the JavaScript code using the persistent context
        context.with(|ctx| {
            ctx.eval::<String, _>(code)
                .map_err(|e| SNCoreError::js_eval(format!("JS eval error: {:?}", e)))
        })
    }

    /// Evaluates JavaScript module code
    /// The renderer is automatically bound to globalThis.solidNative
    /// Uses a persistent context, so JavaScript state persists across calls
    pub fn eval_module(&self, code: String, module_name: String) -> Result<String, SNCoreError> {
        let context = self.context.lock().unwrap();

        // Execute as module using the persistent context
        context.with(|ctx| {
            // For module evaluation, we need to compile and execute differently
            // This is a simplified version - in production you'd want better module handling
            let wrapped_code = format!(
                r#"
                (function() {{
                    {}
                }})();
                "#,
                code
            );

            ctx.eval::<String, _>(wrapped_code)
                .map_err(|e| SNCoreError::js_eval(format!("Module eval error for '{}': {:?}", module_name, e)))
        })
    }
}

// Non-uniffi methods (for internal/test use)
impl SolidNativeCore {
    /// Gets access to the renderer (for advanced use cases)
    /// Not exposed via uniffi - for internal use only
    pub fn get_renderer(&self) -> Arc<SolidRenderer> {
        self.renderer.clone()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::delegate::HostDelegate;
    use std::sync::Mutex;

    struct MockDelegate {
        created_nodes: Arc<Mutex<Vec<(String, String)>>>,
    }

    impl MockDelegate {
        fn new() -> Self {
            Self {
                created_nodes: Arc::new(Mutex::new(Vec::new())),
            }
        }

        fn get_created_count(&self) -> usize {
            self.created_nodes.lock().unwrap().len()
        }
    }

    impl HostDelegate for MockDelegate {
        fn on_node_created(&self, node_id: String, node_type: String) {
            self.created_nodes
                .lock()
                .unwrap()
                .push((node_id, node_type));
        }

        fn on_node_removed(&self, _node_id: String) {}
        fn on_children_change(&self, _node_id: String, _node_ids: Vec<String>) {}
        fn on_update_revision_count(&self, _node_id: String) {}
        fn is_text_element_by_node_id(&self, _node_id: String) -> bool {
            false
        }
        fn is_text_element_by_node_type(&self, node_type: String) -> bool {
            node_type == "text"
        }
    }

    #[test]
    fn test_sncore_creation() {
        let delegate = Arc::new(MockDelegate::new());
        let core = SolidNativeCore::new(delegate).unwrap();
        assert!(core.get_renderer().is_text_node("nonexistent".to_string()) == false);
    }

    #[test]
    fn test_host_can_create_elements() {
        let delegate = Arc::new(MockDelegate::new());
        let core = SolidNativeCore::new(delegate.clone()).unwrap();

        let node_id = core.create_element("div".to_string());
        // UUID format check
        assert_eq!(node_id.len(), 36);
        assert_eq!(delegate.get_created_count(), 1);
    }

    #[test]
    fn test_host_can_manipulate_tree() {
        let delegate = Arc::new(MockDelegate::new());
        let core = SolidNativeCore::new(delegate.clone()).unwrap();

        let parent_id = core.create_element("div".to_string());
        let child_id = core.create_element("span".to_string());

        core.insert_node(parent_id.clone(), child_id.clone(), None);

        let found_parent = core.get_parent_node(child_id.clone());
        assert_eq!(found_parent, Some(parent_id.clone()));

        let first_child = core.get_first_child(parent_id);
        assert_eq!(first_child, Some(child_id));
    }

    #[test]
    fn test_eval_js_basic() {
        let delegate = Arc::new(MockDelegate::new());
        let core = SolidNativeCore::new(delegate).unwrap();

        let result = core.eval_js("(2 + 3).toString()".to_string());
        assert_eq!(result, Ok("5".to_string()));
    }

    #[test]
    fn test_eval_js_with_renderer_binding() {
        let delegate = Arc::new(MockDelegate::new());
        let core = SolidNativeCore::new(delegate.clone()).unwrap();

        let result = core
            .eval_js("globalThis.solidNative.createElement('div')".to_string())
            .unwrap();

        // UUID format check
        assert_eq!(result.len(), 36);
        assert_eq!(delegate.get_created_count(), 1);
    }

    #[test]
    fn test_js_and_host_share_renderer() {
        let delegate = Arc::new(MockDelegate::new());
        let core = SolidNativeCore::new(delegate.clone()).unwrap();

        // Host creates parent
        let parent_id = core.create_element("div".to_string());

        // JS creates child
        let child_id: String = core
            .eval_js("globalThis.solidNative.createElement('span')".to_string())
            .unwrap();

        // Host inserts child (avoiding JS insertNode for now)
        core.insert_node(parent_id.clone(), child_id.clone(), None);

        // Host can verify the structure
        let first_child = core.get_first_child(parent_id.clone());
        assert_eq!(first_child, Some(child_id.clone()));

        let parent = core.get_parent_node(child_id);
        assert_eq!(parent, Some(parent_id));
    }

    #[test]
    fn test_multiple_js_executions() {
        let delegate = Arc::new(MockDelegate::new());
        let core = SolidNativeCore::new(delegate.clone()).unwrap();

        // First execution
        core.eval_js("globalThis.solidNative.createElement('div')".to_string())
            .unwrap();

        // Second execution should also work (same context, same renderer)
        core.eval_js("globalThis.solidNative.createElement('span')".to_string())
            .unwrap();

        assert_eq!(delegate.get_created_count(), 2);
    }

    #[test]
    fn test_persistent_context_state() {
        let delegate = Arc::new(MockDelegate::new());
        let core = SolidNativeCore::new(delegate.clone()).unwrap();

        // Set a variable in JS and return it as string
        let result = core
            .eval_js("globalThis.testCounter = 42; testCounter.toString()".to_string())
            .unwrap();

        assert_eq!(result, "42");

        // Read it back in a different eval call
        let result = core
            .eval_js("globalThis.testCounter.toString()".to_string())
            .unwrap();

        assert_eq!(result, "42");

        // Increment it and return new value
        let result = core
            .eval_js("globalThis.testCounter += 8; testCounter.toString()".to_string())
            .unwrap();

        assert_eq!(result, "50");

        // Verify it persisted
        let result = core
            .eval_js("globalThis.testCounter.toString()".to_string())
            .unwrap();

        assert_eq!(result, "50");
    }

    #[test]
    fn test_solidnative_binding_persists() {
        let delegate = Arc::new(MockDelegate::new());
        let core = SolidNativeCore::new(delegate.clone()).unwrap();

        // Verify solidNative is accessible in first call
        let result = core
            .eval_js("(typeof globalThis.solidNative !== 'undefined').toString()".to_string())
            .unwrap();

        assert_eq!(result, "true");

        // Verify it's still accessible in subsequent calls
        let result = core
            .eval_js(
                "(typeof globalThis.solidNative.createElement === 'function').toString()"
                    .to_string(),
            )
            .unwrap();

        assert_eq!(result, "true");

        // Actually use it
        let node_id = core
            .eval_js("globalThis.solidNative.createElement('div')".to_string())
            .unwrap();

        assert_eq!(node_id.len(), 36); // UUID length
        assert_eq!(delegate.get_created_count(), 1);
    }

    #[test]
    fn test_complex_js_tree_building() {
        let delegate = Arc::new(MockDelegate::new());
        let core = SolidNativeCore::new(delegate.clone()).unwrap();

        // Build: <div><h1>Title</h1><p>Text</p></div>
        // Create nodes via JS
        let root_id = core
            .eval_js("globalThis.solidNative.createElement('div')".to_string())
            .unwrap();
        let h1_id = core
            .eval_js("globalThis.solidNative.createElement('h1')".to_string())
            .unwrap();
        let h1_text_id = core
            .eval_js("globalThis.solidNative.createTextNode('Title')".to_string())
            .unwrap();
        let p_id = core
            .eval_js("globalThis.solidNative.createElement('p')".to_string())
            .unwrap();
        let p_text_id = core
            .eval_js("globalThis.solidNative.createTextNode('Text')".to_string())
            .unwrap();

        // Use host API to build structure
        core.insert_node(h1_id.clone(), h1_text_id, None);
        core.insert_node(p_id.clone(), p_text_id, None);
        core.insert_node(root_id.clone(), h1_id.clone(), None);
        core.insert_node(root_id.clone(), p_id.clone(), None);

        // Verify structure from host side
        let first_child = core.get_first_child(root_id.clone()).unwrap();
        let second_child = core.get_next_sibling(first_child.clone()).unwrap();

        assert!(core.get_first_child(first_child).is_some()); // h1 has text child
        assert!(core.get_first_child(second_child).is_some()); // p has text child

        // Should have created 5 nodes: root, h1, h1Text, p, pText
        assert_eq!(delegate.get_created_count(), 5);
    }
}
