use crate::paper_moon::{Renderer, ViewHostReceiver};
use anyhow::Result;
use rquickjs::{Context, Ctx, Exception, Function, Object, Runtime, Value};
use std::cell::RefCell;
use std::rc::Rc;
use std::sync::Arc;
use taffy::NodeId;
use uniffi::HandleAlloc;

pub mod paper_moon;
uniffi::include_scaffolding!("native_core");

// TODO: Make

struct NativeCore<T: ViewHostReceiver> {
    runtime: Runtime,
    context: Context,
    // Need interior mutability as Swift and JS (which is internal clojures)
    // need access to this.
    renderer: Arc<RefCell<Renderer<T>>>,
}

struct DummyViewHostReceiver {}

impl ViewHostReceiver for DummyViewHostReceiver {
    fn update_props(&self) {
        println!("You updated the props!")
    }
}

impl<T> NativeCore<T>
where
    T: ViewHostReceiver,
{
    /// Main constructor. Make the JS Engine and
    /// Swift/Kotlin will pass a reference to the receiver
    /// that handles creating it (itself, for example)
    pub fn new(dummy_view_host: T) -> Self {
        let runtime = Runtime::new().unwrap();
        let context = Context::full(&runtime).unwrap();

        // Since JS and Native need a ref we may need to wrap in an
        // Arc/Rc RefCell
        //
        let renderer = Arc::new(RefCell::new(Renderer::new(dummy_view_host)));

        // TODO: Make renderer available to JS side via a Class.

        Self {
            runtime,
            context,
            renderer,
        }
    }
    /// Download code and start JS engine
    /// TODO: Enventually handle bytecode support
    pub fn start_js_engine(&mut self) {
        self.context.with(|ctx| {
            // self.renderer..create_node("").unwrap();

            let v: Value = ctx.eval("1 + 3").unwrap();
            ctx.globals().set("", "").unwrap();
        });
    }

    pub fn create_root_node(&mut self, view_type: &str) -> Result<NodeId> {
        // self.renderer.create_root_node(view_type)
        Ok(NodeId::new(2))
    }

    /// TODO: we need to expose methods to the
    /// enclosed self value
    fn build_renderer_jsvalue<'a>(&'a self, ctx: Ctx<'a>) -> Result<Value> {
        let new_obj = Object::new(ctx.clone())?;

        let t = ctx.clone();
        
        let func = Function::new(ctx.clone(),  move || {
            // let c = ctx.clone();
            // handle_js_error(c, || {
            //     self.renderer.clone().borrow_mut().create_node("sn_text")
            // });
            // None
            let e = Exception::from_message(t.clone(), "message").unwrap();

            e
        })
        .unwrap()
        .with_name("inc")
        .unwrap();

        Ok(Value::new_bool(ctx, false))
    }
}

fn handle_js_error<T, U: FnOnce() -> Result<T>>(ctx: Ctx, f: U) -> Option<T> {
    let result = f();
    match result {
        Ok(y) => Some(y),
        // Ok(t) => Some(t),
        Err(_) => {
            //
            None
        }
    }
}

fn add(a: u32, b: u32) -> u32 {
    let runtime = Runtime::new().unwrap();

    let context = Context::full(&runtime).unwrap();

    let mut res: u32 = 0;

    context.with(|ctx| {
        let v: u32 = ctx.eval(format!("{} + {}", a, b)).unwrap();
        res = v;
        // println!("Value: {}", v)
    });

    res
}
