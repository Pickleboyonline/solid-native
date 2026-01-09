//
//  ManualTreeExample.swift
//  SNSwiftUI Example
//
//  Example of building a tree manually without the Rust core
//

import SwiftUI
import SNSwiftUI

struct ManualTreeExample: View {
    @StateObject private var manager = SNSwiftUIManager()

    var body: some View {
        VStack(spacing: 20) {
            Text("Manual Tree Building Example")
                .font(.headline)

            SNSwiftUIRenderer(viewTree: manager.viewTree)
                .onAppear {
                    buildManualTree()
                }

            Button("Add More Content") {
                addMoreContent()
            }
            .padding()
        }
    }

    private func buildManualTree() {
        // Create root without the Rust core
        let root = manager.createRoot(type: "vstack")
        manager.setProperty(node: root, key: "spacing", value: "16")
        manager.setProperty(node: root, key: "padding", value: "20")

        // Create a title
        let title = manager.createTextNode(text: "Welcome to SNSwiftUI!")
        manager.setProperty(node: title, key: "color", value: "blue")
        manager.insertNode(parent: root, child: title)

        // Create a button
        let button = manager.createElement(type: "button")
        manager.setProperty(node: button, key: "backgroundColor", value: "green")
        manager.setProperty(node: button, key: "cornerRadius", value: "8")
        manager.setProperty(node: button, key: "padding", value: "12")

        let buttonText = manager.createTextNode(text: "Click Me!")
        manager.setProperty(node: buttonText, key: "color", value: "white")

        manager.insertNode(parent: button, child: buttonText)
        manager.insertNode(parent: root, child: button)

        print("✅ Manual tree built successfully!")
        manager.debugPrintTree()
    }

    private func addMoreContent() {
        // Find the root node
        guard let root = manager.viewTree.root else {
            print("❌ No root node found")
            return
        }

        // Add a new text element
        let newText = manager.createTextNode(text: "Dynamically added content ✨")
        manager.setProperty(node: newText, key: "color", value: "purple")
        manager.insertNode(parent: root, child: newText)

        print("✅ Added new content!")
    }
}

// MARK: - Preview

#Preview {
    ManualTreeExample()
}
