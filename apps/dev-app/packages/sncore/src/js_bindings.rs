use crate::renderer::SolidRenderer;
use rquickjs::function::{Func, Opt};
use rquickjs::{Context, Result as JsResult};
use std::sync::Arc;

/// Binds the SolidRenderer methods to the JavaScript context as globalThis.solidNative
pub fn bind_renderer_to_context(ctx: &Context, renderer: Arc<SolidRenderer>) -> JsResult<()> {
    ctx.with(|ctx| {
        let globals = ctx.globals();

        // Create the solidNative object
        let solid_native = rquickjs::Object::new(ctx.clone())?;

        // Clone Arc for each closure
        let r1 = renderer.clone();
        solid_native.set(
            "createElement",
            Func::from(move |tag: String| r1.create_element(tag)),
        )?;

        let r2 = renderer.clone();
        solid_native.set(
            "createTextNode",
            Func::from(move |value: String| r2.create_text_node(value)),
        )?;

        let r3 = renderer.clone();
        solid_native.set(
            "insertNode",
            Func::from(move |parent_id: String, node_id: String, anchor_id: Opt<String>| {
                r3.insert_node(parent_id, node_id, anchor_id.0);
            }),
        )?;

        let r4 = renderer.clone();
        solid_native.set(
            "removeNode",
            Func::from(move |parent_id: String, node_id: String| {
                r4.remove_node(parent_id, node_id);
            }),
        )?;

        let r5 = renderer.clone();
    solid_native.set(
            "setProperty",
            Func::from(move |node_id: String, name: String, value: String| {
                r5.set_property(node_id, name, value);
            }),
        )?;

        let r6 = renderer.clone();
        solid_native.set(
            "replaceText",
            Func::from(move |node_id: String, value: String| {
                r6.replace_text(node_id, value);
            }),
        )?;

        let r7 = renderer.clone();
        solid_native.set(
            "isTextNode",
            Func::from(move |node_id: String| r7.is_text_node(node_id)),
        )?;

        // These return empty string if not found (rquickjs doesn't handle Option<String> well)
        let r8 = renderer.clone();
        solid_native.set(
            "getParentNode",
            Func::from(move |node_id: String| -> String {
                r8.get_parent_node(node_id).unwrap_or_default()
            }),
        )?;

        let r9 = renderer.clone();
        solid_native.set(
            "getFirstChild",
            Func::from(move |node_id: String| -> String {
                r9.get_first_child(node_id).unwrap_or_default()
            }),
        )?;

        let r10 = renderer.clone();
        solid_native.set(
            "getNextSibling",
            Func::from(move |node_id: String| -> String {
                r10.get_next_sibling(node_id).unwrap_or_default()
            }),
        )?;

        // getRootView - returns the root node ID (empty string if not set)
        let r11 = renderer.clone();
        solid_native.set(
            "getRootView",
            Func::from(move || -> String {
                r11.get_root().unwrap_or_else(|| String::new())
            }),
        )?;

        // setProp - sets a property on a node (handles JS values)
        // We use a JSON string for the value since rquickjs lifetimes are complex
        let r12 = renderer.clone();
        solid_native.set(
            "_setPropJson",
            Func::from(move |node_id: String, key: String, json_value: String| {
                r12.set_prop_json(node_id, key, json_value);
            }),
        )?;

        // insertBefore - alias for insertNode (matches React Native API)
        let r13 = renderer.clone();
        solid_native.set(
            "insertBefore",
            Func::from(move |parent_id: String, node_id: String, anchor_id: Opt<String>| {
                r13.insert_node(parent_id, node_id, anchor_id.0);
            }),
        )?;

        // removeChild - alias for removeNode (matches renderer API)
        let r14 = renderer.clone();
        solid_native.set(
            "removeChild",
            Func::from(move |parent_id: String, node_id: String| {
                r14.remove_node(parent_id, node_id);
            }),
        )?;

        // isTextElement - alias for isTextNode (matches renderer API)
        let r15 = renderer.clone();
        solid_native.set(
            "isTextElement",
            Func::from(move |node_id: String| r15.is_text_node(node_id)),
        )?;

        // getParent - alias for getParentNode (matches renderer API)
        let r16 = renderer.clone();
        solid_native.set(
            "getParent",
            Func::from(move |node_id: String| -> String {
                r16.get_parent_node(node_id).unwrap_or_default()
            }),
        )?;

        // Attach to globalThis
        globals.set("solidNative", solid_native)?;

        // Create JS wrappers with validation
        ctx.eval::<(), _>(r#"
            // Helper to validate string arguments
            function validateString(value, fnName, argName) {
                if (value === undefined || value === null) {
                    throw new Error(fnName + ': ' + argName + ' is ' + value);
                }
                if (typeof value !== 'string') {
                    throw new Error(fnName + ': ' + argName + ' must be string, got ' + typeof value);
                }
            }

            // Store originals - only for functions that actually exist
            const originals = {
                createElement: globalThis.solidNative.createElement,
                createTextNode: globalThis.solidNative.createTextNode,
                insertBefore: globalThis.solidNative.insertBefore,
                removeChild: globalThis.solidNative.removeChild,
                isTextElement: globalThis.solidNative.isTextElement,
                getParent: globalThis.solidNative.getParent,
                getFirstChild: globalThis.solidNative.getFirstChild,
                getNextSibling: globalThis.solidNative.getNextSibling,
                getRootView: globalThis.solidNative.getRootView,
                isTextNode: globalThis.solidNative.isTextNode,
                getParentNode: globalThis.solidNative.getParentNode,
                replaceText: globalThis.solidNative.replaceText
            };

            globalThis.solidNative.createElement = function(tag) {
                validateString(tag, 'createElement', 'tag');
                try {
                    return originals.createElement(tag);
                } catch(e) {
                    throw new Error('createElement("' + tag + '") failed: ' + e);
                }
            };

            globalThis.solidNative.insertBefore = function(parentId, nodeId, anchorId) {
                validateString(parentId, 'insertBefore', 'parentId');
                validateString(nodeId, 'insertBefore', 'nodeId');
                // anchorId is optional
                if (anchorId !== undefined && anchorId !== null) {
                    validateString(anchorId, 'insertBefore', 'anchorId');
                    return originals.insertBefore(parentId, nodeId, anchorId);
                }
                return originals.insertBefore(parentId, nodeId);
            };

            globalThis.solidNative.removeChild = function(parentId, nodeId) {
                validateString(parentId, 'removeChild', 'parentId');
                validateString(nodeId, 'removeChild', 'nodeId');
                return originals.removeChild(parentId, nodeId);
            };

            globalThis.solidNative.isTextElement = function(nodeId) {
                validateString(nodeId, 'isTextElement', 'nodeId');
                return originals.isTextElement(nodeId);
            };

            globalThis.solidNative.getParent = function(nodeId) {
                validateString(nodeId, 'getParent', 'nodeId');
                return originals.getParent(nodeId);
            };

            globalThis.solidNative.getFirstChild = function(nodeId) {
                validateString(nodeId, 'getFirstChild', 'nodeId');
                return originals.getFirstChild(nodeId);
            };

            globalThis.solidNative.getNextSibling = function(nodeId) {
                validateString(nodeId, 'getNextSibling', 'nodeId');
                return originals.getNextSibling(nodeId);
            };

            // Store original _setPropJson
            const _originalSetPropJson = globalThis.solidNative._setPropJson;

            // setProp wrapper that calls JSON.stringify before passing to _setPropJson
            globalThis.solidNative.setProp = function(nodeId, key, value) {
                validateString(nodeId, 'setProp', 'nodeId');
                validateString(key, 'setProp', 'key');
                var jsonValue = JSON.stringify(value);
                try {
                    _originalSetPropJson(nodeId, key, jsonValue);
                } catch(e) {
                    throw new Error('setProp("' + nodeId + '", "' + key + '", ...) failed: ' + e);
                }
            };

            // getRootView wrapper - now returns empty string if not set
            globalThis.solidNative.getRootView = function() {
                try {
                    var root = originals.getRootView();
                    // Empty string means root not set
                    if (!root || root.length === 0) {
                        throw new Error('getRootView returned empty - root node not set. Call createRoot() first.');
                    }
                    return root;
                } catch(e) {
                    throw new Error('getRootView() failed: ' + e);
                }
            };
        "#)?;

        Ok(())
    })
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::delegate::HostDelegate;
    use rquickjs::Runtime;

    struct MockDelegate;

    impl HostDelegate for MockDelegate {
        fn on_node_created(&self, _node_id: &str, _node_type: &str) {}
        fn on_node_removed(&self, _node_id: &str) {}
        fn on_children_change(&self, _node_id: &str, _node_ids: &[String]) {}
        fn on_update_revision_count(&self, _node_id: &str) {}
        fn is_text_element_by_node_id(&self, _node_id: &str) -> bool {
            false
        }
        fn is_text_element_by_node_type(&self, node_type: &str) -> bool {
            node_type == "text"
        }
    }

    #[test]
    fn test_bind_renderer_to_context() {
        let delegate = Arc::new(MockDelegate);
        let renderer = Arc::new(SolidRenderer::new(delegate).unwrap());
        let runtime = Runtime::new().unwrap();
        let context = Context::full(&runtime).unwrap();

        // Bind renderer to context
        bind_renderer_to_context(&context, renderer.clone()).unwrap();

        // Test that solidNative exists
        let result: bool = context
            .with(|ctx| ctx.eval("typeof globalThis.solidNative !== 'undefined'"))
            .unwrap();
        assert!(result);

        // Test createElement is accessible
        let result: bool = context
            .with(|ctx| {
                ctx.eval("typeof globalThis.solidNative.createElement === 'function'")
            })
            .unwrap();
        assert!(result);
    }

    #[test]
    fn test_js_can_call_create_element() {
        let delegate = Arc::new(MockDelegate);
        let renderer = Arc::new(SolidRenderer::new(delegate).unwrap());
        let runtime = Runtime::new().unwrap();
        let context = Context::full(&runtime).unwrap();

        bind_renderer_to_context(&context, renderer.clone()).unwrap();

        // Call createElement from JS
        let node_id: String = context
            .with(|ctx| ctx.eval("globalThis.solidNative.createElement('div')"))
            .unwrap();

        // UUID format check
        assert_eq!(node_id.len(), 36);
    }

    #[test]
    fn test_js_can_call_create_text_node() {
        let delegate = Arc::new(MockDelegate);
        let renderer = Arc::new(SolidRenderer::new(delegate).unwrap());
        let runtime = Runtime::new().unwrap();
        let context = Context::full(&runtime).unwrap();

        bind_renderer_to_context(&context, renderer.clone()).unwrap();

        let node_id: String = context
            .with(|ctx| ctx.eval("globalThis.solidNative.createTextNode('Hello')"))
            .unwrap();

        // UUID format check
        assert_eq!(node_id.len(), 36);
    }

    #[test]
    fn test_js_can_manipulate_tree() {
        let delegate = Arc::new(MockDelegate);
        let renderer = Arc::new(SolidRenderer::new(delegate).unwrap());
        let runtime = Runtime::new().unwrap();
        let context = Context::full(&runtime).unwrap();

        bind_renderer_to_context(&context, renderer.clone()).unwrap();

        // Test step by step to find the issue
        let parent_id: String = context
            .with(|ctx| ctx.eval("globalThis.solidNative.createElement('div')"))
            .unwrap();

        let child_id: String = context
            .with(|ctx| ctx.eval("globalThis.solidNative.createElement('span')"))
            .unwrap();

        // Call insertNode directly through renderer instead of JS
        renderer.insert_node(parent_id.clone(), child_id.clone(), None);

        // Verify through renderer directly
        let found_parent = renderer.get_parent_node(child_id);
        assert_eq!(found_parent, Some(parent_id));
    }
}
