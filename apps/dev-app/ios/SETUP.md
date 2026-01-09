# DevApp Setup Guide

This guide explains how to set up and run the DevApp iOS application with SolidNative.

## Prerequisites

- Xcode 15.0 or later
- macOS 14.0 or later
- Rust toolchain installed
- `just` command runner installed

## Project Structure

```
ios/DevApp/
├── DevApp.xcodeproj/       # Xcode project
├── DevApp/
│   ├── DevAppApp.swift     # App entry point
│   ├── ContentView.swift   # Main UI with SolidNative integration
│   └── Assets.xcassets/    # App assets
```

## Setup Steps

### 1. Build iOS Bindings

First, build the Rust library and generate Swift bindings:

```bash
cd apps/dev-app/packages/sncore
just build-ios
```

This will:
- Build the Rust library for iOS/macOS
- Generate Swift bindings via uniffi
- Create the SNCore package at `build/swift/SNCore/`

### 2. Add Swift Packages to Xcode

Open the Xcode project:

```bash
cd apps/dev-app/ios
open DevApp.xcodeproj
```

Add the local Swift packages:

1. **Add SNCore Package:**
   - File → Add Package Dependencies
   - Click "Add Local..."
   - Navigate to: `apps/dev-app/packages/sncore/build/swift/SNCore`
   - Click "Add Package"
   - Select your target (DevApp) and click "Add Package"

2. **Add SNSwiftUI Package:**
   - File → Add Package Dependencies
   - Click "Add Local..."
   - Navigate to: `apps/dev-app/packages/sn-swiftui`
   - Click "Add Package"
   - Select your target (DevApp) and click "Add Package"

### 3. Verify Package Dependencies

In your target's "Frameworks, Libraries, and Embedded Content" section, you should see:
- SNCore
- SNSwiftUI

### 4. Build and Run

1. Select your target device/simulator
2. Build and run (⌘R)

## How It Works

### ContentView.swift

The main view demonstrates:

1. **Initialization:**
   ```swift
   try manager.initializeWithCore()
   ```
   - Creates Rust runtime and renderer
   - Connects Swift delegate to Rust

2. **Root Node Creation:**
   ```swift
   let rootId = try manager.createRootFromCore(tag: "vstack")
   ```
   - Creates root container in Rust
   - Syncs to Swift view tree
   - **IMPORTANT:** Must be done before building UI

3. **JavaScript UI Building:**
   ```swift
   try manager.evaluateJavaScript(jsCode)
   ```
   - Runs JS code in QuickJS runtime
   - Uses `solidNative` global API
   - Builds UI tree that renders as SwiftUI

### JavaScript API

Available in `solidNative` global:

```javascript
// Create elements
const elementId = solidNative.createElement('button');
const textId = solidNative.createTextNode('Hello!');

// Set properties
solidNative.setProperty(elementId, 'backgroundColor', 'blue');
solidNative.setProperty(elementId, 'cornerRadius', '8');

// Build tree
solidNative.insertNode(parentId, childId);
solidNative.removeNode(parentId, childId);
```

## Supported View Types

- `vstack` - Vertical stack (VStack)
- `hstack` - Horizontal stack (HStack)
- `zstack` - Depth stack (ZStack)
- `scrollview` - Scrollable container
- `button` - Interactive button
- `image` - Image view (system images)
- `text` - Text node
- `spacer` - Flexible space

## Supported Properties

- Layout: `padding`, `spacing`, `width`, `height`
- Colors: `backgroundColor`, `foregroundColor`, `color`
- Styling: `cornerRadius`
- Image: `src`, `name`

## Troubleshooting

### "Module 'SNCore' not found"

- Ensure you've built the iOS bindings: `just build-ios`
- Check that SNCore package is added in Xcode
- Verify package path points to `build/swift/SNCore/`

### "Module 'SNSwiftUI' not found"

- Ensure SNSwiftUI package is added in Xcode
- Verify package path points to `packages/sn-swiftui/`

### Build Errors

If you get build errors after updating Rust code:

1. Rebuild iOS bindings:
   ```bash
   cd apps/dev-app/packages/sncore
   just build-ios
   ```

2. In Xcode:
   - Product → Clean Build Folder (⇧⌘K)
   - Product → Build (⌘B)

### Runtime Errors

Enable detailed logging:

```swift
// In ContentView.swift, check the console for:
print("🚀 Initializing SolidNative Core...")
print("✅ Core initialized")
print("✅ Root node created with ID: ...")
```

## Development Workflow

1. Make changes to Rust code in `packages/sncore/src/`
2. Rebuild iOS bindings: `just build-ios`
3. Xcode will automatically pick up changes
4. Build and run in Xcode

## Next Steps

- Add event handling (button presses, gestures)
- Implement more complex UI patterns
- Try building UI with SolidJS (transpiled to JS)
- Experiment with different SwiftUI components

## Resources

- SNSwiftUI README: `packages/sn-swiftui/README.md`
- Example apps: `packages/sn-swiftui/Examples/`
- Rust docs: `packages/sncore/src/`
