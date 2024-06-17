swiftc \
    -module-name                      \    # Name for resulting Swift module
    -emit-module -emit-module-path ./        \    # Output directory for resulting module
    -parse-as-library                        \   
    -L ./target/release/                       \    # Directory containing compiled Rust crate
    -libnative_core                                \    # Name of compiled Rust crate cdylib
    -Xcc -fmodule-map-file=native_coreFFI.modulemap  \   # The modulemap file from above
    native_core.swift                                \   # The generated bindings file


swiftc
    -module-name native_core                         # Name for resulting Swift module
    -emit-module -emit-module-path ./            # Output directory for resulting module
    -parse-as-library
    -L ./target/debug/                           # Directory containing compiled Rust crate
    -lexample                                    # Name of compiled Rust crate cdylib
    -Xcc -fmodule-map-file=exampleFFI.modulemap  # The modulemap file from above
    example.swift                                # The generated bindings file