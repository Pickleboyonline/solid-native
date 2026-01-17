# Solid-Native Dev-App

Self-contained development app for Solid Native with class-based architecture.

## Project Structure

```
dev-app/
├── deno.json                    # Deno config with solid-native import
├── src/
│   ├── index.ts                 # Entry point
│   ├── App.tsx                  # Main app component
│   ├── server.ts                # Dev server (esbuild + serve)
│   └── bundle.ts                # Bundle script
├── packages/
│   ├── solid-native/            # JS component package
│   │   ├── mod.ts               # Main exports
│   │   ├── renderer.ts          # SolidJS universal renderer
│   │   └── views/
│   │       ├── mod.ts           # Re-exports
│   │       ├── types.ts         # FlexStyle, TextStyle
│   │       ├── view.tsx         # View component
│   │       ├── text.tsx         # Text component
│   │       └── image.tsx        # Image component
│   ├── sncore/                  # Rust core (FFI to JS runtime)
│   ├── sn-swiftui/              # Swift UI layer
│   │   └── Sources/SNSwiftUI/
│   │       ├── Views/
│   │       │   ├── util/
│   │       │   │   ├── SolidNativeView.swift      # Protocol + BaseView
│   │       │   │   ├── SolidNativeViewWrapper.swift # Wrapper class
│   │       │   │   └── Color+Hex.swift            # Hex color support
│   │       │   ├── SNView.swift               # Container (uses Flex)
│   │       │   ├── SNText.swift               # Text with styling
│   │       │   └── SNImage.swift              # AsyncImage
│   │       ├── HostReceiver.swift             # Registry + delegate
│   │       └── SNSwiftUI.swift                # Public exports
│   └── Yoga-SwiftUI/            # Submodule for flexbox layout
└── ios/
    └── DevApp/
        ├── DevAppApp.swift      # App entry
        └── ContentView.swift    # Main view using HostReceiver
```

## Architecture

### Key Simplifications
- **No `layoutMetrics`** - Yoga-SwiftUI handles layout directly in SwiftUI
- **No ZStack positioning** - Props pass to Yoga flexbox view modifiers
- **Views as package** - `packages/solid-native/` with deno.json import mapping

### Data Flow
```
JS (SolidJS) → Rust (sncore) → Swift (HostReceiver) → SwiftUI
```

1. SolidJS renderer calls `solidNative.*` APIs
2. Rust core processes and calls `HostDelegate` callbacks
3. `HostReceiver` updates `SolidNativeViewWrapper` props/children
4. SwiftUI observes wrapper changes and re-renders

### Component Registration
Components are registered in `HostReceiver.swift`:
```swift
let viewTypeRegistry: [String: any SolidNativeView.Type] = [
    SNView.name: SNView.self,   // "sn_view"
    SNText.name: SNText.self,   // "sn_text"
    SNImage.name: SNImage.self, // "sn_image"
]
```

### Creating New Components

**JS Side** (`packages/solid-native/views/`):
```typescript
export function MyComponent(props: MyProps) {
  return <sn_mycomponent {...props} />;
}
```

**Swift Side** (`packages/sn-swiftui/Sources/SNSwiftUI/Views/`):
```swift
class SNMyComponent: SolidNativeView {
    static var name: String { "sn_mycomponent" }

    func render() -> some View {
        // Access props via self.props
        // Access children via self.children
    }
}
```

Then register in `HostReceiver.viewTypeRegistry`.

## Commands

```bash
# Start dev server (bundles JS and serves on :8080)
deno task dev

# Build bundle only
deno task bundle

# Build Rust core
cd packages/sncore && just build-ios
```

## Key Files

| File | Purpose |
|------|---------|
| `packages/solid-native/renderer.ts` | SolidJS ↔ Rust binding |
| `packages/sn-swiftui/.../HostReceiver.swift` | Swift component registry |
| `packages/sn-swiftui/.../SolidNativeView.swift` | Component protocol |
| `packages/sn-swiftui/.../SNView.swift` | Container with Flex layout |
| `ios/DevApp/ContentView.swift` | iOS app integration |

## Props System

Props are `[String: JsValue]` where `JsValue` is an enum:
```swift
public enum JsValue {
    case null
    case undefined
    case boolean(value: Bool)
    case number(value: Double)
    case string(value: String)
    case array(values: [JsValue])
    case object(properties: [String: JsValue])
}
```

Extract style props:
```swift
if let style = props["style"], case .object(let styleProps) = style {
    if let color = styleProps["color"], case .string(let hex) = color {
        // use hex color
    }
}
```

## Layout

SNView uses Yoga-SwiftUI's `Flex` container:
```swift
Flex(
    direction: .column,
    justifyContent: .center,
    alignItems: .center
) {
    ForEach(children, id: \.self) { nodeId in
        childWrapper.render()
            .applyFlexChildProps(from: childWrapper.props)
    }
}
```

Flex props are extracted from `style` object and mapped to Yoga types.

## Dependencies

- **SNCore**: Rust FFI layer (uniffi-generated)
- **Yoga-SwiftUI**: Flexbox layout for SwiftUI
- **yoga**: Facebook's layout engine (via Yoga-SwiftUI)
- **solid-js**: Reactive UI framework
- **esbuild**: JS bundling
