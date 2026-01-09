//
//  BasicExample.swift
//  SNSwiftUI Example
//
//  Example of using SNSwiftUI with the Rust core
//

import SwiftUI
import SNSwiftUI

struct BasicExample: View {
    @StateObject private var manager = SNSwiftUIManager()
    @State private var errorMessage: String?

    var body: some View {
        VStack(spacing: 20) {
            if let error = errorMessage {
                Text("Error: \(error)")
                    .foregroundColor(.red)
                    .padding()
            }

            SNSwiftUIRenderer(viewTree: manager.viewTree)
                .onAppear {
                    setupUI()
                }

            Button("Run JavaScript Example") {
                runJSExample()
            }
            .padding()
        }
    }

    private func setupUI() {
        do {
            // Initialize the core with our delegate
            try manager.initializeWithCore()

            // Method 1: Create root from core (this is the recommended way)
            let rootId = try manager.createRootFromCore(tag: "vstack")

            // Method 2: Build UI using JavaScript
            try manager.evaluateJavaScript("""
            // The root node is already created, so get it
            const rootId = '\(rootId)';

            // Create a button
            const buttonId = solidNative.createElement('button');
            solidNative.setProperty(buttonId, 'backgroundColor', 'blue');
            solidNative.setProperty(buttonId, 'cornerRadius', '8');
            solidNative.setProperty(buttonId, 'padding', '12');

            // Create button text
            const textId = solidNative.createTextNode('Hello from Rust + JS!');
            solidNative.setProperty(textId, 'color', 'white');

            // Build the tree
            solidNative.insertNode(buttonId, textId);
            solidNative.insertNode(rootId, buttonId);
            """)

            print("✅ UI setup complete!")
        } catch {
            errorMessage = error.localizedDescription
            print("❌ Error: \(error)")
        }
    }

    private func runJSExample() {
        do {
            guard let rootId = manager.getRootFromCore() else {
                errorMessage = "No root node found"
                return
            }

            // Add another element via JavaScript
            try manager.evaluateJavaScript("""
            const rootId = '\(rootId)';

            // Create a text element
            const textId = solidNative.createTextNode('JavaScript works! 🎉');
            solidNative.setProperty(textId, 'color', 'green');

            solidNative.insertNode(rootId, textId);
            """)

            print("✅ JavaScript executed successfully!")
        } catch {
            errorMessage = error.localizedDescription
            print("❌ JavaScript error: \(error)")
        }
    }
}

// MARK: - Preview

#Preview {
    BasicExample()
}
