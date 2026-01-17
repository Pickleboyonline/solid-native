# SNCore

Rust core library for Solid Native. Bridges JavaScript (via QuickJS) and native platforms (iOS/Android) using UniFFI for FFI bindings.

## Build Commands

```bash
# Build iOS framework (output: build/swift/SNCore/)
just build-ios

# Build Android library (output: android-sncore/)
just build-android

# Build all platforms
just build-all

# Run tests
cargo test
```

## Output Locations

- **iOS**: `build/swift/SNCore/` - Auto-generated Swift package
- **Android**: `android-sncore/` - Requires manual config

The rationale: iOS bindings are completely auto-generated via `cargo-swift`, whereas Android needs additional configuration.

## Architecture

### Core Components

| Module | Purpose |
|--------|---------|
| `core.rs` | `SolidNativeCore` - Main entry point, manages QuickJS runtime |
| `renderer.rs` | `SolidRenderer` - UI tree management, notifies host via delegate |
| `delegate.rs` | `HostDelegate` trait - Callbacks implemented by Swift/Kotlin |
| `tree.rs` | `UITree` - Slotmap-based tree structure |
| `jsvalue.rs` | `JSValue` - Cross-platform JS value representation |
| `js_bindings.rs` | Binds renderer to `globalThis.solidNative` |
| `node.rs` | `Node` types (Element, Text) |
| `text.rs` | `TextDescriptor` for styled text segments |
| `error.rs` | Error types (Runtime, Renderer, JsEval, Context) |

### JS Bindings (`globalThis.solidNative`)

```javascript
solidNative.createElement(tag)        // Returns node ID (UUID)
solidNative.createTextNode(value)     // Returns node ID
solidNative.setProp(nodeId, key, val) // Auto-serializes value to JSON
solidNative.insertBefore(parent, node, anchor?)
solidNative.removeNode(parent, node)
solidNative.getRootView()             // Returns root node ID
solidNative.getParentNode(nodeId)
solidNative.getFirstChild(nodeId)
solidNative.getNextSibling(nodeId)
```

### HostDelegate Callbacks

Swift/Kotlin implements these to receive UI updates:

```rust
fn on_node_created(&self, node_id: String, node_type: String);
fn on_prop_updated(&self, node_id: String, key: String, value: JSValue);
fn on_children_change(&self, node_id: String, node_ids: Vec<String>);
fn on_text_descriptors_change(&self, node_id: String, descriptors: Vec<TextDescriptor>);
fn on_update_revision_count(&self, node_id: String);  // Signals UI refresh
```

## Key Implementation Details

- Node IDs are UUIDs (36 characters)
- JS context persists across `eval_js` calls (state is maintained)
- `setProp` uses a JS wrapper that calls `JSON.stringify` before passing to Rust
- Properties stored as JSON internally, converted to `JSValue` for callbacks
- Error handling extracts actual JS exception messages via `ctx.catch()`

## Dependencies

- `rquickjs` - QuickJS JavaScript engine bindings
- `uniffi` - FFI bindings generation for Swift/Kotlin
- `slotmap` - Efficient tree storage
- `serde_json` - JSON serialization
- `uuid` - Node ID generation
- `thiserror` - Error handling
