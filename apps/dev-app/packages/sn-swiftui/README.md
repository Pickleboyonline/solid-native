# SNSwiftUI

A Swift package for managing iOS/macOS UI trees with SwiftUI, designed to work with the SolidNative Core Rust library.

## Overview

SNSwiftUI provides a SwiftUI-based rendering system that integrates with the SolidNative Core (SNCore) Rust library via uniffi bindings. It manages a virtual tree of view nodes that can be manipulated from JavaScript running in the Rust QuickJS runtime and renders them as native SwiftUI views.

## Features

- **View Tree Management**: Thread-safe view tree with efficient node lookup
- **HostDelegate Implementation**: Bridges Rust renderer callbacks to SwiftUI
- **SwiftUI Rendering**: Renders virtual nodes as native SwiftUI views
- **Flexible Layout**: Supports VStack, HStack, ZStack, ScrollView, and more
- **Styling Support**: Properties for colors, padding, corner radius, frames, etc.
- **Text Nodes**: Efficient text rendering
- **Manual Tree Building**: Test and develop without the full Rust runtime

## Architecture

```
┌─────────────────────┐
│   JavaScript Code   │
│   (QuickJS Runtime) │
└──────────┬──────────┘
           │
           ▼
┌─────────────────────┐
│    SNCore (Rust)    │
│  Renderer + Runtime │
└──────────┬──────────┘
           │ uniffi
           ▼
┌─────────────────────┐
│  SNHostDelegateImpl │
│   (Swift Bridge)    │
└──────────┬──────────┘
           │
           ▼
┌─────────────────────┐
│      ViewTree       │
│   (State Manager)   │
└──────────┬──────────┘
           │
           ▼
┌─────────────────────┐
│ SNSwiftUIRenderer   │
│  (SwiftUI Views)    │
└─────────────────────┘
```

## Installation

Add this package as a dependency in your Xcode project or Swift Package:

```swift
dependencies: [
    .package(path: "../packages/sn-swiftui")
]
```

## Usage

### Basic Setup

```swift
import SwiftUI
import SNSwiftUI

struct ContentView: View {
    @StateObject private var manager = SNSwiftUIManager()

    var body: some View {
        SNSwiftUIRenderer(viewTree: manager.viewTree)
            .onAppear {
                buildUI()
            }
    }

    private func buildUI() {
        // Create a simple UI tree
        let root = manager.createRoot(type: "vstack")
        manager.setProperty(node: root, key: "spacing", value: "16")
        manager.setProperty(node: root, key: "padding", value: "20")

        // Add a button
        let button = manager.createElement(type: "button")
        manager.setProperty(node: button, key: "backgroundColor", value: "blue")
        manager.setProperty(node: button, key: "cornerRadius", value: "8")
        manager.setProperty(node: button, key: "padding", value: "12")

        let buttonText = manager.createTextNode(text: "Click Me")
        manager.setProperty(node: buttonText, key: "color", value: "white")

        manager.insertNode(parent: button, child: buttonText)
        manager.insertNode(parent: root, child: button)

        // Add a text label
        let label = manager.createTextNode(text: "Hello, SwiftUI!")
        manager.setProperty(node: label, key: "color", value: "gray")

        manager.insertNode(parent: root, child: label)
    }
}
```

### Integration with SNCore (once uniffi exports are available)

```swift
import SNSwiftUI
import SNCore

struct ContentView: View {
    @StateObject private var manager = SNSwiftUIManager()

    var body: some View {
        SNSwiftUIRenderer(viewTree: manager.viewTree)
            .onAppear {
                initializeCoreAndRunJS()
            }
    }

    private func initializeCoreAndRunJS() {
        do {
            // Initialize the Rust core with our delegate
            try manager.initializeWithCore()

            // Run JavaScript that builds UI
            let jsCode = """
            const root = solidNative.createElement('vstack');
            solidNative.setProperty(root, 'spacing', '16');

            const button = solidNative.createElement('button');
            const text = solidNative.createTextNode('Hello from JS!');

            solidNative.insertNode(button, text);
            solidNative.insertNode(root, button);
            """

            try manager.evaluateJavaScript(jsCode)
        } catch {
            print("Error: \\(error)")
        }
    }
}
```

## Supported View Types

| Type | Description | SwiftUI Equivalent |
|------|-------------|-------------------|
| `vstack` | Vertical stack | `VStack` |
| `hstack` | Horizontal stack | `HStack` |
| `zstack` | Depth stack | `ZStack` |
| `scrollview` | Scrollable container | `ScrollView` |
| `button` | Interactive button | `Button` |
| `image` | Image view | `Image` |
| `text` | Text node | `Text` |
| `spacer` | Flexible space | `Spacer` |
| `view`, `div`, `container` | Generic container | `VStack` |

## Supported Properties

- **Layout**: `padding`, `spacing`, `width`, `height`
- **Colors**: `backgroundColor`, `foregroundColor`, `color`
- **Styling**: `cornerRadius`
- **Image**: `src`, `name` (for system images)

### Color Values

Colors can be specified as:
- Named colors: `"red"`, `"blue"`, `"green"`, etc.
- Hex colors: `"#FF0000"`, `"#00FF00AA"`

## API Reference

### SNSwiftUIManager

The main manager class that coordinates the view tree and rendering.

#### Methods

- `createRoot(type:)` - Creates and sets the root node
- `createElement(type:)` - Creates a new element node
- `createTextNode(text:)` - Creates a new text node
- `insertNode(parent:child:before:)` - Inserts a child into a parent
- `removeNode(parent:child:)` - Removes a child from a parent
- `setProperty(node:key:value:)` - Sets a property on a node
- `clear()` - Clears the entire tree
- `debugPrintTree()` - Prints tree structure for debugging

### SNHostDelegateImpl

Implements the HostDelegate protocol from SNCore.

#### Methods (matching Rust trait)

- `onNodeCreated(nodeId:nodeType:)` - Called when a node is created
- `onNodeRemoved(nodeId:)` - Called when a node is removed
- `onChildrenChange(nodeId:nodeIds:)` - Called when children change
- `onUpdateRevisionCount(nodeId:)` - Signals UI should update
- `isTextElementByNodeId(nodeId:)` - Checks if node is text
- `isTextElementByNodeType(nodeType:)` - Checks if type is text

### ViewTree

Thread-safe tree management.

#### Methods

- `addNode(_:)` - Add a node to the index
- `getNode(_:)` - Get a node by ID
- `removeNode(_:)` - Remove a node by ID
- `setRoot(_:)` - Set the root node
- `clear()` - Clear all nodes
- `printTree()` - Print tree structure

## Development

### Building

```bash
cd packages/sn-swiftui
swift build
```

### Testing

```bash
swift test
```

### Integration with Xcode Project

1. Add the package to your Xcode project via File > Add Packages
2. Select "Add Local..." and choose the `sn-swiftui` directory
3. Import `SNSwiftUI` in your Swift files

## TODO

- [ ] Add uniffi conformance once SNCore exports are ready
- [ ] Implement event handling (button presses, gestures)
- [ ] Add more SwiftUI components (List, TextField, etc.)
- [ ] Support for animations and transitions
- [ ] Better property parsing and validation
- [ ] Performance optimizations for large trees

## License

MIT

## Contributing

Contributions welcome! Please ensure tests pass before submitting PRs.
