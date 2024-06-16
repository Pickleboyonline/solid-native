use rquickjs::{Context, Runtime};

pub mod paper_moon;


uniffi::include_scaffolding!("native_core");

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