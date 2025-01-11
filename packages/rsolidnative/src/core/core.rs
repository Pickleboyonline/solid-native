use rquickjs::{Context, Function, Object, Runtime, Value, Ctx};
use std::error::Error;

const GLOBAL_CORE_NAME: &str = "_SolidNativeCore";

pub struct Core {
    rt: Runtime,
    ctx: Context,
}

pub trait Module {
    fn define<'a, 'b>(&self, ctx: Ctx<'b>) -> ModuleDefinition<'a>;
}

pub trait HostModule {
    fn define(&self, wrapper: &HostContextWrapper) -> HostModuleDefinition;
}

pub struct ModuleDefinition<'a> {
    pub name: String,
    pub value: Value<'a>,
}

pub struct HostModuleDefinition {
    pub name: String,
}

pub struct HostContextWrapper<'a, 'b> {
    ctx: &'a Context,
    obj: Object<'b>,
}

impl Core {
    pub fn new() -> Result<Self, Box<dyn Error>> {
        let rt = Runtime::new()?;
        let ctx = Context::full(&rt)?;

        ctx.with(|ctx| -> Result<(), rquickjs::Error> {
            let globals = ctx.globals();

            ctx.spawn();
        
            // Create core object
            let core = Object::new(ctx.clone())?;
            let modules = Object::new(ctx.clone())?;
            
            core.set("modules", modules)?;
            globals.set(GLOBAL_CORE_NAME, core)?;
            
            Ok(())
        })?;

        Ok(Core { rt, ctx })
    }

    pub fn register_go_module(&self, module: impl Module) -> Result<(), Box<dyn Error>> {

        self.ctx.with(move |ctx| {
            let definition = module.define(ctx.clone());
            let globals = ctx.globals();

            let core: Object = globals.get(GLOBAL_CORE_NAME)?;
            let modules: Object = core.get("modules")?;
            
            modules.set(definition.name, definition.value)?;
            
            Ok(())
        })


    }

    fn do_something<'a>(module: impl Module + 'a, ctx: Ctx<'a>) -> ModuleDefinition<'a> {
        module.define(ctx)
    }

    pub async fn start_from_server(&self, url: &str) -> Result<(), Box<dyn Error>> {
        let response = reqwest::get(url).await?;
        let js_to_eval = response.text().await?;
        
        self.ctx.with(|ctx| {
            ctx.eval::<(), _>(js_to_eval)?;
            Ok(())
        })
    }

    pub fn start_from_js(&self, js: &str) -> Result<(), Box<dyn Error>> {
        self.ctx.with(|ctx| {
            ctx.eval::<(), _>(js)?;
            Ok(())
        })
    }

    pub fn start_with_debugger(&self, _debug_server_url: &str) -> Result<(), Box<dyn Error>> {
        // TODO: Implement debugger functionality
        Ok(())
    }
}

// The Drop trait implementation isn't needed as rquickjs handles cleanup automatically

// This won't compile
struct Container<'a> {
    value: &'a str
}

trait Process {
    fn process(&self, input: &str) -> Container;  // Error: Container needs a lifetime parameter
}

// This compiles
// trait Process {
//     fn process<'a>(&self, input: &'a str) -> Container<'a>;
// }

fn dsa(x: impl Process) {
    // let y = x.process();
    let x = x.process("dsa");
}