pub mod paper_moon;


uniffi::include_scaffolding!("native_core");

fn add(a: u32, b: u32) -> u32 {
    (a + b).into()
}