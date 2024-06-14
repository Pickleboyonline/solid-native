use deno_core::error::AnyError;
use deno_core::PollEventLoopOptions;
use std::env;
use std::rc::Rc;

async fn run_js(file_path: &str) -> Result<(), AnyError> {
    let flags = deno_core::v8_set_flags(vec![
        "".to_owned(),
        "--jitless".to_owned(),
        "--no-expose-wasm".to_owned()
    ]);

    println!("Unsupported flags: {:?}", flags);

    let main_module = deno_core::resolve_path(file_path, env::current_dir().unwrap().as_path())?;

    let mut js_runtime = deno_core::JsRuntime::new(deno_core::RuntimeOptions {
        module_loader: Some(Rc::new(deno_core::FsModuleLoader)),
        ..Default::default()
    });

    let mod_id = js_runtime.load_main_es_module(&main_module).await?;
    let result = js_runtime.mod_evaluate(mod_id);
    js_runtime
        .run_event_loop(PollEventLoopOptions::default())
        .await?;
    result.await
}

fn main() {
    let runtime = tokio::runtime::Builder::new_current_thread()
      .enable_all()
      .build()
      .unwrap();

    if let Err(error) = runtime.block_on(run_js("./example.js")) {
      eprintln!("error: {}", error);
    }
  }