// Module declarations
pub mod core;
pub mod delegate;
pub mod error;
pub mod js_bindings;
pub mod node;
pub mod renderer;
pub mod tree;

// Re-export main types for convenience
pub use core::SolidNativeCore;
pub use delegate::HostDelegate;
pub use error::SNCoreError;
pub use node::{Node, NodeKey, NodeType};
pub use renderer::SolidRenderer;
pub use tree::UITree;

uniffi::setup_scaffolding!();

// Example of using uniffi
// Uses JS Engine for it.
// #[uniffi::export]
// fn add(a: u32, b: u32) -> u32 {
//     let runtime = Runtime::new().unwrap();
//     let context = Context::full(&runtime).unwrap();

//     let expression = format!("{} + {} + 1", a, b);

//     let result = context.with(|ctx| ctx.eval::<u32, _>(expression).unwrap());
//     result
// }

/*
What's needed:
- 1 SNCore object that handles the JS engine.
- It should be able to take a reference to the "receiver" object
- We are using the delegate pattern here
- The host platform receives function calls and data
*/
