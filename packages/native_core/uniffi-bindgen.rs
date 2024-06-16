

fn main() {
    uniffi::uniffi_bindgen_main()
    // uniffi::generate_bindings(udl_file, config_file_override, binding_generator, out_dir_override, library_file, crate_name, try_format_code)

    /*
    Tut on iOS:
    https://krirogn.dev/blog/integrate-rust-in-ios
    https://github.com/antoniusnaumann/cargo-swift

    ! Have to build iOS from source to make it work :(
    ! Ensure CCache works. Build takes like 10 min
    https://crates.io/crates/v8
     */
}