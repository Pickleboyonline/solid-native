use rquickjs::{runtime, Context, Runtime};

// https://duktape.org/ <= has debugger
fn main() {
    let runtime = Runtime::new().unwrap();

    let context = Context::full(&runtime).unwrap();

    context.with(|ctx| {
        let v: u64 = ctx.eval("1 + 2").unwrap();

        println!("Value: {}", v)
    })
    
}

// use deno_core::error::AnyError;
// use deno_core::{extension, op2, PollEventLoopOptions};
// use std::env;
// use std::rc::Rc;

// // How to write functions
// #[op2(fast)]
// fn op_adder(num: f64) -> f64 {

//     println!("Hi, im from rust!");
//     num + 1.0
// }

// async fn run_js(file_path: &str) -> Result<(), AnyError> {
//     let main_module = deno_core::resolve_path(file_path, env::current_dir().unwrap().as_path())?;

//     extension!(my_extension, ops = [op_adder]);
//     /*
//      * TODO: We need to improve runtime
//      * - Expose renderer to runtime
//      * - Expose set timeout.
//      * - For now, we can just resolve the path from the URL since it support it like
//      *   iOS variant JSCore
//      * - Javascript HostObject thing for JSI like experience or just some macros.
//      * - We use a library and hook into it via "JSI". Code in Rust.
//      * - Project will prob just have a build macro that links binaries for rust
//      * - Detemine Rust/Swift/Kotline interop with JSI and UniFFI
//      *   UniFFI handles talking between Rust <=> Kotline/Swift.
//      *
//      * For now, dont worry about the module interop. Just get it working in this project in an ad-hoc fashion
//      * Then refactor to modularize it.
//      */
//     let mut js_runtime = deno_core::JsRuntime::new(deno_core::RuntimeOptions {
//         module_loader: Some(Rc::new(deno_core::FsModuleLoader)),
//         extensions: vec![my_extension::init_ops_and_esm()],
//         ..Default::default()
//     });

//     let mod_id = js_runtime.load_main_es_module(&main_module).await?;
//     let result = js_runtime.mod_evaluate(mod_id);
//     js_runtime
//         .run_event_loop(PollEventLoopOptions::default())
//         .await?;
//     result.await
// }


// fn set_flags() {
//     deno_core::v8_set_flags(vec![
//         "".to_owned(),
//         "--jitless".to_owned(),
//         "--no-expose-wasm".to_owned(),
//         // Enable to save memory to trade for performnace
//         // "--lite-mode".to_owned()
//     ]);
// }

// fn main() {
//     // ! Immediatly configure the engine not to use jit. Prob need to have some check however as Android
//     // ! can support it.
//     // set_flags();

//     let runtime = tokio::runtime::Builder::new_current_thread()
//         .enable_all()
//         .build()
//         .unwrap();

//     if let Err(error) = runtime.block_on(run_js("./example.js")) {
//         eprintln!("error: {}", error);
//     }
// }
