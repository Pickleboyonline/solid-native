# Changelog

## [Unreleased]

### Added
- **Text descriptor system** for React Native-style text rendering
  - TextDescriptor struct with text content and styles (HashMap<String, JSValue>)
  - generate_text_descriptors() function that walks text node hierarchy
  - Merges parent styles with child styles (children override parents)
  - on_text_descriptors_change callback in HostDelegate
  - Automatic text descriptor generation on text node changes (insert, remove, property set)
  - Swift SNHostDelegateImpl handles text descriptors
- **JSValue enum** for representing JavaScript values across FFI boundary
  - Supports: Null, Undefined, Boolean, Number, String, Array, Object
  - Fully uniffi-compatible with HashMap-based objects
  - Utility methods for type checking and value extraction
- **on_prop_updated callback** in HostDelegate trait
  - Called when properties are updated on nodes
  - Receives JSValue for type-safe property values
  - Integrated with SolidRenderer.set_property()
- **Root node management**
  - UITree.set_root() - Sets a node as the root
  - UITree.get_root() - Gets the root node key
  - UITree.get_root_id() - Gets the root node ID
  - SolidRenderer.create_root() - Creates and sets a root node
  - SolidRenderer.set_root() - Sets an existing node as root
  - SolidRenderer.get_root() - Gets the root node ID
  - SolidNativeCore exposes all root node methods via uniffi
- **Thread safety** for SolidNativeCore
  - Implemented Send + Sync traits (with unsafe impl)
  - Safe Arc + Mutex synchronization for runtime and context
- **SNCoreError enum** for proper error handling
  - RuntimeError, RendererError, JsEvalError, ContextError
  - uniffi-compatible with thiserror integration
  - Replaces String-based errors

### Changed
- HostDelegate trait methods now use owned types (String, Vec<String>) instead of references
  - Better uniffi compatibility
  - Updated all delegate callbacks and implementations
- SolidNativeCore.new() now returns Result<Arc<Self>, SNCoreError>
  - Properly wrapped for uniffi object lifetime management
- Property updates now trigger on_prop_updated callback in addition to on_update_revision_count

### Fixed
- QuickJS Runtime and Context not being thread-safe
  - Added unsafe Send + Sync implementations with proper synchronization
- Uniffi tuple support issues
  - Changed JSValue::Object from Vec<(String, JSValue)> to HashMap<String, JSValue>

## [0.1.0] - Initial Release

### Added
- SolidNativeCore struct for managing JavaScript runtime and renderer
- SolidRenderer for UI tree management
- UITree with normalized node storage using slotmap
- HostDelegate trait for platform callbacks
- JavaScript bindings via QuickJS
- Basic uniffi exports for Swift/Kotlin
- Node creation, insertion, removal, and property setting
- Tree traversal methods (parent, first child, next sibling)
