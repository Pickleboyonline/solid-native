fn main() {
    uniffi::generate_scaffolding("./src/native_core.udl").unwrap();
}