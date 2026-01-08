use rquickjs::{Context, Runtime};

uniffi::setup_scaffolding!();

/// Uses JS Engine for it.
#[uniffi::export]
fn add(a: u32, b: u32) -> u32 {
    let runtime = Runtime::new().unwrap();
    let context = Context::full(&runtime).unwrap();

    let expression = format!("{} + {} + 1", a, b);

    let result = context.with(|ctx| {
        ctx.eval::<u32, _>(expression).unwrap()
    });
    // a + b
    result
}
