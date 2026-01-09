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

### Integration with SNCore

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

            // Create a root node (IMPORTANT: Always create a root first!)
            let rootId = try manager.createRootFromCore(tag: "vstack")

            // Run JavaScript that builds UI
            let jsCode = """
            // Use the root node we created
            const rootId = '\(rootId)';

            solidNative.setProperty(rootId, 'spacing', '16');
            solidNative.setProperty(rootId, 'padding', '20');

            const button = solidNative.createElement('button');
            solidNative.setProperty(button, 'backgroundColor', 'blue');
            solidNative.setProperty(button, 'cornerRadius', '8');

            const text = solidNative.createTextNode('Hello from JS!');
            solidNative.setProperty(text, 'color', 'white');

            solidNative.insertNode(button, text);
            solidNative.insertNode(rootId, button);
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

**With Rust Core:**
- `initializeWithCore()` - Initializes the Rust core with the Swift delegate
- `createRootFromCore(tag:)` - Creates a root node in Rust and syncs to Swift (use this!)
- `getRootFromCore()` - Gets the root node ID from Rust core
- `evaluateJavaScript(_:)` - Evaluates JavaScript code in the Rust runtime
- `evaluateModule(_:name:)` - Evaluates JavaScript module code

**Manual Tree Building (without Rust Core):**
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

## Root Node Management

**IMPORTANT:** You must create a root node before building your UI tree. The root node is the top-level container for all UI elements.

### Creating a Root Node

**With Rust Core (Recommended):**
```swift
let rootId = try manager.createRootFromCore(tag: "vstack")
```

**Manual (for testing without Rust):**
```swift
let root = manager.createRoot(type: "vstack")
```

### Why Do You Need a Root Node?

The root node serves as:
1. **Entry point** - The starting point for rendering the UI tree
2. **Container** - Holds all child elements
3. **Synchronization point** - Keeps Rust and Swift trees in sync

### Methods Available in Rust Core

```swift
// Create and set root
core.createRoot(tag: "vstack") -> String  // Returns root node ID

// Set existing node as root
core.setRoot(nodeId: String) -> Bool  // Returns true if successful

// Get current root
core.getRoot() -> String?  // Returns root node ID or nil
```

## Examples

Check the `Examples/` directory for complete examples:
- **BasicExample.swift** - Using SNSwiftUI with the Rust core and JavaScript
- **ManualTreeExample.swift** - Building trees manually without Rust core

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
3. Import both `SNSwiftUI` and `SNCore` in your Swift files

## TODO

- [x] Add uniffi conformance for SNCore exports
- [x] Root node management
- [ ] Implement event handling (button presses, gestures)
- [ ] Add more SwiftUI components (List, TextField, etc.)
- [ ] Support for animations and transitions
- [ ] Better property parsing and validation
- [ ] Performance optimizations for large trees

## License

MIT

## Contributing

Contributions welcome! Please ensure tests pass before submitting PRs.
