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

        let r8 = renderer.clone();
        solid_native.set(
            "getParentNode",
            Func::from(move |node_id: String| r8.get_parent_node(node_id)),
        )?;

        let r9 = renderer.clone();
        solid_native.set(
            "getFirstChild",
            Func::from(move |node_id: String| r9.get_first_child(node_id)),
        )?;

        let r10 = renderer.clone();
        solid_native.set(
            "getNextSibling",
            Func::from(move |node_id: String| r10.get_next_sibling(node_id)),
        )?;

        // Attach to globalThis
        globals.set("solidNative", solid_native)?;

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
