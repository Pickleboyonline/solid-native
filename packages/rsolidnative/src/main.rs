use anyhow::{Ok, Result};
use log::info;
use rquickjs::{qjs::JSValue, Context, Runtime, Value};
use rquickjs_extra::console::{Console, Formatter};
use rsolidnative::hello;

fn main() -> Result<()> {
    env_logger::init();
    hello();

    let runtime = Runtime::new()?;

    // For a full async model, my guess is to have an actor
    // model, where every ctx invokation is followed
    // by a loop that exewcutes pending jobs
    // Every context call should do that.
    // Channels can help with this.
    let ctx = Context::full(&runtime)?;

    info!("erm?");

    ctx.with(|ctx| {
        let console = Console::new("hello", Formatter::default());
        ctx.globals().set("console", console).unwrap();
        ctx.eval::<(), _>("console.log('test')").unwrap();
    });

    let result = ctx.with(|ctx| {
        ctx.eval::<String, _>(
            r#"
        console.log("HELLO???");
        "pizza"
        "#,
        )
    })?;

    println!("Hello, world! {}", result);
    Ok(())
}
