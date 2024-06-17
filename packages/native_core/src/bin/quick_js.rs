use rquickjs::loader::Bundle;
use rquickjs::{embed, runtime, CatchResultExt, Context, Ctx, Exception, Function, Object, Runtime, Error, Value, Result};
use uniffi::HandleAlloc;
use std::cell::Cell;
use std::rc::Rc;
use std::sync::Arc;


macro_rules! expand_function {
    (
        $func_name:ident,
        $ctx:ident : $ctx_ty:ty,
        $( $arg:ident : $arg_ty:ty ),* $(,)?
        => $ret_ty:ty $body:block
    ) => {
        {
            fn $func_name($ctx: $ctx_ty, $( $arg : $arg_ty ),*) -> $ret_ty $body

        let func = Function::new($ctx.clone(), $func_name)
            .unwrap()
            .with_name(stringify!($func_name))
            .unwrap();

        func
        }
    };
}

/*
! load the `my_module. js` file and name it myModule
! use when in prod or just have js file.
! How to do bytecode =>
static BUNDLE: Bundle = embed! {
"myModule": "./my_module.js",
};
*/

struct QuickJs {
    runtime: Runtime,
    context: Context,
    value: Rc<Cell<u32>>,
}

impl QuickJs {
    fn new() -> Self {
        let runtime = Runtime::new().unwrap();

        let context = Context::full(&runtime).unwrap();


        Self {
            runtime,
            context,
            value: Rc::new(Cell::new(10)),
        }
    }

    fn run_js(&mut self) {
        self.context.with(|ctx: Ctx| {
            let value = self.value.clone();
            let closure_ctx = ctx.clone();


            // fn inc(ctx: Ctx) -> Result<Object> {
            //     // Ok(Object::new(ctx).unwrap())
            //     Err(Error::Exception)
            // }

            let inc = expand_function!(inc, ctx: Ctx, a: i32, b: i32 => i32 {
                    // Err(Error::Exception)
                a+b
            });


            ctx.globals().prop("inc", inc).unwrap();

            let res: Value = ctx.eval("try { inc(7, 10 ) } catch (e) { 2 }").unwrap();

            println!("Bruh: {:?}", res);
            println!("Res: {:?}", ctx.catch());
            println!("Res: {:?}", ctx.catch());
            println!("Value: {}", self.value.get())
        })
    }

    fn inc_value(&mut self) {
        self.value.set(self.value.get() + 1);
    }
}


fn adder(a: u32, b: u32) -> u32 {
    println!("Hi from rust!");
    a + b
}

/// Right now we are using QuickJS since it has better docs
/// I recently learned C/C++ and build tools so I could use
/// v8 or duktape, both have a debugger.
/// TODO: Determine how big v8 executables are.
/// Duktape has a size of ~1 MB => extremely tiny.
/// You could also use v8 for debug apps and QuickJS for running
/// https://duktape.org/ <= has debugger
/// Duktape does have rust bindings but they are old.


fn old_js() {
    let runtime = Runtime::new().unwrap();

    let context = Context::full(&runtime).unwrap();

    context.with(|ctx| {
        let func = Function::new(ctx.clone(), adder)
            .unwrap()
            .with_name("adder")
            .unwrap();

        ctx.globals().prop("adder", func).unwrap();

        let v: u64 = ctx.eval("adder(1,2)").unwrap();

        println!("Value: {}", v)
    });
}

fn main() {
    let mut q = QuickJs::new();
    q.run_js();
}